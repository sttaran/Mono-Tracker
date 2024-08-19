package fundraising

type Platform string

const (
	PlatformMonobank Platform = "monobank"
)

type FundraisingURL string

type Fundraising struct {
	ID          int            `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Description string         `json:"description" db:"description"`
	Goal        float64        `json:"goal" db:"goal"`
	URL         FundraisingURL `json:"url" db:"url"`
}

type FundraisingHistory struct {
	ID            int     `json:"id" db:"id"`
	FundraisingID int     `json:"fundraising_id" db:"fundraising_id"`
	Raised        float64 `json:"raised" db:"raised"`
	SyncTime      string  `json:"sync_time" db:"sync_time"`
}

func NewFundraising(url FundraisingURL) *Fundraising {
	return &Fundraising{
		ID:          0,
		Name:        "",
		Description: "",
		Goal:        0,
		URL:         url,
	}
}
