package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Usuario struct {
	Id        string    `json:"id"`
	Nombres   string    `json:"nombres"`
	Apellidos string    `json:"apellidos`
	Creado    time.Time `json:creado`
}

var usuarioStore = make(map[string]Usuario)
var id int

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var usuarios []Usuario
	for _, v := range usuarioStore {
		usuarios = append(usuarios, v)
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(usuarios)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		panic(err)
	}

	usuario.Creado = time.Now()
	id++

	k := strconv.Itoa(id)
	usuario.Id = k
	usuarioStore[k] = usuario

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	j, err := json.Marshal(usuario)
	if err != nil {
		panic(err)
	}

	w.Write(j)
}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// usuario := usuarioStore[id]
	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		panic(err)
	}

	if usuarioOld, ok := usuarioStore[id]; ok {
		usuario.Id = id
		usuario.Creado = usuarioOld.Creado
		delete(usuarioStore, id)
		usuarioStore[id] = usuario
	} else {
		log.Printf("No se encontró el id %s", id)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := usuarioStore[id]; ok {
		delete(usuarioStore, id)
	} else {
		log.Printf("No se encontró el usuario con id %s", id)
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/users", GetUsers).Methods("GET").Queries("q", "{q}")
	r.HandleFunc("/api/users", PostUsers).Methods("POST")
	r.HandleFunc("/api/users/{id}", PutUsers).Methods("PUT")
	r.HandleFunc("/api/users/{id}", DeleteUsers).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Escuchando...")
	server.ListenAndServe()
}
