package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spsa/swagger-automation/api"

	_ "github.com/spsa/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
)

func main() {
	// Echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set up basic auth with username=foo and password=bar
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(username, password string, c echo.Context) (bool, error) {
			if username == "foo" && password == "bar" {
				return true, nil
			}
			return false, nil
		},
	}))

	// Route => handler
	e.POST("/foo", api.FooBarHandler)
	e.POST("/login", api.LoginHandler)
	e.GET("/info", api.InfoHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
