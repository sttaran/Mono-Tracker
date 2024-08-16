package fundraising

import (
	"errors"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"mono-tracker/internal/domain/fundraising_history"
	"mono-tracker/pkg"
	urlNet "net/url"
	"strconv"
	"strings"
	"time"
)

type PlaywrightMonoFundraisingService struct {
	db *pkg.SQLiteClient
	pw *playwright.Playwright
}

func (s *PlaywrightMonoFundraisingService) GetFundraisingHistory(id int) ([]fundraising_history.FundraisingHistory, error) {
	rows, err := s.db.Db.Query("SELECT * FROM fundraising_history WHERE fundraising_id = ? ORDER BY sync_time DESC", id)
	if err != nil {
		return nil, err
	}
	history := []fundraising_history.FundraisingHistory{}
	for rows.Next() {
		h := fundraising_history.FundraisingHistory{}
		err = rows.Scan(&h.ID, &h.FundraisingID, &h.Raised, &h.SyncTime)
		if err != nil {
			return nil, err
		}
		history = append(history, h)
	}
	return history, nil
}

func (s *PlaywrightMonoFundraisingService) SynchronizeFundraising(id int, initialSync bool) error {
	var url string
	err := s.db.Db.QueryRow("SELECT url FROM fundraising WHERE id = ?", id).Scan(&url)
	if err != nil {
		return err
	}

	browser, err := s.pw.Chromium.Launch()
	if err != nil {
		return fmt.Errorf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		return fmt.Errorf("could not create page: %v", err)
	}
	defer func(page playwright.Page) {
		if err = page.Close(); err != nil {
			println("could not close page: %v", err)
		}
	}(page)

	if _, err = page.Goto(url); err != nil {
		return fmt.Errorf("could not goto: %v", err)
	}

	errFinished := page.Locator(".done-jar-sub-text").WaitFor(
		playwright.LocatorWaitForOptions{
			Timeout: playwright.Float(2000),
		})
	if errFinished == nil {
		return errors.New("fundraising is finished")
	}

	raised, err := page.Locator(".stats-data").First().Locator(".stats-data-value").InnerText()
	raised = strings.Replace(raised, "₴", "", -1)
	raised = strings.ReplaceAll(raised, "\u00a0", "")

	// convert raised to float
	raisedParsed, err := strconv.ParseFloat(raised, 64)
	if err != nil {
		return fmt.Errorf("could not parse raised: %v", err)
	}

	goal, err := page.Locator(".stats-data").Last().Locator(".stats-data-value").InnerText()
	if err != nil {
		return fmt.Errorf("could not get goal: %v", err)
	}

	if initialSync {
		goal = strings.Replace(goal, "₴", "", -1)
		goal = strings.ReplaceAll(goal, "\u00a0", "")
		// convert goal to float
		goalParsed, err := strconv.ParseFloat(goal, 64)
		if err != nil {
			return fmt.Errorf("could not parse goal: %v", err)
		}

		description, err := page.Locator(".description-box").InnerText()
		name, err := page.Locator("h1").InnerText()

		_, err = s.db.Db.Exec("UPDATE fundraising SET name = ?, description = ?, goal = ? WHERE url = ?", name, description, goalParsed, url)
		if err != nil {
			return err
		}
	}

	_, err = s.db.Db.Exec("INSERT INTO fundraising_history (fundraising_id, raised, sync_time) VALUES (?, ?, ?)", id, raisedParsed, time.Now().Format(time.RFC3339))
	if err != nil {
		return err
	}
	return nil

}

func NewFundraisingService(db *pkg.SQLiteClient, pw *playwright.Playwright) IFundraisingService {
	return &PlaywrightMonoFundraisingService{
		db: db,
		pw: pw,
	}
}

func (s *PlaywrightMonoFundraisingService) GetFundraisingByID(id int) (*Fundraising, error) {
	return &Fundraising{}, nil
}

func (s *PlaywrightMonoFundraisingService) GetFundraisings() ([]*FundraisingWithHistory, error) {
	rows, err := s.db.Db.Query("SELECT * FROM fundraising")
	if err != nil {
		return nil, err
	}

	fundraisings := []*FundraisingWithHistory{}
	for rows.Next() {
		fundraising := &FundraisingWithHistory{}
		err = rows.Scan(&fundraising.ID, &fundraising.Name, &fundraising.Description, &fundraising.Goal, &fundraising.URL)
		if err != nil {
			return nil, err
		}
		fundraising.History, err = s.GetFundraisingHistory(fundraising.ID)
		if err != nil {
			return nil, err
		}
		fundraisings = append(fundraisings, fundraising)
	}

	return fundraisings, nil
}

func (s *PlaywrightMonoFundraisingService) isFundraisingValid(url string) (bool, error) {
	_, err := urlNet.ParseRequestURI(url)
	if err != nil {
		return false, errors.New("invalid url")
	}
	browser, err := s.pw.Chromium.Launch()
	if err != nil {
		return false, fmt.Errorf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		return false, fmt.Errorf("could not create page: %v", err)
	}

	defer func(page playwright.Page) {
		if err = page.Close(); err != nil {
			println("could not close page: %v", err)
		}
	}(page)

	if _, err = page.Goto(url); err != nil {
		return false, fmt.Errorf("could not goto: %v", err)
	}

	errFinished := page.Locator(".done-jar-sub-text").WaitFor(
		playwright.LocatorWaitForOptions{
			Timeout: playwright.Float(2000),
		})
	if errFinished == nil {
		return false, errors.New("fundraising is finished")
	}
	return true, nil
}

func (s *PlaywrightMonoFundraisingService) CreateFundraising(fundraising *Fundraising) (id int, err error) {
	if _, err = s.isFundraisingValid(string(fundraising.URL)); err != nil {
		return 0, err
	}

	result, err := s.db.Db.Exec("INSERT INTO fundraising (name, description, goal, url) VALUES (?, ?, ?, ?)", fundraising.Name, fundraising.Description, fundraising.Goal, fundraising.URL)
	if err != nil {
		return 0, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id64), nil

}

func (s *PlaywrightMonoFundraisingService) UpdateFundraising(fundraising *Fundraising) error {
	return nil
}

func (s *PlaywrightMonoFundraisingService) DeleteFundraising(id int) error {
	return nil
}
