package controllers

import (
	"../models"
	"../utilities"
	"encoding/json"
	// "github.com/gorilla/mux"
	"net/http"
	"fmt"

)

type Assignment struct{}

func (u * Assignment) Create_assignment(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	data := map[string]string{"name": "",  "unique_assignment": ""}
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	db := utilities.GetDB(r)
	A := new(models.Assignment)
	A.NewAssignment(db , data["name"] , data["unique_assignment"])
}

// func (u * Assignment) Get(w http.ResponseWriter, r *http.Request){
// 	vars = mux.Vars(r)
// 	id := vars["id"]
// 	db := utilities.GetDB(r)
// 	assignment := new(models.Assignment)
// 	err := assignment.Get(db , id)
// 	if err != nil{
// 		w.WriteHeader(404)
// 	}else{
// 		w.write(assignment)
// 	}

// }