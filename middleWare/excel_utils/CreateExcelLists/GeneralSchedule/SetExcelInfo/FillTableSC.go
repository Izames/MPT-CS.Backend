package SetExcelInfo

import (
	"MPT-CS/middleWare/excel_utils/GetExcelInfo"
	"MPT-CS/models"
	"github.com/xuri/excelize/v2"
)

func FillTableSC(file *excelize.File) {
	var TDs []models.TeacherDay
	TDs = GetExcelInfo.GetTDs()
	SetTDsSC(TDs, file)

}
