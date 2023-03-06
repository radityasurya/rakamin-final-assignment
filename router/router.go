package router

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

type PhotoRouteController struct {
	photoController controllers.PhotoController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func NewRoutePhotoController(photoController controllers.PhotoController) PhotoRouteController {
	return PhotoRouteController{photoController}
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

	router.PUT("/:userId", middleware.DeserializeUser(), uc.userController.UpdateUser)
	router.GET("/:userId", middleware.DeserializeUser(), uc.userController.FindUserById)
	router.DELETE("/:userId", middleware.DeserializeUser(), uc.userController.DeleteUser)
}

func (pc *PhotoRouteController) PhotoRoute(rg *gin.RouterGroup) {
	router := rg.Group("photos")

	router.Use(middleware.DeserializeUser())

	router.POST("/", pc.photoController.CreatePhoto)
	router.GET("/", pc.photoController.FindPhotos)
	router.PUT("/:photoId", pc.photoController.UpdatePhoto)
	router.GET("/:photoId", pc.photoController.FindPhotoById)
	router.DELETE("/:photoId", pc.photoController.DeletePhoto)
}
