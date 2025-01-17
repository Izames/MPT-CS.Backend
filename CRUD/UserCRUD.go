package CRUD

import (
	"MPT-CS/DataBase"
	"MPT-CS/models"
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Create_user(user models.User) int {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal("Failed to generate salt: ", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password: ", err)
	}
	user.Password = string(hashedPassword)
	user.Salt = base64.StdEncoding.EncodeToString(salt)

	if result := DataBase.DB.Create(&user).Error; result != nil {
		log.Println("Error of creating:", result)
		return 500
	}
	return 201
}
func Update_user(user models.User) int {
	salt := make([]byte, 16)  // Выберите подходящий размер соли
	_, err := rand.Read(salt) // Заполняем соль случайными байтами
	if err != nil {
		log.Fatal("Failed to generate salt: ", err)
	}
	// Хешируем пароль с солью
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password: ", err)
	}
	// Сохраняем соль и хеш пароля в базу данных
	user.Password = string(hashedPassword)              // Хеш
	user.Salt = base64.StdEncoding.EncodeToString(salt) // Соль

	if result := DataBase.DB.Model(&user).Updates(models.User{Password: user.Password}).Error; result != nil {
		log.Println("Error of creating:", result)
		return 500
	}
	return 201
}
