package wordsearch

<<<<<<< Updated upstream
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	panic("Please implement the Solve function")
||||||| Stash base
=======
import (
	"errors"
	"strings"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	resultMap := map[string][2][2]int{}
	returnErro := true

	for _, word := range words {
		res, ok := leftToRight(word, puzzle)
		if ok {
			resultMap[word] = res
			returnErro = false
			continue
		}
		res, ok = reversed(word, puzzle, leftToRight)
		if ok {
			resultMap[word] = res
			returnErro = false
			continue
		}
		res, ok = topToBottom(word, puzzle)
		if ok {
			resultMap[word] = res
			returnErro = false
			continue
		}
		res, ok = reversed(word, puzzle, topToBottom)
		if ok {
			resultMap[word] = res
			returnErro = false
			continue
		}
		res, ok = diagonalFromTopLeft(word, puzzle)
		if ok {
			resultMap[word] = res
			returnErro = false
			continue
		}
		res, ok = reversed(word, puzzle, diagonalFromTopLeft)
		if ok {
			resultMap[word] = res
			returnErro = false
			continue
		}
		res, ok = diagonalFromBottomLeft(word, puzzle)
		if ok {
			resultMap[word] = res
			returnErro = false
			continue
		}
		res, ok = reversed(word, puzzle, diagonalFromBottomLeft)
		if ok {
			resultMap[word] = res
			returnErro = false
			continue
		}
		returnErro = true
	}

	if returnErro {
		return resultMap, errors.New("not found")
	}

	return resultMap, nil
}

func reversed(word string, puzzle []string, x func(string, []string) ([2][2]int, bool)) ([2][2]int, bool) {
	runStr := []rune(word)
	for i, j := 0, len(runStr)-1; i < j; i, j = i+1, j-1 {
		runStr[i], runStr[j] = runStr[j], runStr[i]
	}

	reversed := string(runStr)

	idx, ok := x(reversed, puzzle)
	if ok {
		start := idx[0]
		end := idx[1]
		return [2][2]int{{end[0], end[1]}, {start[0], start[1]}}, true
	}
	return idx, false
}

func leftToRight(word string, puzzle []string) ([2][2]int, bool) {
	for row := 0; row < len(puzzle); row++ {
		startIndex := strings.Index(puzzle[row], word)
		if startIndex == -1 {
			continue
		} else {
			endIndex := startIndex + len(word) - 1
			return [2][2]int{{startIndex, row}, {endIndex, row}}, true
		}
	}

	return [2][2]int{{-1, -1}, {-1, -1}}, false
}

func topToBottom(word string, puzzle []string) ([2][2]int, bool) {
	startIndex := [2]int{-1, -1}
	endIndex := [2]int{-1, -1}
	for col := 0; col < len(puzzle[0]); col++ {
		i := 0
		startIndex[0] = -1
		for row := 0; row < len(puzzle); row++ {
			if string(word[i]) == string(puzzle[row][col]) {
				if startIndex[0] == -1 {
					startIndex[0] = col
					startIndex[1] = row
				}
				endIndex[0] = col
				endIndex[1] = row
				i++
			} else {
				i = 0
			}

			if i == len(word) {
				return [2][2]int{startIndex, endIndex}, true
			}
		}
	}

	return [2][2]int{{-1, -1}, {-1, -1}}, false
}

func diagonalFromBottomLeft(word string, puzzle []string) ([2][2]int, bool) {
	startIndex := [2]int{-1, -1}
	endIndex := [2]int{-1, -1}
	for row := 0; row < len(puzzle); row++ {
		for col := len(puzzle[row]) - 1; col >= 0; col-- {
			if row-len(word)+1 >= 0 && col+len(word) <= len(puzzle[row]) {
				matched := true
				for i := 0; i < len(word); i++ {
					if puzzle[row-i][col+i] != word[i] {
						matched = false
						break
					}
				}
				if matched {
					startIndex = [2]int{col, row}
					endIndex = [2]int{col + len(word) - 1, row - len(word) + 1}
					return [2][2]int{startIndex, endIndex}, true
				}
			}
		}
	}

	return [2][2]int{startIndex, endIndex}, false
}

func diagonalFromTopLeft(word string, puzzle []string) ([2][2]int, bool) {
	startIndex := [2]int{-1, -1}
	endIndex := [2]int{-1, -1}
	for row := 0; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[row]); col++ {
			if row+len(word) <= len(puzzle) && col+len(word) <= len(puzzle[row]) {
				matched := true
				for i := 0; i < len(word); i++ {
					if puzzle[row+i][col+i] != word[i] {
						matched = false
						break
					}
				}
				if matched {
					startIndex = [2]int{col, row}
					endIndex = [2]int{col + len(word) - 1, row + len(word) - 1}
					return [2][2]int{startIndex, endIndex}, true
				}
			}
		}
	}

	return [2][2]int{startIndex, endIndex}, false
>>>>>>> Stashed changes
}
