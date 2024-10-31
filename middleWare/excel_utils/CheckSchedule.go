package excel_utils

import (
	"MPT-CS/middleWare/excel_utils/GetExcelInfo"
	"MPT-CS/middleWare/excel_utils/SetExcelInfo"
	"MPT-CS/middleWare/excel_utils/TableTemplate"
	"MPT-CS/models"
	"github.com/gin-gonic/gin"
	"log"
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
	file := TableTemplate.CreateTable(teachers)
	SetExcelInfo.FillTable(file)
	if err := file.SaveAs("students.xlsx"); err != nil {
		log.Fatal(err)
	}
	c.File("students.xlsx")
	os.Remove("students.xlsx")
	for i := 0; i < len(models.SchedulesName); i++ {
		CloseFiles(models.SchedulesName[i])
	}
	CloseFiles(models.Prepods)
}
