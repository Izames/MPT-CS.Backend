package TableTemplate

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func TemplateISCCreate(teachers []string, file *excelize.File, StartRow int) (*excelize.File, int) {

	for i := 0; i < len(teachers); i++ {
		file.MergeCell("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[5], StartRow), fmt.Sprintf("%s%d", models.Columns[7], StartRow))

		for j := 0; j < 6; j++ {
			file.MergeCell("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[5+j], StartRow+2), fmt.Sprintf("%s%d", models.Columns[5+j], StartRow+3))
		}

		for j := 0; j < 6; j++ {
			file.MergeCell("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[3], StartRow+4+j*2), fmt.Sprintf("%s%d", models.Columns[3], StartRow+5+j*2))
		}

		file.SetColWidth("Индивидуальное", models.Columns[0], models.Columns[0], 5)
		file.SetColWidth("Индивидуальное", models.Columns[1], models.Columns[1], 30)
		file.SetColWidth("Индивидуальное", models.Columns[3], models.Columns[3], 15)
		file.SetColWidth("Индивидуальное", models.Columns[4], models.Columns[4], 15)

		StartRow = StartRow + 17
	}

	return file, StartRow
}
