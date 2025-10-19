package utils

import (
	"strings"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
	"github.com/yunarsuanto/base-go/constants"
)

type ExcelStyles struct {
	Header  int
	Warning int
}

func NewExcelStyles(xlsx *excelize.File) (ExcelStyles, *constants.ErrorResponse) {
	var result ExcelStyles

	header, err := xlsx.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		return result, ErrorInternalServer(err.Error())
	}
	warning, err := xlsx.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#FFCCCC"},
		},
	})
	if err != nil {
		return result, ErrorInternalServer(err.Error())
	}

	result = ExcelStyles{
		Header:  header,
		Warning: warning,
	}

	return result, nil
}

func ExcelColumnToNumber(col string) int {
	col = strings.ToUpper(col)
	result := 0
	for i := 0; i < len(col); i++ {
		result = result*26 + int(col[i]-'A'+1)
	}
	return result
}

func NumberToExcelColumn(n int) string {
	result := ""
	for n > 0 {
		n--
		ch := rune('A' + (n % 26))
		result = string(ch) + result
		n /= 26
	}
	return result
}

func GetMaxTextLength(currentLength int, text string) int {
	textLength := utf8.RuneCountInString(text)
	if textLength > currentLength {
		return textLength
	}

	return currentLength
}
