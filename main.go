package main

import (
	"fmt"
	"qoin-dadu/src"
)

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {

	// var N = flag.Int64("N", 2, "type your N")
	// var M = flag.Int64("M", 2, "type your M")
	// flag.Parse()

	var N int
	var M int

	fmt.Printf("Masukan jumlah pemain: ")
	_, err := fmt.Scanln(&N)
	if err != nil {
		fmt.Println("Input salah.")
		return
	}
	if N < 2 {
		fmt.Println("Jumlah pemain tidak boleh kurang dari 2")
		return
	}

	fmt.Printf("Masukan jumlah dadu: ")
	_, err = fmt.Scanln(&M)
	if err != nil {
		fmt.Println("Input salah.")
		return
	}
	if M < 1 {
		fmt.Println("Jumlah pemain tidak boleh kurang dari 1")
		return
	}

	fmt.Println()
	fmt.Printf("Pemain = %d, Dadu = %d\n", N, M)

	// players := make([]src.Player, *N)
	players := make([]src.Player, 0)

	for i := 0; i < N; i++ {
		player := src.Player{Dice: make([]int, M)}
		players = append(players, player)
	}

	// fmt.Println(players)

	isPlaying := true
	lastStand := -1
	count := 1

	fmt.Println("==================")
	for isPlaying {
		fmt.Printf("Giliran %d lempar dadu:\n", count)

		// players rolling the dice
		src.RollDice(&players)

		// evaluation
		src.EvaluateDice(&players)

		fmt.Println("==================")

		count++
		isGameOver, lastStandPlayer := src.CheckGameOver(&players)
		isPlaying = !isGameOver
		lastStand = lastStandPlayer
	}

	mostPointsPlayers := src.MostPointsPlayers(&players)
	winner := ""
	for i := 0; i < len(*mostPointsPlayers); i++ {
		winner += fmt.Sprintf("#%d", (*mostPointsPlayers)[i])

		if i < len(*mostPointsPlayers)-1 {
			winner += ", "
		}
	}

	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", lastStand)
	fmt.Printf("Game dimenangkan oleh pemain %s karena memiliki poin lebih banyak dari pemain lainnya.", winner)
}
