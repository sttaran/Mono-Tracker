package app

import (
	"mono-tracker/internal/domain/fundraising"
)

func (a *App) CreateFundraising(fundraisingURL fundraising.FundraisingURL) (int, error) {
	id, err := a.fundraisingService.CreateFundraising(fundraising.NewFundraising(fundraisingURL))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (a *App) GetFundraisingList() ([]*fundraising.FundraisingWithHistory, error) {
	f, err := a.fundraisingService.GetFundraisings()
	if err != nil {
		println(err.Error())
		return []*fundraising.FundraisingWithHistory{}, err
	}

	return f, nil
}

func (a *App) SyncFundraising(id int) error {
	err := a.fundraisingService.SynchronizeFundraising(id)
	if err != nil {
		return err
	}

	return nil
}
