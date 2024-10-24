package ResetPassword

import (
	"MPT-CS/CRUD"
	"MPT-CS/middleWare"
	"MPT-CS/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SavePinCode(c *gin.Context) {
	var err error
	var input models.PinCodeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !middleWare.IsValidEmail(input.Usermail) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат email"})
		return
	}

	pinCode := models.ResetPassword{
		Usermail: input.Usermail,
	}

	savedPin, pin := CRUD.Create_pin(pinCode)

	if savedPin == 201 {
		err = SendPinCode(pinCode.Usermail, pin)
	}
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"pin": savedPin})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"EmailError": err})

	}
}
