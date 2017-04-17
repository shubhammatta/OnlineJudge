package utilities

import (
	"gopkg.in/mgo.v2"
	"net/http"
	
)

const (
	mongoURI = "localhost:27017"
	dbname   = "test"
)

var db * mgo.Database

func GetDB(r *http.Request) *mgo.Database {
	ds, err := mgo.Dial(mongoURI)

	if err != nil {
		panic(err)
	}

	db = ds.DB(dbname)
	return db
}
