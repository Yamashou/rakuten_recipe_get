package main

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

func init() {
	loadTemplates()
}

func main() {

	server = gin.Default()

	server.Static("/public/css/", "./public/css")
	server.Static("/public/js/", "./public/js/")
	server.Static("/public/fonts/", "./public/fonts/")
	server.Static("/public/img/", "./public/img/")

	//様々なルーティングの省略

	server.POST("/signuped", SignupedRoute)

	server.Run(":3000")
}

func loadTemplates() {
	baseTemplate := "templates/layout/_base.html"
	templates = make(map[string]*template.Template)
	templates["signuped"] = template.Must(template.ParseFiles(baseTemplate, "./signuped.html"))
}

func signupedRoute(g *gin.Context) {
	g.Request.ParseForm()
	fmt.Println(g.Request.Form["id"])
}
