package main

import (
	"flag"
	"fmt"
	"math/rand"
	"qoin-dadu/src"
)

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {

	var N = flag.Int64("N", 0, "type your N")
	var M = flag.Int64("M", 0, "type your M")
	flag.Parse()

	fmt.Printf("Pemain = %d, Dadu = %d\n", *N, *M)

	// players := make([]src.Player, *N)
	players := make([]src.Player, 0)

	for i := 0; int64(i) < *N; i++ {
		player := src.Player{Dice: make([]int, *M)}
		players = append(players, player)
	}

	// fmt.Println(players)

	isPlaying := true
	count := 1

	fmt.Println("==================")
	for isPlaying {
		fmt.Printf("Giliran %d lempar dadu:\n", count)

		// players rolling the dice
		for i := 0; i < len(players); i++ {
			player := players[i]

			// rolling the dice
			for j := 0; j < len(player.Dice); j++ {
				result := rand.Intn(6) + 1
				// fmt.Println(result)
				player.Dice[j] = result
			}
		}

		// evaluation
		for i := 0; i < len(players); i++ {
			player := &players[i]

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

		isPlaying = false
		count++
	}

	fmt.Println(players)

}
