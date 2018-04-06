package main

import (
	"html/template"
	"net/http"

	"github.com/aneelraju/udemy/WebDevWithGolang/work/042_mongodb/10_hands-on_solution/controllers"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	c := controllers.NewController(tpl)
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/bar", c.Bar)
	http.HandleFunc("/signup", c.SignUp)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// % go run main.go; # start server
// http://localhost:8080
