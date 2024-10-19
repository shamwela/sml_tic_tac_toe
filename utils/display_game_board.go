package utils

import (
	"fmt"
)

func DisplayGameBoard(gameData [3][3]int) {
	fmt.Println()

	for _, row := range gameData {
		joinedRow := fmt.Sprintf("%d, %d, %d", row[0], row[1], row[2])
		fmt.Println(joinedRow)
	}

	fmt.Println()
}
