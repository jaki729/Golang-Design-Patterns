package main

import (
    "fmt"
    "GoLang-Design-Patterns/AdapterPattern"
    "GoLang-Design-Patterns/StatePattern"
    "GoLang-Design-Patterns/StrategyPattern"
    "GoLang-Design-Patterns/ObserverPattern"
    "GoLang-Design-Patterns/PrototypePattern"
    "GoLang-Design-Patterns/SingletonPattern"
)

func main() {
    // 1. Adapter Design Pattern Demo
    fmt.Println("Adapter Design Pattern Demo")
    adapter.DemoAdapterPattern()

    // 2. State Design Pattern Demo
    fmt.Println("\nState Design Pattern Demo")
    state.DemoStatePattern()

    // 3. Strategy Design Pattern Demo
    fmt.Println("\nStrategy Design Pattern Demo")
    strategy.DemoStrategyPattern()

    // 4. Observer Design Pattern Demo
    fmt.Println("\nObserver Design Pattern Demo")
    observer.DemoObserverPattern()

    // 5. Prototype Design Pattern Demo
    fmt.Println("\nPrototype Design Pattern Demo")
    prototype.DemoPrototypePattern()

    // 6. Singleton Design Pattern Demo
    fmt.Println("\nSingleton Design Pattern Demo")
    singleton.DemoSingletonPattern()
}