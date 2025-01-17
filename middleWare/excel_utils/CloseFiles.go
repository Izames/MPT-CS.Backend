package excel_utils

import "github.com/xuri/excelize/v2"

func CloseFiles(file *excelize.File) {
	file.Close()
}
