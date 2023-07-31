package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxnum := 100
	secretNum := rand.Intn(maxnum)
	// fmt.Println(secretNum)
	fmt.Printf("Tell me what you guess:")
	for {
		var inputNumer int
		_, err := fmt.Scanf("%v\n", &inputNumer)
		if err != nil {
			fmt.Println("Input error!")
			continue
		}
		if inputNumer == secretNum {
			fmt.Println("Correct! You Legend")
			break
		} else if inputNumer > secretNum {
			fmt.Println("The secret number is lesser")
		} else {
			fmt.Println("The secret number is greater")
		}
	}
}
