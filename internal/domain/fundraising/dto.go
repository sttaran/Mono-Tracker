package fundraising

type FundraisingWithHistory struct {
	Fundraising
	History []FundraisingHistory `json:"history"`
}

type FundaisingInfo struct {
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Goal        float64 `json:"goal" db:"goal"`
	Raised      float64 `json:"raised" db:"raised"`
}

type fetchListDTOSort struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}

type FetchListDTO struct {
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
	Sort  fetchListDTOSort `json:"sort"`
}

type FetchListResponse struct {
	Total int                       `json:"total"`
	Data  []*FundraisingWithHistory `json:"data"`
}
