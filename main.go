package main

import (
	"net/http"
	"websocket/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	utils.LoadConfig()
	// db.Connect()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "change files")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
