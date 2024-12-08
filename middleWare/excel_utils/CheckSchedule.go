package excel_utils

import (
	"MPT-CS/middleWare/excel_utils/CreateExcelLists/GeneralSchedule/SetExcelInfo"
	"MPT-CS/middleWare/excel_utils/GetExcelInfo"
	"MPT-CS/middleWare/excel_utils/fileProcessing"
	"MPT-CS/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CheckSchedule(c *gin.Context) {

	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Deserealization(form)

	teachers, err := GetExcelInfo.GetPrepods()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file := CreateTable(teachers)
	SetExcelInfo.FillTableSC(file)

	files := SliceISC(file)
	zipFile := fileProcessing.ZippingFiles(file, files)

	file.Close()
	c.File(zipFile)
	os.Remove("students.xlsx")
	for i := 0; i < len(models.SchedulesName); i++ {
		CloseFiles(models.SchedulesName[i])
	}
	CloseFiles(models.Prepods)
}
