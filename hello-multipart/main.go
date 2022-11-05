package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/1995parham-learning/go101/hello-multipart/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/form", handler.Handler)

	if err := e.Start(":1378"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("echo initiation failed: %s", err.Error())
	}
}
