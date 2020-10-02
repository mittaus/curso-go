package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default() //por defecto http://localhost:8080

	//Crea un enrutador sin ningún middleware de forma predeterminada
	r := gin.New()

	//El middleware Logger escribirá los registros en gin.DefaultWriter incluso si lo configura con GIN_MODE = release.
	r.Use(gin.Logger())

	//El middleware Recovery hará que se recupere de cualquier pánico y escribirá un error 500.
	r.Use(gin.Recovery())

	usuarioHandler := NewUsuario()
	r.GET("/api/users", usuarioHandler.GetUsers)
	r.PUT("/api/users/:id", usuarioHandler.PutUsers)
	r.DELETE("/api/users/:id", usuarioHandler.DeleteUsers)
	r.POST("/api/users", usuarioHandler.PostUsers)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000") // el servidor estará disponible  desde "localhost:8080"
}
