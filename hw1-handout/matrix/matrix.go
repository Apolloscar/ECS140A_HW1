package matrix

// If needed, you may define helper functions here.

// AreAdjacent returns true iff a and b are adjacent in lst.
func AreAdjacent(lst []int, a, b int) bool {
	//panic("TODO: implement this!")
	for i := 0; i < len(lst); i++ { //traverse thorough slice
		if lst[i] == a { // check if current element is equal to a
			var before, after int = i - 1, i + 1 //get adjacent inices to current element

			if before >= 0 { //check if index within range
				if lst[before] == b { //check if element is equal to b and return true if so
					return true
				}
			}

			if after < len(lst) { //check if index within range
				if lst[after] == b { //check if element is equal to b and return true if so
					return true
				}
			}

		}
	}
	return false //slice has been traversed entirely with no adjacencies so return false
}

// Transpose returns the transpose of the 2D matrix mat.
func Transpose(mat [][]int) [][]int {
	//panic("TODO: implement this!")
	if mat == nil { //transpose of nill is also nill
		return nil
	}
	var m int = len(mat) //get number of rows
	if m == 0 {          // if no rows then matrix is empty so transpose is also empty
		return make([][]int, 0)
	}
	var n int = len(mat[0])     // get number columns
	tra_mat := make([][]int, n) // make a new matrix with transposed length of rows and columns
	for i := 0; i < n; i++ {
		tra_mat[i] = make([]int, m)
	}

	for i := 0; i < n; i++ { //each element in mat goes to inerse row and column of transposed mat
		for j := 0; j < m; j++ {
			tra_mat[i][j] = mat[j][i]
		}
	}

	return tra_mat
}

// AreNeighbors returns true iff a and b are neighbors in the 2D matrix mat.
func AreNeighbors(mat [][]int, a, b int) bool {
	//panic("TODO: implement this!")
	if mat == nil { //no elements so no neighbors
		return false
	}
	var m int = len(mat) //get number of rows
	if m == 0 {          // if no rows then matrix is empty so no neighbors
		return false
	}
	var n int = len(mat[0]) // get number of columns

	for i := 0; i < m; i++ { // traverse through all elements
		for j := 0; j < n; j++ {
			if mat[i][j] == a { // found a, now check if b is a neighbor

				var row_before, row_after, column_before, column_after int = i - 1, i + 1, j - 1, j + 1

				if row_before >= 0 { // if row before exists and is b, they are neighbors
					if mat[row_before][j] == b {
						return true
					}
				}

				if row_after < m { // if row after exists and is b, they are neighbors
					if mat[row_after][j] == b {
						return true
					}
				}

				if column_before >= 0 { // if column before exists and is b, they are neighbors
					if mat[i][column_before] == b {
						return true
					}
				}

				if column_after < n { // if column after exists and is b, they are neighbors
					if mat[i][column_after] == b {
						return true
					}
				}

			}
		}
	}

	return false //went through all elements and no neighbors were found
}
