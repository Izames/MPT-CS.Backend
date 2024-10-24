package CRUD

import (
	"MPT-CS/DataBase"
	"MPT-CS/models"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Create_pin(pin models.ResetPassword) (int, string) {

	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	randomNumber := rand.Intn(1000000)          // Генерируем число от 0 до 999999
	pin.Pin = fmt.Sprintf("%06d", randomNumber) // Форматируем число с ведущими нулями

	if result := DataBase.DB.Create(&pin).Error; result != nil {
		log.Println("Error of creating:", result)
		return 500, ""
	}
	return 201, pin.Pin
}
