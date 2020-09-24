# Simple CRUD con GO
Este proyecto consiste en una API REST para la gestion de contactos

Instalación de GORM y Gorilla Mux

```bash
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
```
Estructura del proyecto
```
.
├── controllers/
  ├── ...
├── models/
  └── ...
├── utils/
  └── ...
└── main.go
```
Los datos de los contactos serán:
* Nombre
* Edad
* Teléfono
* Dirección
* E-mail
* Descripción