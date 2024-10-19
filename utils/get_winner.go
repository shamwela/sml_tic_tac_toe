package utils

func GetWinner(gameData [3][3]int) int {
	// Find a row match
	for _, row := range gameData {
		if row[0] == row[1] && row[1] == row[2] {
			return row[0]
		}
	}

	// Find a column match
	for column := 0; column <= 2; column++ {
		if gameData[0][column] == gameData[1][column] && gameData[1][column] == gameData[2][column] {
			return gameData[0][column]
		}
	}

	// Find a diagonal match
	if gameData[0][0] == gameData[1][1] && gameData[1][1] == gameData[2][2] {
		return gameData[0][0]
	}

	if gameData[0][2] == gameData[1][1] && gameData[1][1] == gameData[2][0] {
		return gameData[0][2]
	}

	// No match
	return 0
}
