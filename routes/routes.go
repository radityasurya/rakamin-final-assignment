package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/radityasurya/btpn-syariah-final/controllers"
	"github.com/radityasurya/btpn-syariah-final/middleware"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

type UserRouteController struct {
	userController controllers.UserController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (ac *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.GET("/refresh", ac.authController.RefreshAccessToken)
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.POST("/register", uc.userController.RegisterUser)
	router.POST("/login", uc.userController.LoginUser)
	router.GET("/logout", middleware.DeserializeUser(), uc.userController.LogoutUser)

	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
}
