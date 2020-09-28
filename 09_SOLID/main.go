package main

import (
	"example.com/username/solid/src/DIP"
	"example.com/username/solid/src/ISP"
	LSP "example.com/username/solid/src/LSP/1_caso_figuras"
	"example.com/username/solid/src/OCP"
	"example.com/username/solid/src/SRP"
)

func main() {
	SRP.Run()
	OCP.Run()
	LSP.Run()
	ISP.Run()
	DIP.Run()
}
