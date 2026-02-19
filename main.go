package main

import (
	"fmt"
	"math/rand"
)

// [9]int becuase it is a array of 9 numbers which are strings
var board = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func printBoard() {
	fmt.Printf(" %s | %s | %s \n", board[0], board[1], board[2])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", board[3], board[4], board[5])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", board[6], board[7], board[8])
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func computerMove() {
	for {
		choice := rand.Intn(9) // random num between 1-8

		if board[choice] != "X" && board[choice] != "O" {
			board[choice] = "O"
			break
		}
	}
}

func isBoardFull() bool {
	for _, spot := range board {
		if spot != "X" && spot != "O" {
			return false
		}
	}
	return true
}

func checkWinner() string {
	// check rows
	if board[0] == board[1] && board[1] == board[2] {
		return board[0]
	}
	if board[3] == board[4] && board[4] == board[5] {
		return board[3]
	}
	if board[6] == board[7] && board[7] == board[8] {
		return board[6]
	}

	// check columns
	if board[0] == board[3] && board[3] == board[6] {
		return board[0]
	}
	if board[1] == board[4] && board[4] == board[7] {
		return board[1]
	}
	if board[2] == board[5] && board[5] == board[8] {
		return board[2]
	}

	// check diagonals
	if board[0] == board[4] && board[4] == board[8] {
		return board[0]
	}
	if board[2] == board[4] && board[4] == board[6] {
		return board[2]
	}

	return ""
}

func main() {
	for {
		clearScreen()

		printBoard()

		fmt.Println("Where would you like to go? (or 0 to exit): ")
		choice := 0
		fmt.Scan(&choice)

		if choice == 0 {
			break
		}

		// make sure it doesnt crash when user says like 6455
		if choice < 1 || choice > 9 {
			fmt.Println("Invalid choice")
			continue
		}

		if board[choice-1] != "X" && board[choice-1] != "O" {
			board[choice-1] = "X"
			computerMove()
		} else {
			fmt.Println("Spot already taken")
		}

		winner := checkWinner()
		if winner != "" {
			fmt.Printf("%s wins!\n", winner)
			break
		}

		if isBoardFull() {
			fmt.Println("It's a tie!")
			break
		}

		clearScreen()
		printBoard()
	}
}
