package SetExcelInfo

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func SetScheduleISC(TDs []models.TeacherDay, file *excelize.File) *excelize.File {
	rows, _ := file.GetRows("Индивидуальное")
	for j := 2; j < len(rows); j += 17 {
		RowTeacher, _ := file.GetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[1], j))
		for i := 0; i < len(TDs); i++ {
			if RowTeacher == "Преподаватель "+TDs[i].Teacher {
				err := TDs[i].Error
				class := TDs[i].Class
				color := TDs[i].Color
				den := TDs[i].Denominator
				num := TDs[i].Numerator
				lessonNum := TDs[i].LessonNum
				weekDay := TDs[i].WeekDay
				col := lessonNum + 5
				row := j + 3 + (weekDay-1)*2
				if num && !den {
					file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[col], row), class+err)
					SetCellColor(file, col, row, color)
				}
				if !num && den {
					file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[col], row+1), class+err)
					SetCellColor(file, col, row+1, color)

				}
				if num && den {
					file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[col], row), class+err)
					SetCellColor(file, col, row, color)
					file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[col], row+1), class+err)
					SetCellColor(file, col, row+1, color)
				}
			}
		}
	}
	return file
}
