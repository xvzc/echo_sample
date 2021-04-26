package main

import (
	db "echo_ex/db_module"
	_ "fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.PUT("/todo", db.InsertTodo)
	e.GET("/todo", db.SelectTodo)
	e.PATCH("/todo", db.UpdateTodo)
	e.DELETE("/todo", db.DeleteTodo)

	e.Logger.Fatal(e.Start(":1323"))
}
