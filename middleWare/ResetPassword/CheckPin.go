package ResetPassword

import (
	"MPT-CS/DataBase"
	"MPT-CS/middleWare"
	"MPT-CS/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CheckPin(context *gin.Context) {
	var input models.PinCodeInput

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	pin, err := FindPinByMail(input.Usermail)
	if err != nil {
		log.Println("срок действия пин кода прошел!")
		context.JSON(http.StatusBadRequest, gin.H{"error": "срок действия пин кода прошел!"})
		return
	}
	if pin.Pin != input.Pin {
		log.Println("пин коды не соответствуют")
		context.JSON(http.StatusBadRequest, gin.H{"error": "пин коды не соответствуют"})
		return
	}
	if len(input.Password) < 6 {
		log.Println("пароль должен быть длиннее 6 символов")
		context.JSON(http.StatusBadRequest, gin.H{"error": "пароль должен быть длиннее 6 символов"})
		return
	}
	user, _ := middleWare.FindUserByUsername(input.Usermail)
	user.Password = input.Password
	result := DataBase.Update_user(user)
	context.JSON(http.StatusOK, gin.H{"result": result})
}

func FindPinByMail(mail string) (models.ResetPassword, error) {
	var pin models.ResetPassword
	err := DataBase.DB.Where("usermail= ? AND ended > ?", mail, time.Now()).First(&pin).Error
	if err != nil {
		return models.ResetPassword{}, err
	}
	return pin, nil
}
