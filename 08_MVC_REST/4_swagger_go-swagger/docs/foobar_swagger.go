package docs

import "github.com/spsa/swagger-automation/api"

// swagger:route POST /foo otros-tag idOfFoobarEndpoint
// Inicio de sesión por usuario y contraseña
// responses:
//   200: foobarResponse

// Respuesta cuando el inicio de sesión es satisfactorio, en response
// swagger:response foobarResponse
type foobarResponseWrapper struct {
	// in:body
	Body api.FooBarResponse
}

// swagger:parameters idOfFoobarEndpoint
type foobarParamsWrapper struct {
	// Parámetros para iniciar sesión, en request body
	// in:body
	Body api.FooBarRequest
}

// swagger:route GET /info/{id} seguridad-tag xyz_1
// Información de usuario
// responses:
//   200: usuario

// swagger:parameters xyz_1
type usuarioRequest struct {
	// in:path
	ID string `json:"id"`
}

// Respuesta ok
// swagger:response usuario
type usuarioWrapper struct {
	// in:path
	ID string `json:"id"`
	// in:body
	Body api.Usuario
}

// swagger:route POST /login seguridad-tag xyz_2
// Inicio de sesión
// responses:
//   200: usuario

// swagger:parameters xyz_2
type autentificacionWrapper struct {
	// in:body
	Body api.Autentificacion
}
