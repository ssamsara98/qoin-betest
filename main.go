package main

import (
	"flag"
	"fmt"
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
		src.RollDice(&players)

		// evaluation
		src.EvaluateDice(&players)

		isPlaying = false
		count++
	}
	fmt.Println("==================")

	// fmt.Println(players)

}
