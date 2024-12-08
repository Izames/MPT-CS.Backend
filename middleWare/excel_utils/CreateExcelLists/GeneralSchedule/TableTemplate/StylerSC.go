package TableTemplate

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func StylerSC(file *excelize.File, teachers []string, StartRow int) *excelize.File {

	SmallText, _ := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size: 8, // Устанавливаем размер шрифта в 8 пунктов
		},
	})

	BFull, _ := file.NewStyle(&excelize.Style{ // Bold Borders
		Border: []excelize.Border{
			{
				Type:  "right",
				Color: "000000",
				Style: 2,
			},
			{
				Type:  "left",
				Color: "000000",
				Style: 2,
			},
			{
				Type:  "top",
				Color: "000000",
				Style: 2,
			},
			{
				Type:  "bottom",
				Color: "000000",
				Style: 2,
			},
		},
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "center",
		},
	})

	BR, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "right",
				Color: "000000",
				Style: 2,
			},
		},
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "center",
		},
		Font: &excelize.Font{
			Size: 8,
		},
	})

	BrBottomlight, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "right",
				Color: "000000",
				Style: 2,
			},
			{
				Type:  "bottom",
				Color: "000000",
				Style: 1,
			},
		},
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "center",
		},
		Font: &excelize.Font{
			Size: 8,
		},
	})

	BltrLight, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "000000",
				Style: 1,
			},
			{
				Type:  "top",
				Color: "000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "000000",
				Style: 1,
			},
		},
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "center",
		},
		Font: &excelize.Font{
			Size: 8,
		},
	})
	BlbrLight, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "000000",
				Style: 1,
			},
		},
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "center",
		},
		Font: &excelize.Font{
			Size: 8,
		},
	})
	BB, _ := file.NewStyle(&excelize.Style{ //center text border bottom
		Border: []excelize.Border{
			{
				Type:  "bottom",
				Color: "000000",
				Style: 2,
			},
		},
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "center",
		},
		Font: &excelize.Font{
			Size: 8,
		},
	})

	BTRF8, _ := file.NewStyle(&excelize.Style{ //center text bold border top right font 8
		Border: []excelize.Border{

			{
				Type:  "right",
				Color: "000000",
				Style: 2,
			},
			{
				Type:  "top",
				Color: "000000",
				Style: 2,
			},
		},
		Font: &excelize.Font{
			Size: 8, // Устанавливаем размер шрифта в 8 пунктов
		},
	})

	BTF8, _ := file.NewStyle(&excelize.Style{ //center text bold border top font 8
		Border: []excelize.Border{
			{
				Type:  "top",
				Color: "000000",
				Style: 2,
			},
		},
		Font: &excelize.Font{
			Size: 8, // Устанавливаем размер шрифта в 8 пунктов
		},
	})

	BRF8, _ := file.NewStyle(&excelize.Style{ //center text bold border right font 8
		Border: []excelize.Border{
			{
				Type:  "right",
				Color: "000000",
				Style: 2,
			},
		},
		Font: &excelize.Font{
			Size: 8, // Устанавливаем размер шрифта в 8 пунктов
		},
	})
	BDDF8, _ := file.NewStyle(&excelize.Style{ //center text bold border DiagonalDown font 8
		Border: []excelize.Border{
			{
				Type:  "diagonalDown",
				Color: "000000",
				Style: 2,
			},
		},
		Font: &excelize.Font{
			Size: 8, // Устанавливаем размер шрифта в 8 пунктов
		},
	})

	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[0], StartRow+1), fmt.Sprintf("%s%d", models.Columns[2], StartRow+4), SmallText)
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[3], StartRow+1), fmt.Sprintf("%s%d", models.Columns[44], StartRow+4), BFull)

	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[0], StartRow+1), fmt.Sprintf("%s%d", models.Columns[0], StartRow+1), BTRF8)

	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[1], StartRow+1), fmt.Sprintf("%s%d", models.Columns[2], StartRow+1), BTF8)
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[1], StartRow+2), fmt.Sprintf("%s%d", models.Columns[2], StartRow+2), BB)
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[0], StartRow+2), fmt.Sprintf("%s%d", models.Columns[0], StartRow+4), BRF8)

	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[1], StartRow+3), fmt.Sprintf("%s%d", models.Columns[1], StartRow+3), BDDF8)
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[2], StartRow+4), fmt.Sprintf("%s%d", models.Columns[2], StartRow+4), BDDF8)

	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[0], StartRow+5), fmt.Sprintf("%s%d", models.Columns[2], (StartRow+6)+(len(teachers)-1)*2), BFull)
	for i := 0; i < len(teachers); i++ {
		file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[3], (StartRow+5)+i*2), fmt.Sprintf("%s%d", models.Columns[44], (StartRow+5)+i*2), BltrLight)
		file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[3], (StartRow+6)+i*2), fmt.Sprintf("%s%d", models.Columns[44], (StartRow+6)+i*2), BlbrLight)
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < len(teachers); j++ {
			file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[9+i*7], StartRow+5+j*2), fmt.Sprintf("%s%d", models.Columns[9+i*7], StartRow+5+j*2), BR)
			file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[9+i*7], StartRow+6+j*2), fmt.Sprintf("%s%d", models.Columns[9+i*7], StartRow+6+j*2), BrBottomlight)
		}
	}
	return file
}
