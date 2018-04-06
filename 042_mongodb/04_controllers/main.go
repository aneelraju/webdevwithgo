package main

import (
	"net/http"

	"github.com/aneelraju/udemy/WebDevWithGolang/work/042_mongodb/04_controllers/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

/*
% go run main.go; # start server
% curl http://localhost:8080/user/1; # GET
% curl -X POST -H "Content-Type: application/json" -d '{"Name":"James Bond","Gender":"male","Age":32,"Id":"777"}' http://localhost:8080/user; # POST
// -X is short for --request Specifies a custom request method to use when communicating with the HTTP server.
// -H is short for --header
// -d is short for --data
% curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/777; # DELETE
*/
