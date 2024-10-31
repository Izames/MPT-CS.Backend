package SetExcelInfo

import (
	"MPT-CS/middleWare/excel_utils/GetExcelInfo"
	"MPT-CS/models"
	"github.com/xuri/excelize/v2"
)

func FillTable(file *excelize.File) {
	var TDs []models.TeacherDay
	TDs = GetExcelInfo.GetTDs()
	SetTDs(TDs, file)
}
