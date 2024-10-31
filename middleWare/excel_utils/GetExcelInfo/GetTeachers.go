package GetExcelInfo

import (
	"MPT-CS/models"
)

func GetPrepods() ([][]string, error) {
	var teachers [][]string
	var teachersGroup []string
	var sheets []string
	// Открытие файла
	f := models.Prepods

	sheets = f.GetSheetList()

	for _, sheet := range sheets {
		rows, _ := f.GetRows(sheet)

		for i := 1; i < len(rows); i++ {
			teachersGroup = append(teachersGroup, rows[i][0])
		}
		teachers = append(teachers, teachersGroup)
		teachersGroup = nil
	}

	return teachers, nil
}
