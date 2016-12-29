// Package pascal implements a function for calculating pascals triangles.
// http://mathworld.wolfram.com/PascalsTriangle.html
package pascal

// Triangle returns the nth pascals triangle.
func Triangle(n int) (triangle [][]int) {
	for i := 0; i < n; i++ {
		triangle = append(triangle, []int{})
		for j := 0; j <= i; j++ {
			value := 0

			if j <= i-1 && i-1 >= 0 {
				value += triangle[i-1][j]
			}

			if j-1 >= 0 && i-1 >= 0 {
				value += triangle[i-1][j-1]
			}

			if value == 0 {
				value = 1
			}

			triangle[i] = append(triangle[i], value)
		}
	}

	return triangle
}
