package store

import "time"

type Event struct {
	ID            string
	AgencyID      string
	Number        string
	TypeCode      string
	ResponderCode string
	CreatedTime   time.Time
	DispatchTime  time.Time
}
