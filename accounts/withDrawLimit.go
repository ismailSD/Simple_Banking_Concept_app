package accounts

import "time"

type WithDrawLimit struct {
	LastWithDrawDate time.Time
	LastLimitSet     float64
	Amount           float64 // by default set to -1.0 means there is no limit per day
}