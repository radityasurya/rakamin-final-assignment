package main

import (
	"fmt"
	"log"

	"github.com/radityasurya/btpn-syariah-final/config"
	"github.com/radityasurya/btpn-syariah-final/database"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	fmt.Println("Starting app...")
	database.ConnectDB(&config)
}
