package main

import (
	"fruits_crud/server"
	"fruits_crud/server/db"
)

func main() {

	db.Db_connection()
	server.Server()

}
