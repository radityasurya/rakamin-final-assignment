package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/radityasurya/btpn-syariah-final/config"
	"github.com/radityasurya/btpn-syariah-final/controllers"
	"github.com/radityasurya/btpn-syariah-final/database"
	"github.com/radityasurya/btpn-syariah-final/router"
)

var (
	server *gin.Engine

	AuthController      controllers.AuthController
	AuthRouteController router.AuthRouteController

	UserController      controllers.UserController
	UserRouteController router.UserRouteController

	PhotoController      controllers.PhotoController
	PhotoRouteController router.PhotoRouteController
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	fmt.Println("Starting app...")
	database.ConnectDB(&config)

	AuthController = controllers.NewAuthController(database.DB)
	AuthRouteController = router.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(database.DB)
	UserRouteController = router.NewRouteUserController(UserController)

	PhotoController = controllers.NewPhotoController(database.DB)
	PhotoRouteController = router.NewRoutePhotoController(PhotoController)

	server = gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/")
	server.GET("/healthz", func(ctx *gin.Context) {
		msg := "OK!"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": msg})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PhotoRouteController.PhotoRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
