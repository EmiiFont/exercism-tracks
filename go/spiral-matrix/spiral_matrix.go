package spiralmatrix

func SpiralMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}
	x, y := 0, 0
	deltaX, deltaY := 1, 0
	for i := 1; i <= size*size; i++ {
		matrix[y][x] = i
		nextPosX, nextPosY := x+deltaX, y+deltaY
		if nextPosX >= size || nextPosX < 0 || nextPosY >= size || nextPosY < 0 || matrix[nextPosY][nextPosX] != 0 {
			// change direction
			deltaX, deltaY = -deltaY, deltaX
		}
		x, y = x+deltaX, y+deltaY
	}
	return matrix
}
