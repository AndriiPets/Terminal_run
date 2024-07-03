package game

type Counter struct {
	val  int
	size int
}

func (c *Counter) inc() {
	c.val++
	if c.val > c.size {
		c.val = 1
	}
}

func copyMatrix(matrix [][]rune) [][]rune {
	n := len(matrix)
	m := len(matrix[0])
	duplicate := make([][]rune, n)
	data := make([]rune, n*m)
	for i := range matrix {
		start := i * m
		end := start + m
		duplicate[i] = data[start:end:end]
		copy(duplicate[i], matrix[i])
	}

	return duplicate
}
