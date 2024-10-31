package GetExcelInfo

import (
	"MPT-CS/middleWare/excel_utils/TDsParams"
	"MPT-CS/models"
)

func GetTDs() []models.TeacherDay {
	var Sheets [][]string
	var TDs []models.TeacherDay
	for i := 0; i < len(models.SchedulesName); i++ {
		Sheets = append(Sheets, GetSheets(models.SchedulesName[i]))
	}
	for i := 0; i < len(Sheets); i++ {
		for j := 0; j < len(Sheets[i]); j++ {
			ReturnedTDs := GetLessons(models.SchedulesName[i], Sheets[i][j])
			for n := 0; n < len(ReturnedTDs); n++ {
				TDs = append(TDs, ReturnedTDs[n])
			}
		}
	}
	TDs = TDsParams.FilterTDs(TDs)
	TDs = TDsParams.SetTDsColors(TDs)
	TDs = TDsParams.CheckCoincidens(TDs)
	TDs = TDsParams.FIOSwap(TDs)
	return TDs
}
