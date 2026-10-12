package matrix

import (
	"errors"
	"strconv"
	"strings"
)

var whitespaceChars = " \t\n\r"

// Define the Matrix type here.
type Matrix [][]int

func validateInput(s string) bool {
	splitted := strings.Split(s, "\n")
	if len(splitted) == 0 {
		return false
	}
	lenghts := []int{}
	for _, row := range splitted {
		numbers := strings.Split(row, " ")
		if len(numbers) == 1 && len(strings.Trim(numbers[0], whitespaceChars)) == 0 {
			return false
		}

		count := 0
		for _, col := range numbers {
			if len(strings.Trim(col, whitespaceChars)) == 0 {
				continue
			}
			count++
		}
		if count > 0 {
			lenghts = append(lenghts, count)
		}
	}
	if len(lenghts) == 0 {
		return false
	}
	for i := 1; i < len(lenghts); i++ {
		if lenghts[i] != lenghts[i-1] {
			return false
		}
	}
	return true
}

func New(s string) (Matrix, error) {
	if !validateInput(s) {
		return nil, errors.New("invalid input")
	}
	splitted := strings.Split(s, "\n")
	// 9 8 7\n19 18 17"
	m := make(Matrix, len(splitted))
	for i, row := range splitted {
		numbers := strings.Split(row, " ")
		m[i] = []int{}

		for _, col := range numbers {
			if len(strings.Trim(col, whitespaceChars)) == 0 {
				continue
			}
			val, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			m[i] = append(m[i], val)
		}

	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))

	for i := range cols {
		cols[i] = make([]int, len(m))
		for j := range cols[i] {
			cols[i][j] = m[j][i]
		}
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i := range m {
		rows[i] = make([]int, len(m[i]))
		copy(rows[i], m[i])
	}
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
