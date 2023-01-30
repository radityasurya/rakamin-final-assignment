package main

import (
	"github.com/radityasurya/rakamin-final-assignment/database"
	"github.com/radityasurya/rakamin-final-assignment/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
