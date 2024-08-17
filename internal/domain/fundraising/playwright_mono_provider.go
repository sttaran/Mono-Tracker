package fundraising

import (
	"errors"
	"fmt"
	"github.com/playwright-community/playwright-go"
	urlNet "net/url"
	"strconv"
	"strings"
)

type PlaywrightMonoFundraisingService struct {
	pw *playwright.Playwright
}

func NewPlaywrightMonoFundraisingService(pw *playwright.Playwright) IFundraisingInfoProvider {
	return &PlaywrightMonoFundraisingService{
		pw: pw,
	}
}

func (s *PlaywrightMonoFundraisingService) SynchronizeFundraising(url string, initialSync bool) (*FundaisingInfo, error) {
	var name, description, goal string
	var goalParsed, raisedParsed float64

	browser, err := s.pw.Chromium.Launch()
	if err != nil {
		return nil, fmt.Errorf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		return nil, fmt.Errorf("could not create page: %v", err)
	}
	defer func(page playwright.Page) {
		if err = page.Close(); err != nil {
			println("could not close page: %v", err)
		}
	}(page)

	if _, err = page.Goto(string(url)); err != nil {
		return nil, fmt.Errorf("could not goto: %v", err)
	}

	errFinished := page.Locator(".done-jar-sub-text").WaitFor(
		playwright.LocatorWaitForOptions{
			Timeout: playwright.Float(2000),
		})
	if errFinished == nil {
		return nil, errors.New("fundraising is finished")
	}

	raised, err := page.Locator(".stats-data").First().Locator(".stats-data-value").InnerText()
	raised = strings.Replace(raised, "₴", "", -1)
	raised = strings.ReplaceAll(raised, "\u00a0", "")

	// convert raised to float
	raisedParsed, err = strconv.ParseFloat(raised, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse raised: %v", err)
	}

	goal, err = page.Locator(".stats-data").Last().Locator(".stats-data-value").InnerText()
	if err != nil {
		return nil, fmt.Errorf("could not get goal: %v", err)
	}

	if initialSync {
		goal = strings.Replace(goal, "₴", "", -1)
		goal = strings.ReplaceAll(goal, "\u00a0", "")
		// convert goal to float
		goalParsed, err = strconv.ParseFloat(goal, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse goal: %v", err)
		}

		description, err = page.Locator(".description-box").InnerText()
		name, err = page.Locator("h1").InnerText()
	}

	return &FundaisingInfo{
		Name:        name,
		Description: description,
		Goal:        goalParsed,
		Raised:      raisedParsed,
	}, nil
}

func (s *PlaywrightMonoFundraisingService) IsFundraisingValid(url string) (bool, error) {
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
