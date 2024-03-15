package tunnels

// use a 3d matrix ([][][]byte) first brackets represent elevation, second represent rows,
// third represent columns.
// to flip between 0 and 1, x = 1 - x;
// use breath first search
// elevators are optional
// move in all possible directions and add it to a matrix, also make a second matrix
// to store visited directions
// store the next possible directions in a array, remove from the array

type Position struct {
	elev int
	row  int
	col  int
}

var directions = []Position{
	{0, 0, 1},  // right
	{0, 0, -1}, // left
	{0, 1, 0},  // forwards
	{0, -1, 0}, // backwards
}

func FindShortestPath(maze [][][]byte, start Position) int {
	mem := memMatrix(len(maze), len(maze[0]), len(maze[0][0]))
	lastRow, lastCol := len(maze[0])-1, len(maze[0][0])-1
	toVisit := []Position{start}
	mem[start.elev][start.row][start.col] = true
	nodes, nextNodes := 1, 0
	moves := 0
	for len(toVisit) > 0 {
		curPos := toVisit[0]
		toVisit = toVisit[1:]
		if curPos.col == lastCol && curPos.row == lastRow-1 {
			return moves // happy path
		}
		for _, dir := range directions {
			nRow, nCol := curPos.row+dir.row, curPos.col+dir.col
			if nRow > lastRow || nRow < 0 || nCol > lastCol || nCol < 0 {
				continue
			}
			if mem[curPos.elev][nRow][nCol] {
				continue
			}
			if maze[curPos.elev][nRow][nCol] == '#' {
				continue
			}
			if maze[curPos.elev][nRow][nCol] == '$' {
				inverse := opposite(curPos.elev)
				toVisit = append(toVisit, Position{inverse, nRow, nCol})
				mem[inverse][nRow][nCol] = true
				nextNodes++
			}
			toVisit = append(toVisit, Position{curPos.elev, nRow, nCol})
			mem[curPos.elev][nRow][nCol] = true
			nextNodes++
		}
		nodes--
		if nodes == 0 {
			nodes, nextNodes = nextNodes, 0
			moves++
		}
	}
	return -1
}

func memMatrix(elv, row, col int) [][][]bool {
	out := make([][][]bool, elv)
	for idx := range out {
		out[idx] = make([][]bool, row)
		for idx2 := range out[idx] {
			out[idx][idx2] = make([]bool, col)
		}
	}
	return out
}

func opposite(x int) int {
	return 1 - x
}
