package SetExcelInfo

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func SetTDsSC(TDs []models.TeacherDay, file *excelize.File) {
	rows, _ := file.GetRows("Sheet1")
	rowsCount := len(rows)
	for i := 0; i < len(TDs); i++ {
		for j := 12; j < rowsCount; j += 1 {
			RowTeacher, _ := file.GetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[1], j))
			if RowTeacher == TDs[i].Teacher {
				var row int
				var col int
				var GroupName string
				var Color string
				var Error string
				Error = TDs[i].Error
				Color = TDs[i].Color
				GroupName = TDs[i].Class
				col = 3 + (TDs[i].WeekDay-1)*7 + TDs[i].LessonNum
				if TDs[i].Numerator && !TDs[i].Denominator {
					row = j
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[col], row), GroupName+Error)
					file = SetCellColorSC(file, col, row, Color)
				}
				if !TDs[i].Numerator && TDs[i].Denominator {
					row = j + 1
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[col], row), GroupName+Error)
					file = SetCellColorSC(file, col, row, Color)
				}
				if TDs[i].Numerator && TDs[i].Denominator {
					row = j
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[col], row), GroupName+Error)
					file = SetCellColorSC(file, col, row, Color)
					row = j + 1
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[col], row), GroupName+Error)
					file = SetCellColorSC(file, col, row, Color)
				}
				break
			}
		}

	}

}
