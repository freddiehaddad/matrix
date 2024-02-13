package matrix

import "sync"

// rotateFour rotates four adjacent values on each side of a square starting
// from the matrix position at row and col.
func rotateFour(wg *sync.WaitGroup, matrix [][]int, row int, col int) {
	defer wg.Done()

	width := len(matrix)

	top := []int{row, col}
	right := []int{col, width - row - 1}
	bottom := []int{right[1], width - col - 1}
	left := []int{bottom[1], row}

	temp := matrix[left[0]][left[1]]
	matrix[left[0]][left[1]] = matrix[bottom[0]][bottom[1]]
	matrix[bottom[0]][bottom[1]] = matrix[right[0]][right[1]]
	matrix[right[0]][right[1]] = matrix[top[0]][top[1]]
	matrix[top[0]][top[1]] = temp
}

// rotateSquare dispatches Go routines that rotate groups of four values from
// matrix where each value is an adjacent position on one of the sides.
// Rotation will happen for all elements making up the width of the square.
func rotateSquare(wg *sync.WaitGroup, matrix [][]int, square int, width int) {
	defer wg.Done()
	for i := 0; i < width-1; i++ {
		wg.Add(1)
		go rotateFour(wg, matrix, square, square+i)
	}
}

// Rotate uses multithreading to rotate a square matrix.  The matrix is treated
// nested squares.  For each square, a Go routine is invoked that creates
// additional Go routines that rotate four adjacent elments on each side of the
// square.  The function will block until all threads have finished.  The
// rotation happens in place.
func Rotate(matrix [][]int) {
	wg := &sync.WaitGroup{}

	numSquares := len(matrix) / 2
	width := len(matrix)

	for i := 0; i < numSquares; i++ {
		wg.Add(1)
		go rotateSquare(wg, matrix, i, width)
		width -= 2
	}

	wg.Wait()
}
