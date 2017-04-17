package controllers

import (
	"../models"
	"../utilities"
	"encoding/json"
//	"github.com/gorilla/mux"
	"net/http"
	"fmt"

)

type Assignment struct{}

func (u * Problem) create_assignment(w http.ResponseWriter, r *http.Request){
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