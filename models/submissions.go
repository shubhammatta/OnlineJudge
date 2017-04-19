package models

import (
	// "crypto/md5"
	// "encoding/hex"
	// "errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "io"
	"time"
)

type Submission struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	ProblemId     string        `bson:"problemId"`
	UserId        string        `bson:"userId"`
	Timestamp     time.Time        `bson:"timestamp"`
	FilePath	  string 		`bson:"filePath"`
}

func (u *Submission) NewSubmission(db *mgo.Database, problemId string, userId string) {
	u.ProblemId = problemId
	u.UserId= userId
	u.ID = bson.NewObjectId()
	u.Timestamp = time.Now()
	u.FilePath = "../../../data/submissions/" + u.ID.String()
	c := db.C("submission")
	c.Insert(&u)
}
/*
func (u *User) Get(db *mgo.Database, id string) error {
	if bson.IsObjectIdHex(id) {
		return db.C("user").FindId(bson.ObjectIdHex(id)).One(&u)
	} else {
		return errors.New("It is not an ID")
	}
}

func (u *User) Authenticate(db *mgo.Database, email string, password string) error {
	h := md5.New()
	io.WriteString(h, password)
	hex_password := hex.EncodeToString(h.Sum(nil))
	err := db.C("user").Find(map[string]string{
		"password" : hex_password,
		"email" :    email,
	}).One(&u)
	return err
}*/
