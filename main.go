package main

import "fmt"

var gameBoard [3][3]string

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

//func winDector([3][3][3]string) bool {
//	nil
//}

func main() {
	viewBoard(gameBoard)
}
