package GetExcelInfo

import "github.com/xuri/excelize/v2"

func GetSheets(file *excelize.File) []string {
	return file.GetSheetList()
}
