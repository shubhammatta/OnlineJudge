package main

import (
	"./controllers"
	"./utilities"
	"github.com/gorilla/mux"
)

func LoadRoutes() *mux.Router {
	index := controllers.Index{}
	auth := controllers.Auth{}
	user := controllers.User{}
	assignment := controllers.Assignment{}
	problem := controllers.Problem{}
	router := mux.NewRouter()
	router.HandleFunc("/", index.Welcome).Methods("GET")
	router.HandleFunc("/sitemap", index.Sitemap)

	a := router.PathPrefix("/auth").Subrouter()
	a.HandleFunc("/login", auth.Login).Methods("POST")
	a.HandleFunc("/logout", auth.Logout).Methods("GET")
	a.HandleFunc("/user", auth.User).Methods("GET")
	a.HandleFunc("/register", auth.Register).Methods("POST")

	a = router.PathPrefix("/users").Subrouter()
	a.HandleFunc("/profile", utilities.AuthenticationHandler(user.Profile)).Methods("GET")
	// Adding problem-submission code from here
	a.HandleFunc("/assignment", utilities.AuthenticationHandler(user.Profile)).Methods("GET")
	a.HandleFunc("/{id}", utilities.AuthenticationHandler(user.Get)).Methods("GET")
	a = router.PathPrefix("/create").Subrouter()
	// Adding assignment-submission code from here
	a.HandleFunc("/assignment", utilities.AuthenticationHandler(assignment.Create_assignment)).Methods("POST")
	a.HandleFunc("/problem",  utilities.AuthenticationHandler(problem.Create_problem)).Methods("POST")
	router.HandleFunc("/assignment/{id}", utilities.AuthenticationHandler(problem.Get_Assignment)).Methods("GET")
	return router
}
