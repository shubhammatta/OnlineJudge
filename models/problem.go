package models

import (
//	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
//	"fmt"
//	"./assignment"
)

type Problem struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Statement     string        `bson:"name"`
	Tests         string        `bson:"test"`
	AssignmentId  string		`bson:"assignmentId"`
}

func (u *Problem) NewProblem(db *mgo.Database, Statement string, Tests string , unique_assignment_ID string) {
	u.Statement = Statement
	u.Tests = Tests
	u.AssignmentId = unique_assignment_ID
	u.ID = bson.NewObjectId()
	c := db.C("problem")
	c.Insert(&u)
}

