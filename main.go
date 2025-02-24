package main

import (
	"fmt"
	"github.com/VancleyVieira/go-rest-api/database"
	"github.com/VancleyVieira/go-rest-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando um servidor com go")
	routes.HandleRequest()
}
