package TDsParams

import "MPT-CS/models"

func SetTDsColors(TDs []models.TeacherDay) []models.TeacherDay {
	for i := 0; i < len(TDs); i++ {
		for j := 0; j < len(models.PlaceColors); j++ {
			if TDs[i].Point == models.PlaceColors[j].Place {
				TDs[i].Color = models.PlaceColors[j].Color
			}
		}

	}
	return TDs
}
