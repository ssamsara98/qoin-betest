package src

import (
	"fmt"
	"math/rand"
)

func RollDice(players *[]Player) {
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
	for i := 0; i < len(*players); i++ {
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
	for i := 0; i < len(*players); i++ {
		player := &(*players)[i]

		increment := 1
		for {
			nextPlayer := &(*players)[(i+increment)%len(*players)]
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
	for i := 0; i < len(*players); i++ {
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
