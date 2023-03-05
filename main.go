package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/radityasurya/btpn-syariah-final/config"
	"github.com/radityasurya/btpn-syariah-final/database"
)

var (
	server *gin.Engine
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	fmt.Println("Starting app...")
	database.ConnectDB(&config)

	server = gin.Default()

	server.GET("/healthz", func(ctx *gin.Context) {
		msg := "OK!"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": msg})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}
