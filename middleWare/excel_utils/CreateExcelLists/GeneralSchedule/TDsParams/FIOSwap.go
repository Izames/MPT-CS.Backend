package TDsParams

import (
	"MPT-CS/models"
	"fmt"
	"strings"
)

func FIOSwap(TDs []models.TeacherDay) []models.TeacherDay {
	for i := 0; i < len(TDs); i++ {
		parts := strings.Split(TDs[i].Teacher, " ")

		// Проверяем, что получили инициалы и фамилию
		if len(parts) == 2 {
			initials := parts[0]
			lastName := parts[1]

			// Форматируем ФИО с фамилией впереди
			TDs[i].Teacher = lastName + " " + initials

		} else {
			fmt.Println("Неверное количество элементов в строке.")
		}
	}
	return TDs
}
