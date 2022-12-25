package main

import (
	"MALIKI-KARIM/database"
	"MALIKI-KARIM/handler"
)

func main() {
	database.StartDb()
	handler.RunServer()
}
