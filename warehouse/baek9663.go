package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution9663 .
func Solution9663() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	chessBoard := make([][]bool, n)
	for i := range chessBoard {
		chessBoard[i] = make([]bool, n)
	}

	fmt.Fprintln(writer, nQueen(chessBoard, 0))
}

func nQueen(board [][]bool, lineNum int) (ans int) {
	for i := 0; i < len(board); i++ {
		if isSafe(board, lineNum, i) {
			if lineNum == len(board)-1 {
				return 1
			}
			board[lineNum][i] = true
			ans += nQueen(board, lineNum+1)
			board[lineNum][i] = false
		}
	}
	return
}

//  0 1 2 3
//[[f f c f]  0 3
// [d f c f]  1 2
// [f d c d]  2 1
// [f f p f]] 3
func isSafe(board [][]bool, lineNum, pos int) bool {
	diaVar := 1
	for i := lineNum - 1; i >= 0; i-- {
		if board[i][pos] || (pos-diaVar >= 0 && board[i][pos-diaVar]) || (pos+diaVar < len(board) && board[i][pos+diaVar]) {
			return false
		}
		diaVar++
	}
	return true
}
