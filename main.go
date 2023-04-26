package main

import (
	"API-GO-REST/models"
	"API-GO-REST/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia 1 teste"},
		{Id: 2, Nome: "Nome2 dfjld", Historia: "Historia 2 teste"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
