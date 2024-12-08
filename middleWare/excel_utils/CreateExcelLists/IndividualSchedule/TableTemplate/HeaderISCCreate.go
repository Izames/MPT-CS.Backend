package TableTemplate

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func HeaderISCCreate(file *excelize.File, StartRow int, teachers []string) *excelize.File {

	for i := 0; i < len(teachers); i++ {
		file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[5], StartRow), models.Year)

		file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[3], StartRow+3), "День недели")
		file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[4], StartRow+2), "№ уч.пары")

		for i := 0; i < 6; i++ {
			file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[5+i], StartRow+2), i)
		}

		for i := 0; i < 6; i++ {
			file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[3], StartRow+4+i*2), models.Week[i])
			file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[4], StartRow+4+i*2), "числ")
			file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[4], StartRow+5+i*2), "знам")
		}

		file.SetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[1], StartRow+1), "Преподаватель "+teachers[i])

		StartRow = StartRow + 17

	}
	return file
}
