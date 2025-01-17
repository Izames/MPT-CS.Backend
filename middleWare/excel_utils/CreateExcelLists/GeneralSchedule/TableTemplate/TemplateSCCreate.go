package TableTemplate

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func TemplateSCCreate(file *excelize.File, teachers []string, StartRow int) *excelize.File {

	file.SetRowHeight("Sheet1", StartRow+1, 10)
	file.SetRowHeight("Sheet1", StartRow+2, 10)
	file.SetRowHeight("Sheet1", StartRow+3, 10)
	file.SetRowHeight("Sheet1", StartRow+4, 10)

	for i := 0; i <= 5; i++ {
		file.MergeCell("Sheet1", fmt.Sprintf("%s%d", models.Columns[3+i*7], StartRow+1), fmt.Sprintf("%s%d", models.Columns[9+i*7], StartRow+2))
		for j := 0; j <= 6; j++ {
			file.MergeCell("Sheet1", fmt.Sprintf("%s%d", models.Columns[3+i*7+j], StartRow+3), fmt.Sprintf("%s%d", models.Columns[3+i*7+j], StartRow+4))
		}
	}

	for i := 0; i < len(teachers); i++ {
		_ = file.MergeCell("Sheet1", fmt.Sprintf("%s%d", models.Columns[0], (StartRow+5)+i*2), fmt.Sprintf("%s%d", models.Columns[0], (StartRow+6)+i*2))
		_ = file.MergeCell("Sheet1", fmt.Sprintf("%s%d", models.Columns[1], (StartRow+5)+i*2), fmt.Sprintf("%s%d", models.Columns[2], (StartRow+6)+i*2))

	}

	file.SetColWidth("Sheet1", models.Columns[1], models.Columns[2], 15)
	file.SetColWidth("Sheet1", models.Columns[0], models.Columns[0], 5)

	return file
}
