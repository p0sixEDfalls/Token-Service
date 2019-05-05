package main

import (
	database "../database"
	server "../server"
)

func main() {
	if !database.Init() {
		panic("Cannot connect to the database")
	}

	if !server.HandleRequests() {
		panic("Cannot run the server")
	}
}
