package main

import (
    "fmt"

	"GoLang-Design-Patterns/AdapterPattern"
	"GoLang-Design-Patterns/FactoryPattern"
	"GoLang-Design-Patterns/ObjectPoolPattern"
	"GoLang-Design-Patterns/ObserverPattern"
	"GoLang-Design-Patterns/PrototypePattern"
	"GoLang-Design-Patterns/SingletonPattern"
	"GoLang-Design-Patterns/StatePattern"
	"GoLang-Design-Patterns/StrategyPattern"
    "GoLang-Design-Patterns/BuilderPattern"
    "GoLang-Design-Patterns/ChainOfResponsibilityPattern"
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

    // 7. Factory Design Pattern Demo
    fmt.Println("\nFactory Design Pattern Demo")
    factory.DemoFactoryPattern()

    // 8. Object Pool Design Pattern Demo
    fmt.Println("\nObject Pool Design Pattern Demo")
    objectpool.DemoObjectPoolPattern()

    // 9. Builder Design Pattern Demo
    fmt.Println("\nBuilder Design Pattern Demo")
    builder.DemoBuilderPattern()

    // 10. Chain of Responsibility Design Pattern Demo
    fmt.Println("\nChain of Responsibility Design Pattern Demo")
    chain.DemoChainOfResponsibilityPattern()
}
