package main

import (
	"echoServise/controller"
	"echoServise/store/filestore"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	e := echo.New()

	s := filestore.NewStore()
	e.GET("/users/:id", controller.GetUser(s))
	e.GET("/users", controller.GetUsersList(s))
	e.POST("/users", controller.AddUser(s))
	e.DELETE("/users/:id", controller.RemoveUser(s))
	e.PUT("/users/:id", controller.UpdateUser(s))
	err := e.Start(":80")
	if err != nil {
		log.Fatal(err)
	}
}
