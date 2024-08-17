package fundraising

type IFundraisingStorage interface {
	GetFundraisings() ([]*FundraisingWithHistory, error)
	GetFundraisingById(id int) (*Fundraising, error)
	CreateFundraising(fundraising *Fundraising) (int, error)
	AddFundraisingHistory(fundraisingID int, raised float64) error
	UpdateFundraising(fundraisingID int, info *Fundraising) error
	DeleteFundraising(id int) error
}

type IFundraisingInfoProvider interface {
	SynchronizeFundraising(url string, initial bool) (*FundaisingInfo, error)
	IsFundraisingValid(url string) (bool, error)
}

type IFundraisingService interface {
	GetFundraisings() ([]*FundraisingWithHistory, error)
	CreateFundraising(fundraising *Fundraising) (int, error)
	DeleteFundraising(id int) error
	SynchronizeFundraising(id int, initial bool) error
}
