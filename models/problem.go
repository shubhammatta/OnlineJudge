package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Problem struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Statement     string        `bson:"name"`
	Tests         string        `bson:"email"`
}

func (u *Problem) NewProblem(db *mgo.Database, Statement string, Tests string) {
	u.Statement = Statement
	u.Tests = Tests
	u.ID = bson.NewObjectId()
	c := db.C("problem")
	c.Insert(&u)
}

func (u *Problem) Get(db *mgo.Database, id string) error {
	if bson.IsObjectIdHex(id) {
		return db.C("problem").FindId(bson.ObjectIdHex(id)).One(&u)
	} else {
		return errors.New("It is not an ID")
	}
}
