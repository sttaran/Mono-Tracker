package fundraising

import "mono-tracker/internal/domain/fundraising_history"

type FundraisingWithHistory struct {
	Fundraising
	History []fundraising_history.FundraisingHistory `json:"history"`
}
