/*
Builder Pattern:
It aims to separate the construction of complex object from its representation,
so, that the same contruction process can create diffrent representations.
It is used to construct a complex step by step and the final step will return the object.
The process of construction an object should be generic so that it can be used to create diffrent reperesentations of the same object.
3 characterictics of Builder Pattern:
1. Product: Complex object that is to be generated.
2. Builder: Object that defines all the steps that must be taken in order to correctly create product.
3. Director: The director-object controls the algorithm that generates the final product object.

WHY:
When we have to build a complex object, one real time example is DOM creation 
where you have to create plenty of nodes and attributes to get your final obejct.

Pros:
1. The parameter to construct are reduced since we divide the object creation.
2. Object is always instantiated in complete state.
Cons:
1. Number of code lines increases since we have to create multiple classes.
*/


package builder

import (
	"fmt"
)

// Product 
type house struct {
	wall string
	window string
	door string
	floor int
}

// Builder
type houseBuilder struct {
	window string
	door string
	wall string
	floor int
}

func getHouseBuilder() houseBuilder {
	return houseBuilder{}
}

func (b *houseBuilder) buildWall() {
	b.wall = "Wooden Wall"
}

func (b *houseBuilder) buildWindow() {
	b.window = "Wooden Window"
}

func (b *houseBuilder) buildDoor() {
	b.door = "Wooden Door"
}

func (b *houseBuilder) buildNumFloor() {
	b.floor = 2
}

func (b *houseBuilder) buildHouse() house {
	return house{
		wall: b.wall,
		door: b.door,
		window: b.window,
		floor: b.floor,
	}
}

// Director
type director struct {
	builder houseBuilder
}

func newDirector(b houseBuilder) *director {
	return &director{
		builder: b,
	}
}

func (b *director) buildHouse() house {
	b.builder.buildWall()
	b.builder.buildWindow()
	b.builder.buildDoor()
	b.builder.buildNumFloor()
	return b.builder.buildHouse()
}

// DemoBuilderPattern
func DemoBuilderPattern() {
	builder := getHouseBuilder()
	director := newDirector(builder)
	house := director.buildHouse()

	// Print the house details
	fmt.Println("House Details:")
	fmt.Println("Wall:", house.wall)
	fmt.Println("Window:", house.window)
	fmt.Println("Door:", house.door)
	fmt.Println("Floor:", house.floor)
	fmt.Println("House built successfully!")
}