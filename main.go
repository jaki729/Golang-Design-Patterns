package main

import (
    "fmt"
    "GoLang-Design-Patterns/AdapterPattern"
    "GoLang-Design-Patterns/StatePattern"
    "GoLang-Design-Patterns/StrategyPattern"
    "GoLang-Design-Patterns/ObserverPattern"
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
}