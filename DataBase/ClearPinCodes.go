package DataBase

import (
	"MPT-CS/models"
	"time"
)

func ClearPincodes() {
	//очищение просроченных паролей каждые 5 секунд
	for {
		DB.Where("ended < ?", time.Now()).Delete(&models.ResetPassword{})
		time.Sleep(5 * time.Second)
	}
}
