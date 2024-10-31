package excel_utils

import (
	"MPT-CS/models"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
)

func Deserealization(form *multipart.Form) {

	colors := form.Value["colors"]
	places := form.Value["places"]
	models.PlaceColors = models.PlaceColorsDefault
	for i := 0; i < len(colors); i++ {
		models.PlaceColors = append(models.PlaceColors, models.Place{
			Place: places[i],
			Color: colors[i],
		})
	}

	schedules := form.File["schedules"]
	models.SchedulesName = []*excelize.File{}
	for i := 0; i < len(schedules); i++ {
		f, _ := schedules[i].Open()
		xlsx, _ := excelize.OpenReader(f)
		models.SchedulesName = append(models.SchedulesName, xlsx)
	}

	prepods := form.File["teachers"]
	models.Prepods = nil
	f, _ := prepods[0].Open()
	models.Prepods, _ = excelize.OpenReader(f)
}
