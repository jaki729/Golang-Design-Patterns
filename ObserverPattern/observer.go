/*
Observer Pattern:
1. Its the behavioral design pattern that allows some objects to notify other objects about changes in their state.
2. Its a publish-subscribe design pattern,  this pattern is widely used in event handling systems.
3. Event handling systems: These are asynchronous systems where message communications are sent from producers to subscribers.
4. It is used when there is a one-to-many relationship between objects such that when one object changes state, 
all its dependents are notified and updated automatically.

Pros and Cons:
Pros:
1. Publisher doesn't need to know implementation details of Subscriber. It just knows it has to call ReactToPublisherMsg.
2. Subscriber can be added or removed at runtime without affecting the Publisher.
3. Publisher and Subscriber can be reused across domains (e.g., logging, events, UI updates, etc.)

Cons:
1. subs.(subscriber) in this code breaks abstraction and is unsafe. 
If Subscriber is an interface, you shouldn’t assume the concrete type unless absolutely necessary.
2. If subscribers are not unregistered properly (e.g., via Unregister method), it may lead to memory retention in long-running systems.
3. Notification chains can become difficult to trace, especially when there are many subscribers or nested notifications.

*/


package observer

import (
	"fmt"
)
// Publisher and Subscriber interfaces
type Publisher interface {
	Register(subscriber Subscriber)
	NotifyAll(msg string)
}

type Subscriber interface {
	ReactToPublisherMsg(msg string)
}

// Publisher
type publisher struct {
	subscriberList []Subscriber
}
// GetNewPublisher creates a new publisher instance
// It initializes the subscriberList to an empty slice
func GetNewPublisher() publisher {
	return publisher{subscriberList: make([]Subscriber, 0)}
}
// Register method allows a subscriber to register with the publisher
// It appends the subscriber to the subscriberList
func (pub *publisher) Register(subs Subscriber) {
	pub.subscriberList = append(pub.subscriberList, subs)
}

// NotifyAll method notifies all registered subscribers with a message
// It iterates through the subscriberList and calls ReactToPublisherMsg on each subscriber
// It uses type assertion to access the concrete type of the subscriber
// This allows the publisher to send messages to all subscribers without needing to know their specific types
// This is useful in scenarios where subscribers may have different implementations of the Subscriber interface
func (pub publisher) NotifyAll(msg string) {
	for _, subs := range pub.subscriberList {
		fmt.Println("Publisher notifying Subscriber with Id ", subs.(subscriber).subscriberId) // Type Assertion
		// Type assertion is used to access the concrete type of the interface
		// Here, we are asserting that subs is of type subscriber to access its subscriberId field
		subs.ReactToPublisherMsg(msg)
	}
}

// Subscriber
type subscriber struct {
	subscriberId string
}
// GetNewSubscriber creates a new subscriber instance
// It initializes the subscriberId to the provided Id
func GetNewSubscriber(Id string) subscriber {
	return subscriber{subscriberId: Id}
}
// ReactToPublisherMsg method is called by the publisher to notify the subscriber with a message
// It prints the message and the subscriberId to the console
// This method is part of the Subscriber interface and must be implemented by any type that wants to be a subscriber
// It allows the subscriber to react to messages sent by the publisher
// This is useful in scenarios where subscribers need to perform specific actions based on the messages they receive
func (s subscriber) ReactToPublisherMsg(msg string) {
	fmt.Println("Subscriber Recieved message", msg, "for subscriber Id ", s.subscriberId)
}

func DemoObserverPattern() {
	pub := GetNewPublisher()
	subs1 := GetNewSubscriber("1")
	subs2 := GetNewSubscriber("2")
	pub.Register(subs1)
	pub.Register(subs2)
	pub.NotifyAll("Thanks for notifying subscriber")
	fmt.Println("Observer Pattern Demo Completed")
}