package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id        string    `json:"id"`
	Nombres   string    `json:"nombres"`
	Apellidos string    `json:"apellidos`
	Creado    time.Time `json:creado`
}

var usuarioStore = make(map[string]Usuario)
var id int

//GetUsers listado
func GetUsers(c *gin.Context) {
	texto := c.Query("q")
	fmt.Println(texto)
	var usuarios []Usuario
	for _, v := range usuarioStore {
		usuarios = append(usuarios, v)
	}
	c.JSON(http.StatusOK, usuarios)
}

func PostUsers(c *gin.Context) {
	var usuario Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuario.Creado = time.Now()
	id++

	k := strconv.Itoa(id)
	usuario.Id = k
	usuarioStore[k] = usuario

	c.JSON(http.StatusOK, usuario)
}

func PutUsers(c *gin.Context) {
	id := c.Param("id")
	var usuario Usuario
	err := c.ShouldBind(&usuario)
	if err != nil {
		panic(err)
	}

	if usuarioOld, ok := usuarioStore[id]; ok {
		usuario.Id = id
		usuario.Creado = usuarioOld.Creado
		delete(usuarioStore, id)
		usuarioStore[id] = usuario
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "title": "Sucedió un error al actualizar"})
	}

	c.Status(http.StatusNoContent)
	return
}

func DeleteUsers(c *gin.Context) {
	id := c.Param("id")

	if _, ok := usuarioStore[id]; ok {
		delete(usuarioStore, id)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "title": fmt.Sprintf("No se encontró el usuario con id %s", id)})
	}
	c.Status(http.StatusNoContent)
	return
}

func main() {
	// r := gin.Default() //por defecto http://localhost:8080

	//Crea un enrutador sin ningún middleware de forma predeterminada
	r := gin.New()

	//El middleware Logger escribirá los registros en gin.DefaultWriter incluso si lo configura con GIN_MODE = release.
	r.Use(gin.Logger())

	//El middleware Recovery hará que se recupere de cualquier pánico y escribirá un error 500.
	r.Use(gin.Recovery())

	r.GET("/api/users", GetUsers)
	r.PUT("/api/users/:id", PutUsers)
	r.DELETE("/api/users/:id", DeleteUsers)
	r.POST("/api/users", PostUsers)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000") // el servidor estará disponible  desde "localhost:8080"
}
