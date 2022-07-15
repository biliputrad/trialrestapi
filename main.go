package main

import (
	"RestAPI-GETNPOST/Handler"
	"RestAPI-GETNPOST/Inisialisasi"
	"RestAPI-GETNPOST/Middleware"
	"RestAPI-GETNPOST/Repository"
	"RestAPI-GETNPOST/Service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db, _ := Inisialisasi.Initialize()
	r := initRouter(db)
	r.Run()
}
func initRouter(db *gorm.DB) *gin.Engine {
	userRepository := Repository.NewRepositoryUser(db)
	userService := Service.NewServiceUser(userRepository)
	userHandler := Handler.NewUserHandler(userService)
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/register", userHandler.RegistrasiHandler)
		api.POST("/login", userHandler.LoginHandler)
		api.GET("/listuser", Middleware.Auth(), userHandler.GetUserList)
		api.POST("/oneuser/:id", Middleware.Auth(), userHandler.GetDataUserById)
		api.PUT("/update/:id", Middleware.Auth(), userHandler.UpdateUser)
		api.DELETE("/delete/:id", Middleware.Auth(), userHandler.DeleteUser)
		api.PUT("/forgetpassword/:id", Middleware.Auth(), userHandler.ForgetPassword)
		api.GET("/pagination", Middleware.Auth(), userHandler.PaginationUser)
		api.POST("/savefiletopdf/:id", Middleware.Auth(), userHandler.ConvertDataToPDF)
	}
	return r
}
