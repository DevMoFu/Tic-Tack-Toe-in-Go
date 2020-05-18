package main

import "fmt"

var gameBoard [3][3]string
var winConditions = [2]string{"xxx", "ooo"}

func viewBoard(b [3][3]string) {
	for i := range gameBoard {
		fmt.Println(b[i])
	}
}

func fillPosition(user string, position uint) {
	u := user
	b := gameBoard[:]
	switch position {
	case 1:
		b[0][0] = u
	case 2:
		b[0][1] = u
	case 3:
		b[0][2] = u
	case 4:
		b[1][0] = u
	case 5:
		b[1][1] = u
	case 6:
		b[1][2] = u
	case 7:
		b[2][0] = u
	case 8:
		b[2][1] = u
	case 9:
		b[2][2] = u
	default:
		fmt.Println("Not a valid position")
	}

	fmt.Println(b)
}

//// Check board conditions
// TODO: Fix return
func checkRows(b [3][3]string) string {
	for i := range b {
		rowResult := fmt.Sprintf("%v%v%v", b[i], b[i+1], b[i+2])
		if rowResult == winConditions[0] || rowResult == winConditions[1] {
			return "win"
		}
	}
}

// TODO: Fix return
func checkColumns(b [3][3]string) string {
	for i := range b {
		columnResult := fmt.Sprintf("%v%v%v", b[0][i], b[1][i], b[2][i])
		if columnResult == winConditions[0] || columnResult == winConditions[1] {
			return "win"
		}
	}
}

// fix return
func checkDiagonal(b [3][3]string) string {
	i, j, k := 0, 1, 2

	diagonalRight := fmt.Sprintf("%v%v%v", b[0][i], b[j][j], b[2][k])
	diagonalLeft := fmt.Sprintf("%v%v%v", b[0][k], b[j][j], b[2][i])
	allDiagonalResults := [2]string{diagonalLeft, diagonalRight}

	for l := range allDiagonalResults {
		for m := range winConditions {
			if allDiagonalResults[i] == winConditions[m] {
				return "win"
			}
		}
	}
}

func gameProgress(b [3][3]string) string {
	// 3 across, 3 down, 3 diagonal

}

func main() {
	viewBoard(gameBoard)
}
