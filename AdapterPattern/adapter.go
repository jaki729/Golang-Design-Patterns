/*
Adapter Design Pattern:
This pattern is used to allow two incompatible interfaces to work together. 
It acts as a bridge between the two interfaces, allowing them to communicate without changing their existing code.

TargetInterface: used by clients
Client: This will interact with the adapter
Adapter: Wrapper which implements TargetInterface and modifies specific request available from Adaptee class
Adaptee: Object which is used by adapter to reuse existing functionality and modify for desired use

Client-->TargetInterface
TargetInterface<--Concrete Prototype
TargetInterface<--Adapter--->Adaptee

Why?:
When you do not need to change the existing object or interface rather
want to add new functionality on top of what is existing
*/

package adapter

import "fmt"

// Target
type device interface {
    charge()
}

// Concrete Prototype Implementation
type apple struct{}

func (a *apple) charge() {
    fmt.Println("Charging APPLE device")
}

// Adaptee
type android struct{}

func (a *android) chargeAndroid() {
    fmt.Println("Charging ANDROID device")
}

// Adapter
type androidAdapter struct {
    android *android
}

// The adapter implements the Target interface by adapting the Adaptee's method
func (ad *androidAdapter) charge() {
    ad.android.chargeAndroid()
}

// Client
type client struct{}

func (c *client) chargeDevice(dev device) {
    dev.charge()
}

func DemoAdapterPattern() {
    // First/Initial Requirement
    appleDevice := &apple{}
    client := &client{}
    client.chargeDevice(appleDevice)

    // Extended requirement i.e. Charge Android Device.
    androidDevice := &android{}
    androidAdapter := &androidAdapter{
        android: androidDevice,
    }
    client.chargeDevice(androidAdapter)
    fmt.Println("Adapter Pattern Demo Completed")
}
