package TableTemplate

import (
	"MPT-CS/middleWare/excel_utils/SetExcelInfo"
	"MPT-CS/models"
	"github.com/xuri/excelize/v2"
)

func CreateTable(teachers [][]string) *excelize.File {
	file := excelize.NewFile()

	StartRow := 7

	for i := 0; i < len(teachers); i++ {
		file = Styler(file, teachers[i], StartRow)

		file = TemplateCreate(file, teachers[i], StartRow)

		file = HeaderCreate(file, models.TableNames[i], StartRow)

		file = SetExcelInfo.FillTeachers(file, teachers[i], StartRow)

		StartRow = StartRow + 5 + len(teachers[i])*2
	}

	return file
}
