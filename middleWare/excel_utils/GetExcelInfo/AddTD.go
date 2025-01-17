package GetExcelInfo

import "MPT-CS/models"

func AddTd(j int, LessonNum int, Class string, Teacher string, Point string, Numenator bool, Denominator bool, TDs []models.TeacherDay) []models.TeacherDay {
	TDs = append(TDs, models.TeacherDay{

		WeekDay:     j + 1,
		LessonNum:   LessonNum,
		Class:       Class,
		Teacher:     Teacher,
		Point:       Point,
		Numerator:   Numenator,
		Denominator: Denominator,
	})
	return TDs
}
