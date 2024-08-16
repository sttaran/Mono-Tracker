package fundraising

type IFundraisingService interface {
	GetFundraisings() ([]*FundraisingWithHistory, error)
	CreateFundraising(fundraising *Fundraising) (int, error)
	DeleteFundraising(id int) error
	SynchronizeFundraising(id int, initial bool) error
}
