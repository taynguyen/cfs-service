package store

type IStore interface {
	Migrate() error

	AddEvent(ev *Event) error
}
