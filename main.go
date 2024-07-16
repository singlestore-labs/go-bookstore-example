package main

import (
	"bookstore/api"
	"bookstore/database"
)

func main() {
	database.InitializeDB()
	api.StartServer()
}
