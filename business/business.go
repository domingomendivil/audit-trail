package business

import (
	"fmt"
	"mingo/audit/db"
	"mingo/audit/model"
	"sync"

	"gopkg.in/go-playground/validator.v9"
)

type AuditTrail interface {
	AddEvent(ev model.Event) error
	Free()
	GetEvents(query model.EventQuery) ([]model.Event, error)
}

type AuditTrailImpl struct {
	dao db.DAO
}

func (auditTrail *AuditTrailImpl) GetEvents(query model.EventQuery) ([]model.Event, error) {
	fmt.Println("AuditTrail.GetEvents")
	return auditTrail.dao.GetEvents(query)
}

func (auditTrail *AuditTrailImpl) AddEvent(ev model.Event) error {
	fmt.Println("AuditTrail.AddEvent")
	v := validator.New()
	fmt.Println("validation errors")
	err := v.Struct(ev)
	if err != nil {

		return err
	} else {
		auditTrail.dao.Insert(ev)
	}
	return nil

}

var auditTrail *AuditTrailImpl
var mu sync.Mutex

func GetAuditTrail() AuditTrail {
	mu.Lock()
	if auditTrail == nil {
		auditTrail = new(AuditTrailImpl)
		auditTrail.dao = db.GetDAO()
	}
	mu.Unlock()
	return auditTrail
}

func (auditTrail *AuditTrailImpl) Free() {
	auditTrail.dao.Free()
}
