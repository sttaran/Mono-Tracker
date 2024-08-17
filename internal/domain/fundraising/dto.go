package fundraising

import "mono-tracker/internal/domain/fundraising_history"

type FundraisingWithHistory struct {
	Fundraising
	History []fundraising_history.FundraisingHistory `json:"history"`
}

type FundaisingInfo struct {
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Goal        float64 `json:"goal" db:"goal"`
	Raised      float64 `json:"raised" db:"raised"`
}
