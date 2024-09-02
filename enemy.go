package main

import (
	"fmt"
	"math/rand"
)

func enemySpeech() {
	fmt.Println("Hello player, let's play for a bit.")
	fmt.Println("If you win i'll let you go back home without any problems.")
	fmt.Println("But if you loose ... Hehe, you will see ...")

}

func enemyChoice() string {

	result := rand.Intn(4-1) + 1

	switch result {
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	default:
		fmt.Println("WTF, something is wrong in the enemy choice ...")
		fmt.Println("Do they really want you to win ?")
		return "restart please"
	}
}
