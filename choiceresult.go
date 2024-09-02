package main

import "fmt"

func sentence(first, second string) {
	fmt.Printf("You choose: %s. The enemy choose: %s\n", first, second)
}

func choosenValue(val, echoice string) {
	if val == "restart" {
		restart = true
	}

	if val == echoice {
		sentence(val, echoice)
		fmt.Printf("Nothing happens\n")
	}

	// player wins
	if val == "1" && echoice == "3" {
		sentence(val, echoice)
		fmt.Println("you win once !")
		wins++
	} else if val == "2" && echoice == "1" {
		sentence(val, echoice)
		fmt.Println("you win once !")
		wins++
	} else if val == "3" && echoice == "2" {
		sentence(val, echoice)
		fmt.Println("you win once !")
		wins++
	}
	// enemy wins
	if echoice == "1" && val == "3" {
		sentence(val, echoice)
		fmt.Println("you loose once !")
		health--
	} else if echoice == "2" && val == "1" {
		sentence(val, echoice)
		fmt.Println("you loose once !")
		health--
	} else if echoice == "3" && val == "2" {
		sentence(val, echoice)
		fmt.Println("you loose once !")
		health--
	}
}
