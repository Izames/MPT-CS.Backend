package middleWare

import (
	"MPT-CS/DataBase"
	"MPT-CS/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(context *gin.Context) {
	var input models.UserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := FindUserByUsername(input.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ValidatePassword(user, input.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwt, err := GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

func FindUserByUsername(username string) (models.User, error) {
	var user models.User
	err := DataBase.DB.Where("email=?", username).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func ValidatePassword(user models.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
