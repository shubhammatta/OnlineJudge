package models

import (
//	"crypto/md5"
//	"encoding/hex"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
//	"io"
	"fmt"
)

type Assignment struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	UniqueId string 		`bson : "uniqueId"`
}

func (u *Assignment) NewAssignment(db *mgo.Database, name string , uniqueId string) {
	u.Name = name
	u.UniqueId = uniqueId;
	c := db.C("assignment")
	c.Insert(&u)
}

func (u *Assignment) Get(db *mgo.Database, uniqueId string) error {
	result := db.C("assignment").Find(bson.M{"UniqueId" : uniqueId}).One(&u)
	if result != nil {
		fmt.Println("assignment   " , result)
		return result
	} else {
		return errors.New("Assignment Not found")
	}
}
/*
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