package TableTemplate

import (
	"github.com/xuri/excelize/v2"
)

func CreateTableICS(teachers [][]string, file *excelize.File) *excelize.File {
	StartRow := 1
	for i := 0; i < len(teachers); i++ {
		file = HeaderISCCreate(file, StartRow, teachers[i])
		file = StylerISC(file, teachers[i], StartRow)
		file, StartRow = TemplateISCCreate(teachers[i], file, StartRow)
	}
	return file
}
