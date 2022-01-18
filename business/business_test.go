package business

import (
	"mingo/audit/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DAOMock struct {
	mock.Mock
}

var daoMock DAOMock

func (m *DAOMock) GetEvents(query model.EventQuery) ([]model.Event, error) {
	args := m.Called(query)
	var events []model.Event
	var err error
	if args.Get(0) == nil {
		events = nil
	} else {
		events = (args.Get(0)).([]model.Event)
	}
	if args.Get(1) == nil {
		err = nil
	} else {
		err = (args.Get(1)).(error)
	}
	return events, err
}

func (m *DAOMock) Insert(event model.Event) {
	m.Called(event)
}
func (m *DAOMock) Free() {
	m.Called()
}

func givenNoEvents(query model.EventQuery) {
	var events []model.Event
	var error error
	daoMock.On("GetEvents", query).Return(events, error)
}

func givenEvents(query model.EventQuery, events []model.Event) {
	daoMock.On("GetEvents", query).Return(events, nil)
}

func whenGetEventsReturnEmpty(t *testing.T, query model.EventQuery) {
	auditTrail := AuditTrailImpl{&daoMock}
	events, err := auditTrail.GetEvents(query)
	assert.Nil(t, events)
	assert.Nil(t, err)
}

func whenGetEventsReturns(t *testing.T, query model.EventQuery, expected []model.Event) {
	auditTrail := AuditTrailImpl{&daoMock}
	events, err := auditTrail.GetEvents(query)
	for index, event := range events {
		assert.Equal(t, expected[index].EventType, event.EventType)
	}
	assert.Nil(t, err)
}

func whenAddEventReturnError(t *testing.T, event model.Event) {
	auditTrail := AuditTrailImpl{&daoMock}
	err := auditTrail.AddEvent(event)
	assert.NotNil(t, err)
}

func TestAddEvent(t *testing.T) {
	var event model.Event
	event.EventType = "CUSTOMER-ADDED"
	whenAddEventReturnError(t, event)

}

func TestGetEvents(t *testing.T) {
	var query model.EventQuery
	query.EventType = "CUSTOMER-ADDED"
	query.StartDate = "2020-03-07 20:15"
	query.EndDate = "2023-03-07 20:15"
	givenNoEvents(query)
	whenGetEventsReturnEmpty(t, query)
}

func TestGetEvents2(t *testing.T) {
	var query model.EventQuery
	query.EventType = "CUSTOMER-ADDED"
	query.StartDate = "2020-03-07 20:15"
	query.EndDate = "2023-03-07 20:15"
	var events []model.Event
	events = make([]model.Event, 1)
	events[0].CreatedTime = "2023-03-07 20:15"
	events[0].EventType = "CUSTOMER-CREATED"
	givenEvents(query, events)
	whenGetEventsReturns(t, query, events)
}
