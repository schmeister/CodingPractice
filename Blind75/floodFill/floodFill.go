package floodFill

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}

	DFS(&image, sr, sc, image[sr][sc], newColor)
	//	Helper(&image, sr, sc, image[sr][sc], newColor)

	return image
}

func DFS(image *[][]int, sr int, sc int, oldColor int, newColor int) {
	if sr < 0 || sc < 0 || sr >= len((*image)) || sc >= len((*image)[0]) {
		return
	}

	if (*image)[sr][sc] != oldColor {
		return
	}
	(*image)[sr][sc] = newColor

	DFS(image, sr-1, sc, oldColor, newColor)
	DFS(image, sr+1, sc, oldColor, newColor)
	DFS(image, sr, sc-1, oldColor, newColor)
	DFS(image, sr, sc+1, oldColor, newColor)
}
