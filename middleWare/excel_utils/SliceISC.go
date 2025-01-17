package excel_utils

import (
	"MPT-CS/middleWare/excel_utils/GetExcelInfo"
	"MPT-CS/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func SliceISC(file *excelize.File) []*excelize.File {
	TDs := GetExcelInfo.GetTDs()
	rows, _ := file.GetRows("Индивидуальное")

	result := make([]*excelize.File, 0, len(rows)/17)

	for i := 1; i < len(rows); i += 17 {
		newFile := excelize.NewFile()
		cellValues := make(map[string]string)
		for j := 1; j < 11; j++ {
			for c := 0; c < 16; c++ {
				cellAddress := fmt.Sprintf("%s%d", models.Columns[j], c+i)
				val, err := file.GetCellValue("Индивидуальное", cellAddress)
				if err == nil {
					cellValues[cellAddress] = val
				}
			}
		}

		for j := 1; j < 11; j++ {
			for c := 0; c < 16; c++ {
				cellAddress := fmt.Sprintf("%s%d", models.Columns[j], c+1)
				newFile.SetCellValue("Sheet1", cellAddress, cellValues[fmt.Sprintf("%s%d", models.Columns[j], c+i)])
			}
		}

		newFile = SetStyle(newFile, 1)

		teacher, _ := file.GetCellValue("Индивидуальное", fmt.Sprintf("%s%d", models.Columns[1], 1+i))

		for q := 0; q < len(TDs); q++ {
			if "Преподаватель "+TDs[q].Teacher == teacher {
				col := TDs[q].LessonNum + 5
				row := (TDs[q].WeekDay-1)*2 + 5
				color := TDs[q].Color
				// Упрощенная условная логика
				if TDs[q].Numerator && TDs[q].Denominator {
					newFile = SetCellColor(newFile, col, row, color)
					newFile = SetCellColor(newFile, col, row+1, color)
				}
				if TDs[q].Denominator {
					newFile = SetCellColor(newFile, col, row+1, color)
				}
				if TDs[q].Numerator {
					newFile = SetCellColor(newFile, col, row, color)
				}
			}
		}
		result = append(result, newFile)
	}
	return result
}
func SetCellColor(file *excelize.File, col int, row int, color string) *excelize.File {
	if color != "" {
		style, _ := file.NewStyle(&excelize.Style{
			Fill: excelize.Fill{
				Type:    "pattern",
				Color:   []string{color},
				Pattern: 1,
			},
			Font: &excelize.Font{
				Size: 10,
			},
		},
		)
		file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[col], row), fmt.Sprintf("%s%d", models.Columns[col], row), style)
	} else {
		style, _ := file.NewStyle(&excelize.Style{
			Font: &excelize.Font{
				Size: 10,
			},
		})
		file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[col], row), fmt.Sprintf("%s%d", models.Columns[col], row), style)
	}
	return file
}
func SetStyle(file *excelize.File, StartRow int) *excelize.File {
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
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[5], StartRow), fmt.Sprintf("%s%d", models.Columns[5], StartRow), FontUnderline)
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[1], StartRow+1), fmt.Sprintf("%s%d", models.Columns[1], StartRow+1), FontUnderline)

	for j := 0; j < 6; j++ {
		file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[5+j], StartRow+2), fmt.Sprintf("%s%d", models.Columns[5+j], StartRow+2), CenterText)
	}
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[5], StartRow+2), fmt.Sprintf("%s%d", models.Columns[10], StartRow+3), BltrLight)
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[3], StartRow+4), fmt.Sprintf("%s%d", models.Columns[10], StartRow+15), BFull)
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[3], StartRow+3), fmt.Sprintf("%s%d", models.Columns[3], StartRow+3), BL)
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[4], StartRow+2), fmt.Sprintf("%s%d", models.Columns[4], StartRow+2), BT)
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[4], StartRow+3), fmt.Sprintf("%s%d", models.Columns[4], StartRow+3), BD)
	file.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", models.Columns[3], StartRow+2), fmt.Sprintf("%s%d", models.Columns[3], StartRow+2), BDTL)

	file.MergeCell("Sheet1", fmt.Sprintf("%s%d", models.Columns[5], StartRow), fmt.Sprintf("%s%d", models.Columns[7], StartRow))

	for j := 0; j < 6; j++ {
		file.MergeCell("Sheet1", fmt.Sprintf("%s%d", models.Columns[5+j], StartRow+2), fmt.Sprintf("%s%d", models.Columns[5+j], StartRow+3))
	}

	for j := 0; j < 6; j++ {
		file.MergeCell("Sheet1", fmt.Sprintf("%s%d", models.Columns[3], StartRow+4+j*2), fmt.Sprintf("%s%d", models.Columns[3], StartRow+5+j*2))
	}

	file.SetColWidth("Sheet1", models.Columns[0], models.Columns[0], 5)
	file.SetColWidth("Sheet1", models.Columns[1], models.Columns[1], 30)
	file.SetColWidth("Sheet1", models.Columns[3], models.Columns[3], 15)
	file.SetColWidth("Sheet1", models.Columns[4], models.Columns[4], 15)

	return file
}
