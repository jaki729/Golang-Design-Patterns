/*
Singleton Pattern:
It is a creational design pattern used when only a single insatnce of the struct should exits.
This single instance is called Singleton object.

Two ways in real time applications:
1. Singleton Pattern with eager initialization:
In eager iniitialization instance of Singleton class is created much before it is actually needed.
Monstly it is done on system startup but it is not recommended as the object is created
even though our application is not using it.
2. Sinleton Pattern with lazy initialization:
With lazy initialization we create instance only when it is needed and not when the class is loaded.
It is the tactic of delaying the creation of an object until the point at which it is needed.
Like calculating the value, other expensive process etc, until the first time it is needed.

IMP:
Singleton object can have only one instance and that ionstance should be accessible globally.
Outer packages/object should be prevented to create instance of the Singleton class.

WHY:
1. DB instance: we only want to create only one instance of the DB object and that instance should be accessible globally.
2. Logger instance: we only want to create only one instance of the Logger object and that instance should be accessible globally.
3. Configuration File: This has a performace benefit as it proivides multiple user to repeatedly access
and read the same configuration file without creating multiple instances of it.
4. Cache: Use cache as singleton object as it can have global point of reference and for all future calls to the cache object
the client application will use the in-memory object.
*/

package singleton

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{} // Mutex is used to provide thread safety for the singleton instance creation
// The config struct represents a configuration object that should only have one instance
// across the application. It is a placeholder for any configuration variables you might need.
// You can add fields to this struct as per your application's requirements.

type config struct {
	// Some Config variables
}

var counter int = 1 // counter is used to track the number of instances created
var singleConfigInstance *config

func getConfigInstance() *config { // Lazy Initialization
	if singleConfigInstance == nil {
		mutex.Lock() // Locking
		defer mutex.Unlock() // Unlocking after the critical section
		if singleConfigInstance == nil {
			fmt.Println("Creting Single Instance Now, and Counter is ", counter)
			singleConfigInstance = &config{}
			counter = counter + 1 // Incrementing the counter for each new instance creation
		} else {
			fmt.Println("Single Instance already created-1, returning that one")
		}
	} else {
		fmt.Println("Single Instance already created-2, returning the same")
	}
	return singleConfigInstance
}

func DemoSingletonPattern() {
	fmt.Println("Starting DemoSingletonPattern...")
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d requesting config instance\n", id)
			getConfigInstance()
		}(i)
	}
	wg.Wait()
	fmt.Println("DemoSingletonPattern finished.")
}