package Roaming

import (
	"MPT-CS/middleWare"
	"MPT-CS/middleWare/ResetPassword"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServeApplication() {
	router := gin.Default()
	//Группа маршрутов без авторизации
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/login", middleWare.Login)
	publicRoutes.POST("/register", middleWare.Register)
	publicRoutes.GET("/test", Test)
	publicRoutes.POST("/sendPin", ResetPassword.SavePinCode)
	publicRoutes.POST("/checkPin", ResetPassword.CheckPin)
	publicRoutes.POST("/ChangePassword", ResetPassword.ChangePassword)

	// Группа маршрутов с авторизацией
	privateRoutes := router.Group("/test")
	privateRoutes.Use(middleWare.JWTAuth()) // Применяем мидлвар только к этой группе

	router.Run(":8080")
	fmt.Println("Server running on port 8000")
}
func Test(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}
