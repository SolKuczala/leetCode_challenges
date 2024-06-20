package main

const (
	up = iota
	down
	left
	right
)

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)    // row
	m := len(matrix[0]) // column
	result := make([]int, n*m)
	// set boundaries
	upb := 0
	rightb := n - 1
	downb := m - 1
	leftb := 0
	direction := right

	idx := 0
	// while
	for x, y := 0, 0; ; {
		// registro y muevo el indice i,j
		result[idx] = matrix[x][y] // row y column
		idx++
		// quiebre
		if upb > downb || leftb > rightb || idx == n*m {
			break
		}
		switch direction {
		case right:
			if y == rightb {
				x++
				direction = down
				// muevo la linea que ya pase
				upb++ // porque llegamos al final del row y bajamos entonces el primero minimo se bloquea
			} else {
				y++
			}
		case down:
			if x == downb {
				y--
				direction = leftb
				rightb--
			} else {
				x++
			}
		case left:
			if y == leftb {
				x--
				direction = up
				downb--
			} else {
				y--
			}
		case up:
			if x == upb {
				y++
				direction = right
				leftb++
			} else {
				x--
			}
		default:
			panic("unsupported direction")
		}
	}

	return result
}
