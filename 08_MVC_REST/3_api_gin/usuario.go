package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type IUsuario interface {
	GetUsers(c *gin.Context)
	PostUsers(c *gin.Context)
	PutUsers(c *gin.Context)
	DeleteUsers(c *gin.Context)
}

type Usuario struct {
	Id        string    `json:"id"`
	Nombres   string    `json:"nombres"`
	Apellidos string    `json:"apellidos`
	Creado    time.Time `json:creado`
}

var usuarioStore = make(map[string]Usuario)
var id int

func NewUsuario() Usuario {
	return Usuario{}
}

//GetUsers listado
func (u Usuario) GetUsers(c *gin.Context) {
	texto := c.Query("q")
	fmt.Println(texto)
	var usuarios []Usuario
	for _, v := range usuarioStore {
		usuarios = append(usuarios, v)
	}
	c.JSON(http.StatusOK, usuarios)
}

func (u Usuario) PostUsers(c *gin.Context) {
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

func (u Usuario) PutUsers(c *gin.Context) {
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

func (u Usuario) DeleteUsers(c *gin.Context) {
	id := c.Param("id")

	if _, ok := usuarioStore[id]; ok {
		delete(usuarioStore, id)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "title": fmt.Sprintf("No se encontró el usuario con id %s", id)})
	}
	c.Status(http.StatusNoContent)
	return
}
