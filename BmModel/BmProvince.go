package BmModel

import (
	"gopkg.in/mgo.v2/bson"
)

type Province struct {
	ID       string        `json:"-"`
	Id_      bson.ObjectId `json:"-" bson:"_id"`
	Title    string        `json:"title" bson:"TITLE"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a Province) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *Province) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *Province) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "title":
			rst["TITLE"] = v[0]
		}
	}
	return rst
}
