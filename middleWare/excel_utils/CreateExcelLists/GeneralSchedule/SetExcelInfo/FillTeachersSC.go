package SetExcelInfo

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func FillTeachersSC(file *excelize.File, teachers []string, StartRow int) *excelize.File {

	for i := 0; i < len(teachers); i++ {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[1], (StartRow+5)+i*2), teachers[i])
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[0], (StartRow+5)+i*2), i+1)
	}
	return file
}
