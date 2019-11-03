package restaurant

import (
	"../pancake"
)

type WaiterXRayVision struct{}

func (s *WaiterXRayVision) FlipDatStack(stack pancake.Stack, callback func(iteration int, currentStack pancake.Stack)) int {

	//The algorithm runs over each pancake, and when it notices a pancake oriented
	//differently, flips all previous panics as part of that stack.

	counter := 0
	currentOrientation := stack[0].FaceUp()
	callback(counter, stack)

	for i, pancake := range stack {
		orientation := pancake.FaceUp()
		if currentOrientation != orientation {
			currentOrientation = orientation
			counter = s.flip(i, stack, counter, callback)
		}
	}

	lastIndex, lastPancake := stack.Last()
	if !lastPancake.FaceUp() {
		counter = s.flip(lastIndex+1, stack, counter, callback)
	}

	return counter
}

func (s *WaiterXRayVision) flip(flipAt int, stack pancake.Stack, counter int, callback func(iteration int, currentOrientation pancake.Stack)) int {
	stack.FlipRange(flipAt)
	counter = counter + 1
	callback(counter, stack)
	return counter
}
