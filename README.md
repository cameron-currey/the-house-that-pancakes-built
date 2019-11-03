# the-house-that-pancakes-built
Revenge of the Pancakes Programming Challenge

Explanation of codebase:

pancake/
	A Pancake is a simple interface that has pancake-y operations a pancake can perform.
		- A PancakeSmiley is an implementation of a Pancake

	A Stack is a collection of Pancakes that allows for somebody to "cut the stack" and FlipAt. It also provides helpers for accessing Last elements and to create a String for the entire stack.

restaurant/
	A WaiterXRayVision is a talented individual who can "FlipDatStack" and show off while they do it with a callback function.
		- They work from the top down on the stack, and flip the entire stack any time they encounter a pancake that is oriented differently.


main.go
	Creates a WaiterXRayVision along with a Stack and prints out its work each time it does a flip to visualize the algorithm so all pancakes are happy face up


The code is setup for extensibility. For example, if we want to add new pancake types to our stack, we can!

*_test.go files use gocheck test framework library (https://labix.org/gocheck). I find the built in Golang testing functionality to be clunky and verbose. GoCheck is a lot more expressive with Assert function calls, and traditional test framework Setup/TearDown support and shared code across tests.