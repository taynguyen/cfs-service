package store

type IStore interface {
	Migrate() error

	AddEvent(ev *Event) error
	SearchEventByTimeRange(o *SearchEventOptions) (*SearchEventResult, error)
}
