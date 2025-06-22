/*
Solid Principles - Liskov Substitution Principle (LSP)
The Liskov Substitution Principle states that objects of a superclass should be replaceable with
objects of a subclass without affecting the correctness of the program.
This means that if a class is a subclass of another class, it should be able to
be used in place of the superclass without causing any issues.
Nutshell, objects can be swapped without altering the correctness of the program.
Go, do not have inheritance like other OOP languages, but it does have interfaces
and struct embedding which can be used to achieve similar behavior.
In Go, we can apply LSP by ensuring that structs and interfaces are designed in such a way that
they can be substituted for each other without breaking the functionality of the code.
*/

package solid

import (
	"encoding/json"
	"io"
	"log"
	 "os"
)

func SendSurvey(s *Survey, w io.Writer) error {
	b, err := json.Marshal(s)
	if err != nil {
   		return err
	}
	_,err = w.Write(b)
	if err != nil {
		return err
	}	
	return nil
}

func example() {
	file,err := os.Open("survey.json")
	if err != nil {
		log.Fatal(err)
	}

	s := &Survey{
		  Title:      "Customer Satisfaction Survey",
	}
	SendSurvey(s, file)
}

// type SurveyMain struct {
// 	Title string `json:"title"`
// }

// func main() {
// 	example()
// }