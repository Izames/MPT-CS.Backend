package excel_utils

import (
	"MPT-CS/middleWare/excel_utils/CreateExcelLists/GeneralSchedule/TableTemplate"
	"MPT-CS/middleWare/excel_utils/CreateExcelLists/IndividualSchedule/SetExcelInfo"
	TableTemplate2 "MPT-CS/middleWare/excel_utils/CreateExcelLists/IndividualSchedule/TableTemplate"
	"github.com/xuri/excelize/v2"
)

func CreateTable(teachers [][]string) *excelize.File {
	file := excelize.NewFile()
	file.NewSheet("Индивидуальное")
	file = TableTemplate.CreateTableCS(teachers, file)
	file = TableTemplate2.CreateTableICS(teachers, file)
	file = SetExcelInfo.FillTableISC(file)
	return file
}
