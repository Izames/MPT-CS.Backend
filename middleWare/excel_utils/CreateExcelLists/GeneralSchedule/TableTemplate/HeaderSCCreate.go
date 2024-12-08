package TableTemplate

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func HeaderSCCreate(file *excelize.File, tableName string, StartRow int) *excelize.File {

	file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[0], StartRow), tableName)
	file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[0], StartRow+1), "№")
	file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[0], StartRow+2), "п.п.")
	file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[2], StartRow+1), "день недели")
	file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[1], StartRow+4), "ФИО")
	file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[2], StartRow+3), "Пара")
	for i := 0; i < len(models.Week); i++ {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[3+i*7], StartRow+1), models.Week[i])
	}

	for i := 0; i < 6; i++ {
		for j := 0; j <= 6; j++ {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[3+i*7+j], StartRow+3), j)
		}
	}

	return file
}
