package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   int64 `bun:",pk,autoincrement"`
	Name string
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "change files")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
