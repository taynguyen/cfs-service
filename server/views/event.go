package views

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/cfs-service/store"
	"gitlab.com/cfs-service/utils"
)

// Define models that used by the handler
type Event struct {
	ID            string `json:"event_id"`
	AgencyID      string `json:"agency_id" validate:"required"`
	Number        string `json:"event_number" validate:"required"`
	TypeCode      string `json:"event_type_code" validate:"required"`
	ResponderCode string `json:"responder" validate:"required"`
	CreatedTime   string `json:"event_time" validate:"required"`
	DispatchTime  string `json:"dispatch_time" validate:"required"`
}

// EventDatetime to customize marshal and unmarshal datetime in JSON
type EventDatetime struct {
	time.Time
}

func (e *Event) ToStoreModel() (*store.Event, error) {
	m := &store.Event{
		ID:            e.ID,
		AgencyID:      e.AgencyID,
		Number:        e.Number,
		TypeCode:      e.TypeCode,
		ResponderCode: e.ResponderCode,
	}

	var err error
	m.CreatedTime, err = utils.StringToTime(e.CreatedTime)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid created time")
	}

	m.DispatchTime, err = utils.StringToTime(e.DispatchTime)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid dispatch time")
	}

	return m, nil
}

type GetEventsResponse struct {
	Events     []*Event
	PagingInfo store.PagingInfo
}

func GetEventResponseFromStoreResponse(r *store.SearchEventResult) *GetEventsResponse {
	m := &GetEventsResponse{
		Events:     make([]*Event, len(r.Events)),
		PagingInfo: r.PageInfo,
	}

	for i := 0; i < len(r.Events); i++ {
		m.Events = append(m.Events, EventFromStoreModel(r.Events[i]))
	}

	return m
}

func EventFromStoreModel(m *store.Event) *Event {
	return &Event{
		ID:            m.ID,
		AgencyID:      m.AgencyID,
		Number:        m.Number,
		TypeCode:      m.TypeCode,
		ResponderCode: m.ResponderCode,
		CreatedTime:   utils.FormatTimeToLongTring(m.CreatedTime),
		DispatchTime:  utils.FormatTimeToLongTring(m.DispatchTime),
	}
}

func (e *Event) Validate() []*ValidationErrorResponse {
	return ValidateStruct(e)
}
