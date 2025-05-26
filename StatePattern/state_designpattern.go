/*
State Design Pattern

What?:
Comes under Behavioral Design Pattern
This pattern is used when you want to change the state of an object based on its interal state.
Example: A vending machine which changes its state based on the coins inserted and items selected.
3 main states:
1. Context: Interface/Implementation for the client to interact with, it contains a reference 
to the current state objects and define their current state.
2. State Interface for the onject states to implement, it contains methods that can be called by the context.
3. Concrete States: Implementations of the state interface, each representing a specific state of the context.

Context (uses) --> State(Actions())
Concrete StateA(Action() methods) and Concrete StateB(Action() methods)  (implements) --> State(Actions())


WHY?:
To make our code loosly coupled, reliable, maintinable, reduce complexity and easy to extend.
Using this pattern we make separate objects/classes for each state and define their behavior
and rely on th econtext of object to provide the behaviour implementation for the current state.
*/

package state

import "fmt"

// state
type tvState interface {
	state()
}

//concrete implementation of state
type on struct{}

func (o *on) state() { // Implementing the Behaviour of ON state
	fmt.Println("TV is ON!")
}

type off struct{}

func (o *off) state() { // Implementing the Behaviour of OFF state
	fmt.Println("TV is OFF!")
}

// Context
type stateContext struct {
	currentTvState tvState
}

func getContext() *stateContext {
	return &stateContext{
		currentTvState: &off{},
	}
}

func (sc *stateContext) setState(state tvState) {
	sc.currentTvState = state
}

func (sc *stateContext) getState() {
	sc.currentTvState.state()
}

//Client
func DemoStatePattern() {
	tvContext := getContext()
	tvContext.getState() // Initial state is OFF
	tvContext.setState(&on{}) // Change state to ON
	tvContext.getState() // Now state is ON
	tvContext.setState(&off{}) // Change state to OFF
	tvContext.getState() // Now state is OFF
	fmt.Println("State Design Pattern Demo Completed")
}