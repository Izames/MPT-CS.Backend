package TableTemplate

import (
	"MPT-CS/middleWare/excel_utils/CreateExcelLists/GeneralSchedule/SetExcelInfo"
	"MPT-CS/models"
	"github.com/xuri/excelize/v2"
)

func CreateTableCS(teachers [][]string, file *excelize.File) *excelize.File {
	StartRow := 7

	for i := 0; i < len(teachers); i++ {
		file = StylerSC(file, teachers[i], StartRow)

		file = TemplateSCCreate(file, teachers[i], StartRow)

		file = HeaderSCCreate(file, models.TableNames[i], StartRow)

		file = SetExcelInfo.FillTeachersSC(file, teachers[i], StartRow)

		StartRow = StartRow + 5 + len(teachers[i])*2
	}

	return file
}
