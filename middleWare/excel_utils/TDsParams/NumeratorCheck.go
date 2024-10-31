package TDsParams

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
	"regexp"
)

func NumenatorCheck(sheet string, i int, j int, n int, file *excelize.File) bool {
	Check1, _ := file.GetCellValue(sheet, fmt.Sprintf("%s%d", models.Columns[2+i*3], 10+14*j+n*2))
	Check2, _ := file.GetCellValue(sheet, fmt.Sprintf("%s%d", models.Columns[2+i*3], 11+14*j+n*2))

	Empty1, _ := regexp.MatchString(`[А-Яа-я]`, Check1)
	Empty1 = !Empty1
	Empty2, _ := regexp.MatchString(`[А-Яа-я]`, Check2)
	Empty2 = !Empty2
	match1, _ := regexp.MatchString(`\.`, Check1) // Проверка на наличие точки
	match2, _ := regexp.MatchString(`\.`, Check2) // Проверка на наличие точки

	if (match1 && match2) || (match1 && Empty2) || (match2 && Empty1) {
		return true
	} else {
		return false
	}
}
