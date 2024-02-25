package main

import (
	"github.com/gjustoo/portfolio-go/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()
	userHandler := handler.UserHandler{}

	app.Static("/static/*,", "static")
	app.GET("/user", userHandler.HandleUserShow)
	//Move this por to a .ENV file
	app.Start(":3000")

}
