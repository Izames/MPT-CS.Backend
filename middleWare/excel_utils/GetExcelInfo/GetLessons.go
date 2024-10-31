package GetExcelInfo

import (
	"MPT-CS/middleWare/excel_utils/TDsParams"
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func GetLessons(file *excelize.File, sheet string) []models.TeacherDay {

	var TDs []models.TeacherDay

	GroupsCount := 0
	var GroupsName []string
	for {
		if Group, _ := file.GetCellValue(sheet, fmt.Sprintf("%s%d", models.Columns[2+GroupsCount*3], 9)); Group != "" {
			GroupsName = append(GroupsName, Group)
			GroupsCount += 3
		} else {
			break
		}
	}

	for i := 0; i < GroupsCount; i++ { //колонка определенной группы
		for j := 0; j < 6; j++ { //день недели

			for n := 1; n < 7; n++ { //строчка
				Lesson, _ := file.GetCellValue(sheet, fmt.Sprintf("%s%d", models.Columns[1+i*3], 11+14*j+n*2))
				LessonNum, _ := strconv.Atoi(Lesson)
				Class, _ := file.GetCellValue(sheet, fmt.Sprintf("%s%d", models.Columns[2+i*3], 9))
				Teacher, _ := file.GetCellValue(sheet, fmt.Sprintf("%s%d", models.Columns[2+i*3], 10+14*j+n*2))
				Point, _ := file.GetCellValue(sheet, fmt.Sprintf("%s%d", models.Columns[2+i*3], 10+j*14))
				match := TDsParams.NumenatorCheck(sheet, i, j, n, file)

				if match {
					IsEnglish := TDsParams.EnglishCheck(Teacher)

					if IsEnglish {
						Teacher1EN, Teacher2EN := GetEnPrepodFromString(Teacher)
						TDs = AddTd(j, LessonNum, Class, Teacher1EN, Point, true, false, TDs)
						TDs = AddTd(j, LessonNum, Class, Teacher2EN, Point, true, false, TDs)
					} else {
						Teacher1 := GetPrepodFromString(Teacher)
						TDs = AddTd(j, LessonNum, Class, Teacher1, Point, true, false, TDs)
					}

					Teacher2, _ := file.GetCellValue(sheet, fmt.Sprintf("%s%d", models.Columns[2+i*3], 11+14*j+n*2))
					IsEnglish = TDsParams.EnglishCheck(Teacher2)
					if IsEnglish {
						Teacher1EN, Teacher2EN := GetEnPrepodFromString(Teacher2)
						TDs = AddTd(j, LessonNum, Class, Teacher1EN, Point, false, true, TDs)
						TDs = AddTd(j, LessonNum, Class, Teacher2EN, Point, false, true, TDs)
					} else {
						Teacher2 = GetPrepodFromString(Teacher2)
						TDs = AddTd(j, LessonNum, Class, Teacher2, Point, false, true, TDs)
					}

				} else {
					Teacher, _ = file.GetCellValue(sheet, fmt.Sprintf("%s%d", models.Columns[2+i*3], 11+14*j+n*2))
					IsEnglish := TDsParams.EnglishCheck(Teacher)
					if IsEnglish {
						Teacher1EN, Teacher2EN := GetEnPrepodFromString(Teacher)
						TDs = AddTd(j, LessonNum, Class, Teacher1EN, Point, true, true, TDs)
						TDs = AddTd(j, LessonNum, Class, Teacher2EN, Point, true, true, TDs)

					} else {
						TDs = AddTd(j, LessonNum, Class, Teacher, Point, true, true, TDs)
					}
				}

			}
		}
	}
	return TDs
}
