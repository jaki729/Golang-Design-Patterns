/*
Prototype Pattern:
It is creational design pattern that lest you copy existing objects without making your code dependent on related objects.
The concept is to copy an existing object instead of creating a new one from scratch.
The existing object acts as a prototype and contains the state of the object.
3 main parts:
1. Ptototype
2. Prototype Registry
3. Client

Prototype patterns are required when object creation is time comsuming or resource intensive.
so instead of creating a new object, we can clone an existing object.
*/
package prototype
import (
	"fmt"
)

type ShapeType int // Define Custom type here

const (
	CircleType ShapeType = 1
	SquareType ShapeType = 2
)

type Shape interface {
	GetId() ShapeType // Get the Shape Id
	PrintTypeProp()   // used for printing proerty values of the Shape
	Clone() Shape     // used for getting DeepCopy
}

// Implementing Circle struct
type Circle struct {
	Id   ShapeType
	Radius int
	Daimeter int
	Circumference int
}

func NewCircle(radius, daimeter, circumference int) Circle {
	return Circle{CircleType, radius, daimeter, circumference}	
}

func (c Circle) GetId() ShapeType {
	return c.Id
}

func (c Circle) Clone() Shape {  // Prototype Method 
	// Deep Copy
	return NewCircle(c.Radius, c.Daimeter, c.Circumference)
}

func (c Circle) PrintTypeProp() {
	fmt.Printf("Circle Properties: Radius: %d, Daimeter: %d, Circumference: %d\n", c.Radius, c.Daimeter, c.Circumference)
}

type Square struct {
	Id     ShapeType
	Length int
	Width  int
}

func NewSquare(length, width int) Square {
	return Square{SquareType, length, width}
}

func (s Square) GetId() ShapeType {
	return s.Id
}

func (s Square) Clone() Shape { // Prototype Method
	// Deep Copy
	return NewSquare(s.Length, s.Width)
}

func (s Square) PrintTypeProp() {
	fmt.Printf("Square Properties: Length: %d, Width: %d\n", s.Length, s.Width)
}

var RegistryList = make(map[int]Shape) // Prototype Registry

func loadToRegistry() {
	// Load Circle to Registry
	circle := NewCircle(10, 20, 30)
	RegistryList[int(circle.GetId())] = circle

	// Load Square to Registry
	square := NewSquare(25, 25)
	RegistryList[int(square.GetId())] = square
}

func DemoPrototypePattern() {
	loadToRegistry() // Load Shapes to Registry

	square := RegistryList[int(SquareType)]
	// Type Assertion is used to check if the interface can be converted to Square type (type casted to Square)
	sq, ok := square.(Square) // Type Assertion to Square
	if ok {
		fmt.Println("Old Square Properties:")
		sq.PrintTypeProp() // Print Old Square Properties
		newSquare := sq.Clone() // Prototype Method to Clone Square
		fmt.Println("New Square Properties Cloned:")
		newSquare.PrintTypeProp() // Print New Square Properties
	}
	circle := RegistryList[int(CircleType)]
	// Type Assertion is used to check if the interface can be converted to Circle type (type casted to Circle)
	cir, ok := circle.(Circle) // Type Assertion to Circle
	if ok {
		fmt.Println("Old Circle Properties:")
		cir.PrintTypeProp() // Print Old Circle Properties
		newCircle := cir.Clone().(Circle) // Prototype Method to Clone Circle
		newCircle.Radius = 20 // Modify the cloned Circle properties
		newCircle.Daimeter = 40 // Modify the cloned Circle properties
		newCircle.Circumference = 60 // Modify the cloned Circle properties
		fmt.Println("New Circle Properties Cloned Changed Properties:")
		newCircle.PrintTypeProp() // Print New Circle Properties
	}
	fmt.Println("Prototype Pattern Demo Completed")
}