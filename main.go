package main

import (
	"./pancake"
	"./restaurant"
	"fmt"
)

func main() {

	fmt.Println("Welcome to the Infinite House of Pancakes. Watch our waiters greatness")

	stack := pancake.NewPancakeSmileyStack(
		false, false, true, true, true, true, false, true, false, false, false, true,
		false, false, false, true, false, true, true, true, true, false, false, true,
		true, false, false, true, false, true, false, true, false, false, false, true,
	)

	waiter := restaurant.WaiterXRayVision{}
	waiter.FlipDatStack(stack, func(flipNumber int, currentStack pancake.Stack) {
		fmt.Println(currentStack, " ", flipNumber)
	})
}
