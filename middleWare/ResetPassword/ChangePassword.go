package ResetPassword

import (
	"MPT-CS/DataBase"
	"MPT-CS/middleWare"
	"MPT-CS/models"
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func ChangePassword(context *gin.Context) {
	var input models.UserInput
	var hashedPassword []byte
	var UpdatedUser models.User

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := middleWare.FindUserByUsername(input.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//в будущем тут не должно этого быть
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		log.Fatal("Failed to generate salt: ", err)
	}
	// Хешируем пароль с солью
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password: ", err)
	}
	user.Password = string(hashedPassword)
	user.Salt = base64.StdEncoding.EncodeToString(salt)

	// Сохранение изменений в базу данных
	if err := DataBase.DB.Save(user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"jwt": UpdatedUser})

}
