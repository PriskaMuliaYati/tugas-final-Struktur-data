package main

import (
    "fmt"
    "math/rand"
    "time"
)

const (
    Empty = " "
    PlayerX = "X"
    PlayerO = "O"
)

var (
    board [3][3]string
)

func main() {
    initializeBoard()
    printBoard()

    rand.Seed(time.Now().UnixNano())

    for !isGameOver() {
        playerMove(PlayerX)
        if isGameOver() {
            break
        }

        computerMove(PlayerO)
        printBoard()
    }

    if isWin(PlayerX) {
        fmt.Println("Player X wins!")
    } else if isWin(PlayerO) {
        fmt.Println("Player O wins!")
    } else {
        fmt.Println("It's a draw!")
    }
}

func initializeBoard() {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            board[i][j] = Empty
        }
    }
}

func printBoard() {
    fmt.Println("-------------")
    for i := 0; i < 3; i++ {
        fmt.Print("| ")
        for j := 0; j < 3; j++ {
            fmt.Print(board[i][j] + " | ")
        }
        fmt.Println()
        fmt.Println("-------------")
    }
}

func playerMove(player string) {
    var row, col int
    for {
        fmt.Printf("Player %s's turn. Enter row (0, 1, 2) and column (0, 1, 2) separated by a space: ", player)
        _, err := fmt.Scanf("%d %d", &row, &col)
        if err != nil || row < 0 || row > 2 || col < 0 || col > 2 || board[row][col] != Empty {
            fmt.Println("Invalid move. Try again.")
            continue
        }
        board[row][col] = player
        break
    }
    printBoard()
}

func computerMove(player string) {
    var row, col int
    for {
        row = rand.Intn(3)
        col = rand.Intn(3)
        if board[row][col] == Empty {
            break
        }
    }
    fmt.Printf("Player %s's turn. Computer chose row %d and column %d\n", player, row, col)
    board[row][col] = player
}

func isGameOver() bool {
    return isWin(PlayerX) || isWin(PlayerO) || isBoardFull()
}

func isWin(player string) bool {
    for i := 0; i < 3; i++ {
        if board[i][0] == player && board[i][1] == player && board[i][2] == player {
            return true
        }
        if board[0][i] == player && board[1][i] == player && board[2][i] == player {
            return true
        }
    }
    if board[0][0] == player && board[1][1] == player && board[2][2] == player {
        return true
    }
    if board[0][2] == player && board[1][1] == player && board[2][0] == player {
        return true
    }
    return false
}

func isBoardFull() bool {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == Empty {
                return false
            }
        }
    }
    return true
}
