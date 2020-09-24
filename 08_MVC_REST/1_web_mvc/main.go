package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	//"log"
)

// Page sirve para la estructura de las páginas
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "./data/" + p.Title + ".txt"
	err := ioutil.WriteFile(filename, p.Body, 0600)
	return err
}

func loadPage(title string) (*Page, error) {
	filename := "./data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	page := &Page{Title: title, Body: body}
	return page, err
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		fmt.Fprintf(w, "<h1>%s</h1>", err)
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func (p Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, p.Title)
}

type Persona struct {
	Nombre   string
	Edad     int
	Operador string
}

const plantillaEnLinea = `
{{range .}}
	{{if .Edad}}
		Nombre: {{.Nombre}} Edad: {{.Edad}}
	{{end}}
{{end}}`

func main() {
	// http.HandleFunc("/view/", viewHandler)
	// http.ListenAndServe(":8080", nil)

	pagina := Page{
		Title: "Hola desde handle",
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hola desde login")
	})

	mux.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hola desde usuario")
	})

	mux.Handle("/mipagina", pagina)

	personas := []Persona{
		{"Alex", 30, "Claro"},
		{"Estefanny", 20, "Movistar"},
		{"José Luis", 0, "Entel"},
		{"Gonzalo", 0, "Bitel"},
	}

	//region template en cadena
	fmt.Println("******** Inicio de template en línea ********")

	templateEnLinea := template.New("plantilla-enlinea")

	templateEnLinea, err := templateEnLinea.Parse(plantillaEnLinea)
	if err != nil {
		panic(err)
	}

	err = templateEnLinea.Execute(os.Stdout, personas)
	if err != nil {
		panic(err)
	}
	//endregion

	//region template de archivos
	fmt.Println("******** Inicio de template desde archivo ********")

	templateEnLayout := template.New("personas")
	templateEnLayoutParseado, err := templateEnLayout.ParseGlob("template/*.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("======> usuarios prepago")
	err = templateEnLayoutParseado.ExecuteTemplate(os.Stdout, "prepago", personas)
	if err != nil {
		panic(err)
	}

	fmt.Println("======> usuarios postpago")
	err = templateEnLayoutParseado.ExecuteTemplate(os.Stdout, "postpago", personas)
	if err != nil {
		panic(err)
	}
	//endregion

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Escuchando peticiones")
	server.ListenAndServe()
}
