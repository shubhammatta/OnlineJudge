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
*/
func (u *Problem) GetById(db *mgo.Database, id string) error {
	if bson.IsObjectIdHex(id) {
		return db.C("problem").FindId(bson.ObjectIdHex(id)).One(&u)
	} else {
		return errors.New("It is not an ID")
	}
}

func (u *Problem) get_Assignment(db * mgo.Database , unique_assignment_ID string) error {
	Assignment := new(models.assignment)
	AssignmentId := Assignment.GET(db, unique_assignment_ID)
	result := db.C("problem").Find(bson.M{"assignmentId" : AssignmentId})
	fmt.Println(result)
	if result != nil{
		return result
	} else {
		return errors.New("Not found")
	}	
}

func (u * Problem) create_problem(w http.ResponseWriter, r *http.Request){

}