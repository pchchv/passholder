package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "Passholder. Version 0.0.1")
}

func generate(c echo.Context) error {
	length := 0
	characters := c.QueryParam("characters")
	l := c.QueryParam("length")
	if l != "" {
		length, _ = strconv.Atoi(l)
	}
	return c.String(http.StatusOK, passwordGenerator(characters, length))
}

func post(c echo.Context) error {
	return c.String(http.StatusOK, "Login added")
}

func routes(e *echo.Echo) {
	e.GET("/ping", ping)
	e.GET("/", ping)
	e.GET("/generate", generate)
	e.POST("/pass", post)
}

func server() {
	e := echo.New()
	routes(e)
	log.Fatal(e.Start(envURL))
}
