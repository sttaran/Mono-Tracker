package fundraising_history

type FundraisingHistory struct {
	ID            int     `json:"id" db:"id"`
	FundraisingID int     `json:"fundraising_id" db:"fundraising_id"`
	Raised        float64 `json:"raised" db:"raised"`
	SyncTime      string  `json:"sync_time" db:"sync_time"`
}
