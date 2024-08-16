package fundraising

type FundraisingURL string

type Fundraising struct {
	ID          int            `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Description string         `json:"description" db:"description"`
	Goal        float64        `json:"goal" db:"goal"`
	URL         FundraisingURL `json:"url" db:"url"`
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
