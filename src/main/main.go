package main

import (
	"fmt"

	database "../database"
	server "../server"
)

func main() {
	if !database.Init() {
		fmt.Println("Cannot connect to the database")
	}

	if !server.HandleRequests() {
		fmt.Println("Cannot run the server")
	}
}
