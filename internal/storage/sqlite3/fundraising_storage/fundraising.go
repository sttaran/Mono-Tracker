package fundraising_storage

import (
	"mono-tracker/internal/domain/fundraising"
	"mono-tracker/internal/domain/fundraising_history"
	"time"
)

func (s *Storage) CreateFundraising(f *fundraising.Fundraising) (int, error) {
	result, err := s.db.Exec("INSERT INTO fundraising (name, description, goal, url) VALUES (?, ?, ?, ?)", f.Name, f.Description, f.Goal, f.URL)
	if err != nil {
		return 0, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id64), nil
}

func (s *Storage) GetFundraisings() ([]*fundraising.FundraisingWithHistory, error) {
	rows, err := s.db.Query("SELECT * FROM fundraising")
	if err != nil {
		return nil, err
	}

	fundraisings := []*fundraising.FundraisingWithHistory{}
	for rows.Next() {
		fundraising := &fundraising.FundraisingWithHistory{}
		err = rows.Scan(&fundraising.ID, &fundraising.Name, &fundraising.Description, &fundraising.Goal, &fundraising.URL)
		if err != nil {
			return nil, err
		}
		fundraising.History, err = s.getFundraisingHistory(fundraising.ID)
		if err != nil {
			return nil, err
		}
		fundraisings = append(fundraisings, fundraising)
	}

	return fundraisings, nil
}

func (s *Storage) getFundraisingHistory(id int) ([]fundraising_history.FundraisingHistory, error) {
	rows, err := s.db.Query("SELECT * FROM fundraising_history WHERE fundraising_id = ? ORDER BY sync_time DESC", id)
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

func (s *Storage) AddFundraisingHistory(fundraisingID int, raised float64) error {
	_, err := s.db.Exec("INSERT INTO fundraising_history (fundraising_id, raised, sync_time) VALUES (?, ?, ?)", fundraisingID, raised, time.Now().Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteFundraising(id int) error {
	_, err := s.db.Exec("DELETE FROM fundraising WHERE id = ?", id)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("DELETE FROM fundraising_history WHERE fundraising_id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetFundraisingById(id int) (*fundraising.Fundraising, error) {
	row := s.db.QueryRow("SELECT * FROM fundraising WHERE id = ?", id)
	f := &fundraising.Fundraising{}
	err := row.Scan(&f.ID, &f.Name, &f.Description, &f.Goal, &f.URL)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *Storage) UpdateFundraising(fundraisingID int, info *fundraising.Fundraising) error {
	_, err := s.db.Exec("UPDATE fundraising SET name = ?, description = ?, goal = ? WHERE id = ?", info.Name, info.Description, info.Goal, fundraisingID)
	if err != nil {
		return err
	}
	return nil
}
