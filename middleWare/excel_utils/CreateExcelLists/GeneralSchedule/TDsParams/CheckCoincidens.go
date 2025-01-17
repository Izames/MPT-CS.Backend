package TDsParams

import "MPT-CS/models"

func CheckCoincidens(TDs []models.TeacherDay) []models.TeacherDay {
	for i := 0; i < len(TDs); i++ {
		for j := i + 1; j < len(TDs); j++ {
			if (TDs[i].Teacher == TDs[j].Teacher) && (TDs[i].WeekDay == TDs[j].WeekDay) && (TDs[i].LessonNum == TDs[j].LessonNum) &&
				((TDs[i].Numerator == TDs[j].Numerator) || (TDs[i].Denominator == TDs[j].Denominator)) {
				TDs[i].Color = "#FF0000"
				TDs[j].Color = "#FF0000"
				TDs[i].Error = "совпадение у группы " + TDs[i].Class + " и " + TDs[j].Class
				TDs[j].Error = "совпадение у группы " + TDs[i].Class + " и " + TDs[j].Class
			}
		}
	}
	//у препода не может быть разные группы в одну пару. Не может совпадать
	//ошибка группы если
	//WeekDay = Weekday,
	//LessonNum = LessonNum
	//Numerator= Numerator || Denominator = Den

	return TDs
}
