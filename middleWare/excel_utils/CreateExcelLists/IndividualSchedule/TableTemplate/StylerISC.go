package TableTemplate

import (
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func StylerISC(file *excelize.File, teachers []string, StartRow int) *excelize.File {
	BD, _ := file.NewStyle(&excelize.Style{ //center text bold border DiagonalDown font 8
		Border: []excelize.Border{
			{
				Type:  "diagonalDown",
				Color: "000000",
				Style: 1,
			},
		},
	})
	BDTL, _ := file.NewStyle(&excelize.Style{ //center text bold border DiagonalDown font 8
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
				Type:  "diagonalDown",
				Color: "000000",
				Style: 1,
			},
		},
	})
	BFull, _ := file.NewStyle(&excelize.Style{ // Bold Borders
		Border: []excelize.Border{
			{
				Type:  "right",
				Color: "000000",
				Style: 1,
			},
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
				Type:  "bottom",
				Color: "000000",
				Style: 1,
			},
		},
	})
	BT, _ := file.NewStyle(&excelize.Style{ // Bold Borders
		Border: []excelize.Border{
			{
				Type:  "top",
				Color: "000000",
				Style: 1,
			},
		},
	})
	BL, _ := file.NewStyle(&excelize.Style{ // Bold Borders
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "000000",
				Style: 1,
			},
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
	FontUnderline, _ := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Underline: "single",
			Bold:      true,
		},
	})
	CenterText, _ := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	for i := 0; i < len(teachers); i++ {
		file.SetCellStyle("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[5], StartRow), fmt.Sprintf("%s%d", models.Columns[5], StartRow), FontUnderline)
		file.SetCellStyle("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[1], StartRow+1), fmt.Sprintf("%s%d", models.Columns[1], StartRow+1), FontUnderline)

		for j := 0; j < 6; j++ {
			file.SetCellStyle("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[5+j], StartRow+2), fmt.Sprintf("%s%d", models.Columns[5+j], StartRow+2), CenterText)
		}
		file.SetCellStyle("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[5], StartRow+2), fmt.Sprintf("%s%d", models.Columns[10], StartRow+3), BltrLight)
		file.SetCellStyle("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[3], StartRow+4), fmt.Sprintf("%s%d", models.Columns[10], StartRow+15), BFull)
		file.SetCellStyle("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[3], StartRow+3), fmt.Sprintf("%s%d", models.Columns[3], StartRow+3), BL)
		file.SetCellStyle("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[4], StartRow+2), fmt.Sprintf("%s%d", models.Columns[4], StartRow+2), BT)
		file.SetCellStyle("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[4], StartRow+3), fmt.Sprintf("%s%d", models.Columns[4], StartRow+3), BD)
		file.SetCellStyle("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[3], StartRow+2), fmt.Sprintf("%s%d", models.Columns[3], StartRow+2), BDTL)
		StartRow += 17
	}

	return file
}
