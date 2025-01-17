package fileProcessing

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"path/filepath"
)

func ZippingFiles(file *excelize.File, files []*excelize.File) string {
	//сжатие в зип
	tempDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	individualsDir := filepath.Join(tempDir, "individuals")
	err = os.Mkdir(individualsDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(files); i++ {
		name, _ := files[i].GetCellValue("Sheet1", fmt.Sprintf("%s%d", models.Columns[1], 2))
		filePath := filepath.Join(individualsDir, name+".xlsx")
		if err := files[i].SaveAs(filePath); err != nil {
			log.Fatal(err)
		}
	}
	filePath := filepath.Join(tempDir, "students.xlsx") //Construct the full path
	if err := file.SaveAs(filePath); err != nil {
		log.Fatal(err)
	}
	return createZip(tempDir)

}
