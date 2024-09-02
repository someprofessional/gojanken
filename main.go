package main

import "fmt"

var winCondition = false
var lossCondition = false
var restart = false
var wins = 0
var health = 3
var iteration = 0

var currVal string

func main() {
	//	boucle de jeu

	enemySpeech()
	for {
		restart = false
		getStats(iteration, health, wins)
		winlooseConditions(health, wins)

		eChoic := enemyChoice()
		fmt.Scan(&currVal)

		choosenValue(currVal, eChoic)

		iteration++

		if winCondition || lossCondition {
			println("Do you want to play again ?")
			fmt.Scan(&currVal)
			if currVal == "yes" {
				restart = true
			} else {
				fmt.Scan("Okay, goodbye !")
			}
		}
		if restart {
			break
		}
	}

	fmt.Println("You broke free from the game !")
}
