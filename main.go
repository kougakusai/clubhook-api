package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kougakusai/clubhook-api/graph"
	"github.com/kougakusai/clubhook-api/graph/generated"

	"github.com/kougakusai/clubhook-api/db"
)

func main() {
	psql := db.Migrate(db.Connect())

	defer db.Close(psql)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello())
	e.GET("/playground", playgroundHander())
	e.POST("/graphql", gqlHander(psql))

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

func playgroundHander() echo.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/graphql")

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func gqlHander(db *gorm.DB) echo.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
