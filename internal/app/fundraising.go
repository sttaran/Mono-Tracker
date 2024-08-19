package app

import (
	"github.com/labstack/gommon/log"
	"mono-tracker/internal/domain/fundraising"
)

func (a *App) CreateFundraising(fundraisingURL fundraising.FundraisingURL) (int, error) {
	id, err := a.fundraisingService.CreateFundraising(fundraising.NewFundraising(fundraisingURL))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (a *App) GetFundraisingList(dto fundraising.FetchListDTO) (*fundraising.FetchListResponse, error) {
	log.Info(dto)
	f, err := a.fundraisingService.GetFundraisings(&dto)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return f, nil
}

func (a *App) SyncFundraising(id int, initial bool) error {
	err := a.fundraisingService.SynchronizeFundraising(id, initial)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) DeleteFundraising(id int) error {
	err := a.fundraisingService.DeleteFundraising(id)
	if err != nil {
		return err
	}

	return nil
}
