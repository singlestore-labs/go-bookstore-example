package main

import (
	"bookstore/database"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	database.InitializeDB()
}
