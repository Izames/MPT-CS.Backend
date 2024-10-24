package ResetPassword

import (
	"MPT-CS/DataBase"
	"MPT-CS/models"
	"github.com/gin-gonic/gin"
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
		context.JSON(http.StatusBadRequest, gin.H{"error": "срок действия пин кода прошел!"})
		return
	}

	if pin.Pin != input.Pin {
		context.JSON(http.StatusBadRequest, gin.H{"error": "пин коды не соответствуют"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"result": true})
}

func FindPinByMail(mail string) (models.ResetPassword, error) {
	var pin models.ResetPassword
	err := DataBase.DB.Where("usermail= ? AND ended > ?", mail, time.Now()).First(&pin).Error
	if err != nil {
		return models.ResetPassword{}, err
	}
	return pin, nil
}
