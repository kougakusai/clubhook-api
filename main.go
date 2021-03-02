package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/kougakusaiHPTeam/clubhook-api/db"
)

func main() {
	db.Migrate(db.Connect())

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello())

	err := e.Start(getListeningPort())
	if err != nil {
		log.Fatalln(err)
	}
}

func getListeningPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":8000"
}

func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from clubhook-api.")
	}
}
