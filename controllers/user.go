package controllers

import (
	"../models"
	"../utilities"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type User struct{}

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

func (u *User) Profile(w http.ResponseWriter, r *http.Request){
	user_id, _ := utilities.GetUserId(r)
	db := utilities.GetDB(r)
	user := new(models.User)
	user.Get(db, user_id)
	user.Password = ""
	fmt.Println(user)
	type  output struct{
		ID   bson.ObjectId
		Name   string
		Email   string	
	}
	value := output {
		ID : user.ID,
		Name : user.Name,
		Email : user.Email,	
	}
	out, _ := json.Marshal(value)
	w.Write(out)
}

// Adding code from here
func (u *User) Problems(w http.ResponseWriter, r *http.Request){
	user_id, _ := utilities.GetUserId(r)
	db := utilities.GetDB(r)
	user := new(models.User)
	user.Get(db, user_id)
	user.Password = ""
	out, _ := json.Marshal(user)
	w.Write(out)
}
