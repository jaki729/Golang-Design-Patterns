/*
SOLID - Dependency Inversion Principle (DIP)
The Dependency Inversion Principle states that high-level modules should not depend on low-level modules.
Both should depend on abstractions (e.g., interfaces).
This means that the high-level modules should not be tightly coupled with the low-level modules,
but should instead depend on abstractions that can be implemented by the low-level modules.
This principle helps in reducing the coupling between different modules and makes the code more flexible
and easier to maintain.
In Go, we can apply DIP by using interfaces to define the dependencies between high-level and low-level modules.
This allows us to change the implementation of the low-level modules without affecting the high-level modules.
*/

package solid

// ========= SurveyService End =========
// The SurveyManager is responsible for managing the survey operations.
// It depends on the Repository interface to perform operations on surveys form srp.go file.

type SurveyManager struct {
	store Repository
}

func NewSurveyManager(repo Repository) *SurveyManager {
 return &SurveyManager{
  store: repo,
 }
}