package DataBase

import (
	"MPT-CS/models"
	"log"
)

func SetDefault() {
	var user models.User
	//проверка есть ли дефолт юзер в системе
	err := DB.Where("email=?", "geday2@mail.ru").First(&user).Error
	if err == nil {
		log.Println("дефолтный пользователь уже есть в системе")
	} else {
		Create_user(models.User{Email: "geday2@mail.ru", Password: "123456"})
		log.Println("дефолтный пользователь создан")

	}
}
