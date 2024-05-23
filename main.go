package main

import (
	"github.com/pedro-coelho1604/gin-api-rest/database"
	"github.com/pedro-coelho1604/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
