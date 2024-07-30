package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	cleaned := ""

	for _, c := range pt {
		if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= '0' && c <= '9' {
			cleaned += string(unicode.ToLower(c))
		}
	}
	length := len(cleaned)
	if length == 0 {
		return ""
	}

	list := []string{}
	cols, rows := CalculateColsAndRows(length)

	for i := 0; i < length; i += cols {
		if i+cols < length {
			list = append(list, cleaned[i:i+cols])
		} else {
			spacesToAdd := cols - (length - i)
			list = append(list, cleaned[i:]+strings.Repeat(" ", spacesToAdd))
		}
	}

	res := ""

	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			res += string(list[r][c])
		}
		res += " "
	}

	return res[:len(res)-1]
}

func CalculateColsAndRows(length int) (int, int) {
	rows := 0
	cols := 1
	for {
		if rows*cols >= length {
			if math.Abs(float64(cols)-float64(rows)) <= 1 {
				if rows > cols {
					rows, cols = cols, rows
				}
				return cols, rows
			}
		}

		if cols >= rows {
			rows++
		} else {
			cols++
			rows = cols
		}
	}
}
