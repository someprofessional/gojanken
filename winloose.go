package main

import "fmt"

func getStats(iteration, health, wins int) string {
	s := fmt.Sprintf("Iteration : %d -- You currently have : %d health and %d wins", iteration, health, wins)
	return s
}

func winlooseConditions(health int, wins int) {
	if health <= 0 {
		lossCondition = true
		if lossCondition {
			fmt.Println("You lose...")
		}
	}
	if wins >= 3 {
		winCondition = true
		if winCondition {
			fmt.Println("You win !")
		}
	}
}
