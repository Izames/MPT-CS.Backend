package middleWare

import (
	"MPT-CS/CRUD"
	"MPT-CS/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(context *gin.Context) {
	var input models.UserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка формата email
	if !IsValidEmail(input.Email) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат email"})
		return
	}

	if _, err := FindUserByUsername(input.Email); err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "такой пользователь уже есть"})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	savedUser := CRUD.Create_user(user)

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}
