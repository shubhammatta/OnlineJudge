package controllers

import (
	"../models"
	"../utilities"
	"encoding/json"
//	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
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

func (u *Problem) get_Assignment(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	data := map[string]string{"unique_Id": ""}
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(403)
	}
	db := utilities.GetDB(r)
	Assignment := new(models.Assignment)
	AssignmentId := Assignment.Get(db, data["unique_Id"])
	var result []Problem
	err = db.C("problem").Find(bson.M{"assignmentId" : AssignmentId}).All(&result)
	fmt.Println(result)
	if err == nil{
		outData, _ := json.Marshal(result)
		w.Write(outData)
	} else{
		w.WriteHeader(404)
	}	
}

func (u * Problem) create_problem(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	data := map[string]string{"statement": "", "test": "", "unique_assignment": ""}
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	db := utilities.GetDB(r)
	A := new(models.Assignment)
	Id := A.Get(db , data["unique_assignment"])
	problem := new(models.Problem)
	problem.NewProblem(db , data["statement"] , data["test"] , Id)
}