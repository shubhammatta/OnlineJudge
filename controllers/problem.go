package controllers

import (
	"../models"
	"../utilities"
	"encoding/json"
//	"github.com/gorilla/mux"
	"net/http"
//	"fmt"
//	"gopkg.in/mgo.v2/bson"
)

type Problem struct{}
/*
func (u *User) Get(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	db := utilities.GetDB(r)
	user := new(models.User)
	err := user.Get(db, id)
	if err != nil {
		w.WriteHeader(404)
	} else {
		user.Password = ""
		out, _ := json.Marshal(user)
		w.Write(out)
	}
}
*/


// Adding code from here
func (u *Problem) Problems(w http.ResponseWriter, r *http.Request){
	user_id, _ := utilities.GetUserId(r)
	db := utilities.GetDB(r)
	user := new(models.User)
	user.Get(db, user_id)
	user.Password = ""
	out, _ := json.Marshal(user)
	w.Write(out)
}

// Adding code from here
func (u *Problem) Problem_s(w http.ResponseWriter, r *http.Request){
	user_id, _ := utilities.GetUserId(r)
	db := utilities.GetDB(r)
	user := new(models.User)
	user.Get(db, user_id)
	user.Password = ""
	out, _ := json.Marshal(user)
	w.Write(out)
}