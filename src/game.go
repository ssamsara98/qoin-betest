package src

import (
	"fmt"
	"math/rand"
	"time"
)

func RollDice(players *[]Player) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < len(*players); i++ {
		player := &(*players)[i]

		// rolling the dice
		for j := 0; j < len(player.Dice); j++ {
			result := rand.Intn(6) + 1
			// fmt.Println(result)
			player.Dice[j] = result
		}
	}

	// print rolled dice
	PrintPlayersStats(players)
}

func EvaluateDice(players *[]Player) {
	playersLen := len(*players)

	for i := 0; i < playersLen; i++ {
		player := &(*players)[i]

		// count point and eliminating dice
		remainingDice := []int{}
		tempDice := []int{}
		for j := 0; j < len(player.Dice); j++ {
			dice := player.Dice[j]
			// fmt.Println(dice)
			if dice == 6 {
				player.Point++
			} else if dice == 1 {
				tempDice = append(tempDice, dice)
			} else {
				remainingDice = append(remainingDice, dice)
			}
		}
		player.Dice = remainingDice
		player.TempDice = tempDice
	}

	// passing the temp dice
	for i := 0; i < playersLen; i++ {
		player := &(*players)[i]

		increment := 1
		for {
			nextPlayerIdx := (i + increment) % playersLen
			if nextPlayerIdx == i {
				break
			}
			nextPlayer := &(*players)[nextPlayerIdx]
			if len(nextPlayer.Dice) == 0 {
				increment++
				continue
			}

			nextPlayer.Dice = append(nextPlayer.Dice, player.TempDice...)
			break
		}
	}

	// print after evaluation
	fmt.Println("Setelah evaluasi:")
	PrintPlayersStats(players)
}

func PrintPlayersStats(players *[]Player) {
	playersLen := len(*players)
	for i := 0; i < playersLen; i++ {
		player := (*players)[i]

		dice := ""
		if len(player.Dice) == 0 {
			dice = "_ (Berhenti bermain karena tidak memiliki dadu)"
		} else {
			for j := 0; j < len(player.Dice); j++ {
				dice += fmt.Sprint(player.Dice[j])
				if j < len(player.Dice)-1 {
					dice += ","
				}
			}
		}

		fmt.Printf("\tPemain #%d (%d): %s\n", i+1, player.Point, dice)
	}
}

func CheckGameOver(players *[]Player) (bool, int) {
	outOfGameCount := 0
	theNumberOfPlayers := make([]int, 0)
	lastStand := -1

	playersLen := len(*players)

	for i := 0; i < playersLen; i++ {
		player := &(*players)[i]
		if len(player.Dice) == 0 {
			outOfGameCount++
		} else {
			theNumberOfPlayers = append(theNumberOfPlayers, i)
		}
	}
	// fmt.Printf("outOfGame %d\n", outOfGameCount)
	gameOver := outOfGameCount >= playersLen-1

	// find the last stand one
	if gameOver {
		lastStand = theNumberOfPlayers[0]
	}

	return gameOver, lastStand + 1
}

func MostPointsPlayers(players *[]Player) *[]int {
	highestPoint := -1
	newPlayers := make([]int, 0)

	for i := 0; i < len(*players); i++ {
		player := &(*players)[i]
		if player.Point > highestPoint {
			highestPoint = player.Point
			newPlayers = []int{i + 1}
		} else if player.Point == highestPoint {
			newPlayers = append(newPlayers, i+1)
		}
	}

	return &newPlayers
}
