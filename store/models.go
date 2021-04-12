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

type SearchEventOptions struct {
	AgentcyID string
	From, To  time.Time

	// Available : ASC DESC
	// TODO: We could make this more generic by also including field name
	Order string

	PagingOpts *PagingOptions
}

type SearchEventResult struct {
	Events   []*Event
	PageInfo PagingInfo
}

type PagingOptions struct {
	Offset       uint64
	ItemsPerPage uint64
}

type PagingInfo struct {
	Total uint64
	Offet uint64
}
