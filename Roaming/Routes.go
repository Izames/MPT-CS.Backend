package Roaming

import (
	"MPT-CS/middleWare"
	"MPT-CS/middleWare/ResetPassword"
	"MPT-CS/middleWare/excel_utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServeApplication() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Разрешаем запросы с вашего React-приложения
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")             // Разрешаем отправку cookies
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	//Группа маршрутов без авторизации
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/login", middleWare.Login)
	publicRoutes.POST("/register", middleWare.Register)
	publicRoutes.GET("/test", Test)
	publicRoutes.POST("/sendPin", ResetPassword.SavePinCode)
	publicRoutes.POST("/checkPin", ResetPassword.CheckPin)
	publicRoutes.POST("/ChangePassword", ResetPassword.ChangePassword)

	// Группа маршрутов с авторизацией
	privateRoutes := router.Group("/schedule")
	privateRoutes.Use(middleWare.JWTAuth())
	privateRoutes.POST("/CheckSchedule", excel_utils.CheckSchedule)

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
func Test(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}
