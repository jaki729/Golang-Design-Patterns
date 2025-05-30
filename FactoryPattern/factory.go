/*
Factory Pattern:
It is a creational design patten. This pattern provides a way to hide the creation logic of the instances.
Th client only interacts with a factory and tells the kind of instances that needs to be created.
Then the factory interacts with the coresponding concrete onjects and returns the correct instance back to the client.
The Factory Pattern is useful when the exact type of the object to be created is not known until runtime,
or when the creation logic is complex and should be encapsulated in a separate class.
It allows for better separation of concerns and makes the code more maintainable and extensible.


*/

package factory

import "fmt"

// App uses the Factory Pattern to create different types of cars.
type Car interface {
	getCar() string
}

type SedanType struct{
	Name string
}

func getNewSedan() *SedanType {
	return &SedanType{}
}

func (s *SedanType) getCar() string {
	return "Honda Civic"
}

type HatchBackType struct {
	Name string
}

func getNewHatchBack() *HatchBackType {
	return &HatchBackType{}
}
func (h *HatchBackType) getCar() string {
	return "Polo"
}

func DemoFactoryPattern() { // Client Code
	getCarFactory (1)
	getCarFactory (2)
}

func getCarFactory(carType int) { // Fcatory Object or the Factory Method
	var car Car
	if carType == 1 {
		car = getNewSedan()
	} else {
		car = getNewHatchBack()
	}
	carInfo := car.getCar()
	fmt.Println("Car Type: ", carInfo)
	fmt.Println("DemoFactoryPattern finished")
}