package fundraising

import "mono-tracker/internal/domain/fundraising_history"

type IFundraisingService interface {
	GetFundraisingByID(id int) (*Fundraising, error)
	GetFundraisings() ([]*FundraisingWithHistory, error)
	GetFundraisingHistory(id int) ([]fundraising_history.FundraisingHistory, error)
	CreateFundraising(fundraising *Fundraising) (int, error)
	UpdateFundraising(fundraising *Fundraising) error
	DeleteFundraising(id int) error
	SynchronizeFundraising(id int, initial bool) error
}
