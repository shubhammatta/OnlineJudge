package controllers

import (
	"../models"
	"../utilities"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
//	"github.com/gorilla/mux"
	"net/http"
	"fmt"

)

type Submission struct{}

func (u * Submission) create_submission(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	data := map[string]string{"problemId": "",  "userId": ""}
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	db := utilities.GetDB(r)
	A := new(models.Submission)
	A.NewSubmission(db , data["problemId"] , data["userId"])
	
}

func (u * Submission) get_submission_by_user(w http.ResponseWriter , r * http.Request) {
	decoder := json.NewDecoder(r.Body)
	data := map[string]string{"userId": ""}
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	db := utilities.GetDB(r)
	var result []Submission
	err = db.C("submission").Find(bson.M{"userId" : data["userId"]}).All(&result)
	fmt.Println(result)
	if err == nil{
		outData, _ := json.Marshal(result)
		w.Write(outData)
	} else{
		w.WriteHeader(404)
	}	
}