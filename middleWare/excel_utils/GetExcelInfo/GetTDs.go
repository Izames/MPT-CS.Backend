package GetExcelInfo

import (
	TDsParams2 "MPT-CS/middleWare/excel_utils/CreateExcelLists/GeneralSchedule/TDsParams"
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
	TDs = TDsParams2.FilterTDs(TDs)
	TDs = TDsParams2.SetTDsColors(TDs)
	TDs = TDsParams2.CheckCoincidens(TDs)
	TDs = TDsParams2.FIOSwap(TDs)
	return TDs
}
