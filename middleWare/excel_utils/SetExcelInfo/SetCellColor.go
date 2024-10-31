package SetExcelInfo

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func SetCellColor(file *excelize.File, col int, row int, color string) *excelize.File {
	if color != "" {
		style, _ := file.NewStyle(&excelize.Style{
			Fill: excelize.Fill{
				Type:    "pattern",
				Color:   []string{color},
				Pattern: 1,
			},
		},
		)
		file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[col], row), fmt.Sprintf("%s%d", models.Columns[col], row), style)
	}

	return file
}
