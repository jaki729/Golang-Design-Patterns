/*
SOLID - Interface Segregation Principle (ISP)
The Interface Segregation Principle states that no client should be forced to depend on methods it doesnot use.
This means that interfaces should be designed in such a way that they are specific to the client
and do not contain unnecessary methods that the client does not need.
This principle helps in reducing the complexity of the code and makes it easier to maintain and understand.
In Go, we can apply ISP by creating small, specific interfaces that are tailored to the needs
of the client rather than large, general-purpose interfaces.
*/

package solid

type Question interface {
	SetTitle(title string)
}

type QuestionWithOptions interface {
	Question
	AddOption(option string)
}

type TextInputQuestion struct {
 title string
}

func (q *TextInputQuestion) SetTitle(title string) {
 q.title = title
}

type DropDownQuestion struct {
 title   string
 options []string
}

func (q *DropDownQuestion) SetTitle(title string) {
 q.title = title
}	

func (q *DropDownQuestion) AddOption(option string) {
	 q.options = append(q.options, option)
}