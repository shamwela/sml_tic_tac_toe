package main

import (
	"fmt"
	"math/rand"
	"sml_tic_tac_toe/utils"
)

func main() {
	var gameData [3][3]int

	for {
		for {
			var x int
			var y int

			for {
				fmt.Print("X value (1/2/3): ")
				fmt.Scanln(&x)
				if x != 1 && x != 2 && x != 3 {
					fmt.Println("X should be 1 or 2 or 3.")
					continue
				}
				break
			}

			for {
				fmt.Print("Y value (1/2/3): ")
				fmt.Scanln(&y)
				if y != 1 && y != 2 && y != 3 {
					fmt.Println("Y should be 1 or 2 or 3.")
					continue
				}
				break
			}

			isAlreadyTaken := gameData[x-1][y-1] != 0
			if isAlreadyTaken {
				fmt.Println("Not allowed. That place is already taken.")
				continue
			}
			gameData[x-1][y-1] = 1
			break
		}

		utils.DisplayGameBoard(gameData)
		winner := utils.GetWinner(gameData)
		if winner != 0 {
			winnerText := fmt.Sprintf("Winner is the player number %d!", winner)
			fmt.Println(winnerText)
			break
		}

		for {
			computerX := rand.Intn(3)
			computerY := rand.Intn(3)
			isAlreadyTaken := gameData[computerX][computerY] != 0

			if isAlreadyTaken {
				continue
			}

			gameData[computerX][computerY] = 2
			break
		}

		utils.DisplayGameBoard(gameData)
		finalWinner := utils.GetWinner(gameData)
		if finalWinner != 0 {
			winnerText := fmt.Sprintf("Winner is the player number %d!", finalWinner)
			fmt.Println(winnerText)
			break
		}
	}
}
