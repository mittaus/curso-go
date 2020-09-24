package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// FooBarRequest represents body of FooBar request.
type FooBarRequest struct {
	Foo string `json:"foo"`
	Bar []int  `json:"bar"`
}

// FooBarResponse represents body of FooBar response.
type FooBarResponse struct {
	Baz struct {
		Prop string `json:"prop"`
	} `json:"baz"`
}

// Usuario datos de usuario
type Usuario struct {
	ID        string `json:"id"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Token     string `json:"token"`
}

// Autentificacion nombre de usuario y contraseña
type Autentificacion struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// FooBarHandler handles incoming foobar requests
func FooBarHandler(ctx echo.Context) error {
	req := FooBarRequest{}
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	resp := doSthWithRequest(req)
	return ctx.JSON(http.StatusOK, resp)
}

// LoginHandler autentificación de usuarios
func LoginHandler(ctx echo.Context) error {
	req := FooBarRequest{}
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	resp := doSthWithRequest(req)
	return ctx.JSON(http.StatusOK, resp)
}

// InfoHandler información de usuario
func InfoHandler(ctx echo.Context) error {
	req := FooBarRequest{}
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	resp := doSthWithRequest(req)
	return ctx.JSON(http.StatusOK, resp)
}

func doSthWithRequest(req FooBarRequest) FooBarResponse {
	return FooBarResponse{}
}
