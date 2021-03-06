package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! sub")
	})
	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
