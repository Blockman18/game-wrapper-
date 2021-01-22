package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var input string
	fmt.Println("welcome to text game!")
	time.Sleep(3 * time.Second)
	fmt.Println("who would you like to play?")
	time.Sleep(10 * time.Second)
	fmt.Println("Ragnorok 100Hp sword and 50 armor!")
	fmt.Println("Heller 2 m26s 500 ammo 150 armor but very slow!")
	fmt.Scan(&input)

	answer := strings.ToLower(input)

	switch {
	case answer == "ragnorok":
		fmt.Print("You have chosen ragnorok!")
		time.Sleep(5 * time.Second)
	case answer == "heller":
		fmt.Print("You have chosen heller!")
		time.Sleep(5 * time.Second)
	default:
		fmt.Print("Invalid choice")
		time.Sleep(5 * time.Second)
	}
}
