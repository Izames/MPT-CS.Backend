package TDsParams

import "regexp"

func EnglishCheck(Teacher string) bool {
	Check, _ := regexp.MatchString(`[А-Яа-я]\.[А-Яа-я]\. \W+\, [А-Яа-я]\.[А-Яа-я]\. \W+`, Teacher)
	return Check
}
