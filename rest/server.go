package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"go-community/model"
	"go-community/rest/api"
	"log"
)

var Router = gin.Default()
var Render = render.New(render.Options{
	IndentJSON: true,
})

func main() {
	model.RunDb()

	ip := "localhost"
	port := "8080"

	Router.POST("/api/sign-on", api.UserCreate)

	err := Router.Run(ip + ":" + port)
	if err != nil {
		log.Panic(err)
	}
}
