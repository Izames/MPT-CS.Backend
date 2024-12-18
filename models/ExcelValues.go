package models

import (
	excelize "github.com/xuri/excelize/v2"
)

var Columns = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN",
	"AO", "AP", "AQ", "AR", "AS"}

var Week = []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"}

var TableNames = []string{"ЦМК Гуманитарная", "ЦМК Иностранных языков", "ЦМК Математическая", "ЦМК Экономики и права",
	"ЦМК Физической культуры", "ЦМК Биология, химия", "ЦМК Общепрофессиональных дисциплин (программное обеспечение)",
	"ЦМК Спецдисциплин 09.02.01", "ЦМК Спецдисциплин 09.02.07 ИС,ВД,БД, 10.02.05", "ЦМК Спецдисциплин  09.02.07 П,Т",
	"ЦМК Спецдисциплин 09.02.06", "ЦМК Спецдисциплин 09.02.05", "ЦМК Профессиональных модулей 40.02.01",
	"ЦМК Профессиональных модулей", "Вакансии"}

var SchedulesName []*excelize.File

var Prepods *excelize.File
var PlaceColorsDefault = []Place{
	{
		Place: "НАХИМОВСКИЙ",
		Color: "#d7d7d7",
	},
	{
		Place: "ДИСТАНЦИОННОЕ",
		Color: "#42AAFF",
	},
}
var Year string
var PlaceColors []Place
