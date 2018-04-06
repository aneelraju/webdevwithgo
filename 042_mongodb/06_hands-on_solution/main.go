package main

import (
	"net/http"

	"github.com/aneelraju/udemy/WebDevWithGolang/work/042_mongodb/06_hands-on_solution/controllers"
	"github.com/aneelraju/udemy/WebDevWithGolang/work/042_mongodb/06_hands-on_solution/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return make(map[string]models.User)
}

/*
% go run main.go; # start server
% curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user; #POST
% curl http://localhost:8080/user/<enter-user-id-here>
% curl -X DELETE http://localhost:8080/user/<enter-user-id-here>
*/
