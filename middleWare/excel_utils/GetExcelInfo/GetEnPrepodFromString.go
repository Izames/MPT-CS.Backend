package GetExcelInfo

import (
	"regexp"
	"strings"
)

func GetEnPrepodFromString(Teacher string) (string, string) {
	re := regexp.MustCompile(`[А-Яа-я]\.[А-Яа-я]\. \W+\, [А-Яа-я]\.[А-Яа-я]\. \W+`) // Компиляция с флагом regexp.Unicode
	matches := re.FindString(Teacher)
	parts := strings.Split(matches, ", ")
	Teacher1 := parts[0]
	Teacher2 := parts[1]
	Teacher1 = strings.TrimRight(Teacher1, " ")
	Teacher2 = strings.TrimRight(Teacher2, " ")
	return Teacher1, Teacher2
}
