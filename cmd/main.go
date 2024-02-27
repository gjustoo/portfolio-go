package main

import (
	"os"

	"github.com/gjustoo/portfolio-go/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()
	userHandler := handler.LayoutHandler{}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	app.Static("/static/*,", "static")
	app.GET("/", userHandler.HandleLayoutShow)
	//Move this por to a .ENV file
	app.Start(":" + port)

}
