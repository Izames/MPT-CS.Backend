package TDsParams

import "MPT-CS/models"

func FilterTDs(TDs []models.TeacherDay) []models.TeacherDay {
	var TDsFiltered []models.TeacherDay
	for i := 0; i < len(TDs); i++ {
		if TDs[i].Teacher != "" && TDs[i].Teacher != "ПРАКТИКА" {
			TDsFiltered = append(TDsFiltered, TDs[i])
		}
	}
	return TDsFiltered
}
