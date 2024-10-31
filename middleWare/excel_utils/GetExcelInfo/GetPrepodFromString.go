package GetExcelInfo

import (
	"regexp"
	"strings"
)

func GetPrepodFromString(Teacher string) string {
	re := regexp.MustCompile(`[А-Яа-я]\.[А-Яа-я]\. \W+`) // Компиляция с флагом regexp.Unicode
	matches := re.FindString(Teacher)
	Teacher = strings.TrimRight(matches, " ")
	return Teacher
}
