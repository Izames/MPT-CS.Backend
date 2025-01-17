package SetExcelInfo

import (
	"MPT-CS/middleWare/excel_utils/GetExcelInfo"
	"MPT-CS/models"
	"github.com/xuri/excelize/v2"
)

func FillTableISC(file *excelize.File) *excelize.File {
	var TDs []models.TeacherDay
	TDs = GetExcelInfo.GetTDs()
	file = SetScheduleISC(TDs, file)
	return file
}
