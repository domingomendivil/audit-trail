package db

import (
	"fmt"
	"mingo/audit/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddEvent(t *testing.T) {
	dao := GetDAO()
	var ev model.Event
	ev.CreatedTime = "2021-03-07 20:15"
	ev.EventType = "CUSTOMER-DELETED"
	ev.Gravity = "MEDIUM"
	ev.DynamicData = "Dun"
	dao.Insert(ev)
	var query model.EventQuery
	query.EventType = "CUSTOMER-DELETED"
	query.StartDate = "2020-03-07 20:15"
	query.EndDate = "2023-03-07 20:15"
	events, err := dao.GetEvents(query)
	assert.Nil(t, err)
	assert.NotNil(t, events)
	assert.Equal(t, "CUSTOMER-DELETED", events[0].EventType)
	assert.Equal(t, 1, len(events))
}

func TestAddEvent2(t *testing.T) {
	dao := GetDAO()
	var ev model.Event
	ev.CreatedTime = "2025-03-07 20:15"
	ev.EventType = "CUSTOMER-ADDED"
	ev.Gravity = "MEDIUM"
	ev.DynamicData = "Dun"
	dao.Insert(ev)
	var query model.EventQuery
	query.EventType = "CUSTOMER-ADDED"
	query.StartDate = "2020-03-07 20:15"
	query.EndDate = "2023-03-07 20:15"
	events, err := dao.GetEvents(query)
	assert.Nil(t, err)
	assert.Nil(t, events)
	fmt.Println(len(events))

	assert.Equal(t, 0, len(events))
}

func TestAddEvent3(t *testing.T) {
	dao := GetDAO()
	var ev model.Event
	ev.CreatedTime = "2025-03-07 20:15"
	ev.EventType = "CUSTOMER-ADDED"
	ev.Gravity = "MEDIUM"
	ev.DynamicData = "Dun"
	dao.Insert(ev)
	var query model.EventQuery
	query.EventType = "CUSTOMER-ADDED"
	query.StartDate = "2020-03-07 20:15"
	query.EndDate = "2023-03-07 20:15"

	events, err := dao.GetEvents(query)
	assert.Nil(t, err)
	assert.Nil(t, events)
	fmt.Println(len(events))
}

func TestAddEvent5(t *testing.T) {
	dao := GetDAO()
	var ev model.Event
	var ev1 model.Event
	ev.CreatedTime = "2025-03-07 20:15"
	ev.EventType = "DNS-ATTACK"
	ev.Gravity = "MEDIUM"
	ev.DynamicData = "Dun"
	ev1.CreatedTime = "2025-03-07 21:15"
	ev1.EventType = "DNS-ATTACK"
	ev1.Gravity = "HIGH"
	ev1.DynamicData = "Dun"
	dao.Insert(ev)
	dao.Insert(ev1)
	var query model.EventQuery
	query.EventType = "DNS-ATTACK"
	query.StartDate = "2024-03-07 20:15"
	query.EndDate = "2026-03-07 20:15"
	events, err := dao.GetEvents(query)
	assert.Nil(t, err)
	assert.NotNil(t, events)
	assert.Equal(t, 2, len(events))
	var query2 model.EventQuery
	query2.EventType = "USERS-CHANGED"
	query2.StartDate = "2024-03-07 20:15"
	query2.EndDate = "2026-03-07 20:15"
	events2, err := dao.GetEvents(query2)
	assert.Nil(t, err)
	assert.Nil(t, events2)

}
