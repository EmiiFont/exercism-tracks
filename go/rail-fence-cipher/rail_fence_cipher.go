package railfence

func BuildRailTable(message string, rails int, encode bool) [][]rune {
	rows := rails
	cols := len(message)

	rail := make([][]rune, rows)
	for i := range rail {
		rail[i] = make([]rune, cols)
	}
	row := 0
	rowFlip := 1
	for col := 0; col < cols; col++ {
		if encode {
			rail[row][col] = rune(message[col])
		} else {
			rail[row][col] = rune('?')
		}
		if row == rails-1 {
			rowFlip = -1
		}
		if row == 0 {
			rowFlip = 1
		}
		row += rowFlip
	}

	return rail
}

func Encode(message string, rails int) string {
	rail := BuildRailTable(message, rails, true)
	encoded := ""
	for row := 0; row < len(rail); row++ {
		for col := 0; col < len(rail[0]); col++ {
			if rail[row][col] != 0 {
				encoded += string(rail[row][col])
			}
		}
	}

	return encoded
}

func Decode(message string, rails int) string {
	rail := BuildRailTable(message, rails, false)
	key := 0
	for row := 0; row < len(rail); row++ {
		for col := 0; col < len(rail[0]); col++ {
			if rail[row][col] == rune('?') {
				rail[row][col] = rune(message[key])
				key++
			}
		}
	}

	decoded := ""
	row := 0
	rowFlip := 1
	for col := 0; col < len(message); col++ {
		decoded += string(rail[row][col])
		if row == rails-1 {
			rowFlip = -1
		}
		if row == 0 {
			rowFlip = 1
		}
		row += rowFlip
	}
	return decoded
}
