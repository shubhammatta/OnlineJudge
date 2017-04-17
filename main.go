package main

import (
	"github.com/urfave/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"net/http"
)

func main() {
	store := cookiestore.New([]byte("secretkey789"))
	router := LoadRoutes()
	n := negroni.Classic()
	static := negroni.NewStatic(http.Dir("static"))
	static.Prefix = "/static"
	n.Use(static)
	//n.Use(negroni.HandlerFunc(MgoMiddleware))
	n.Use(sessions.Sessions("global_session_store", store))
	n.UseHandler(router)
	n.Run(":5000")
}

