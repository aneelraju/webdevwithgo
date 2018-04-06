package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"

	"github.com/aneelraju/udemy/WebDevWithGolang/work/042_mongodb/05_mongodb/05_update-user-controllers-delete/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}

/*
% mongod --dbpath mongodb/data/db_v3.6.3; # start mongodb server
% go run main.go; # start server
% curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user; #POST
% curl http://localhost:8080/user/<enter-user-id-here>
% curl -X DELETE http://localhost:8080/user/<enter-user-id-here>
*/
