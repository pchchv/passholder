package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ping(c echo.Context) error {
	msg := "Passholder. Version 0.0.0"
	return c.String(http.StatusOK, msg)
}

func routes(e *echo.Echo) {
	e.GET("/ping", ping)
}

func server() {
	e := echo.New()
	routes(e)
	log.Fatal(e.Start(envURL))
}
