/*
Chain of Responsibility Pattern:
It helps to build a loosly coupled systems in softwrae design, where the called only knows the next step to be executed
and nothing else.
It comes unser behavioral design patterns. It lets you create a chain of functions or request handlers, where each handler
can either process the request or pass it to the next handler in the chain.
Every incomming request is passed through chain of functions and each function can decide:
1. Whether to skip or process the request.
2. Whether to pass the request to the next function.

3 participants of Chain of Responsibility Pattern:
1. Handler: First object that receives the request.
2. Concrete implementation of functionality(chain): these are objects that contains logic to be executed 
and decide whether to process the request or pass it to the next handler in the chain.
3. Client: Initiates the request and passes it to the first handler in the chain.

WHY:
1. When there are multiple functionalities to process the same request.
2. When we want system to be loosly coupled so that client should only know about the first function in the chain
of functions and rest function decides of what to do with the request.

Write a middleware.
*/

package chain

import (
	"fmt"
)

// Incoming HTTP request structure
type Request struct {
	UserID string
	AuthToken string
	Data string
}

// Middleware interface defines the method that each handler must implement
type Middleware interface {
	Execute(req *Request)
	SetNext(next Middleware)
}

// BaseMiddleware is a base struct that implements the Middleware interface for chaining
type BaseMiddleware struct {
	next Middleware
}

func (b *BaseMiddleware) SetNext(next Middleware) {
	b.next = next
}

func (b *BaseMiddleware) Execute(req *Request) {
	if b.next != nil {
		b.next.Execute(req)
	}
}

// LoggingMiddleware logs the request details
type LoggingMiddleware struct {
	BaseMiddleware
}

func (l *LoggingMiddleware) Execute(req *Request) {
	fmt.Println("Logging request for user:", req.UserID)
	fmt.Println("Request data:", req.Data)
	l.BaseMiddleware.Execute(req) // Pass to the next middleware
}

// AuthenticationMiddleware checks if the request has a valid authentication token
type AuthenticationMiddleware struct {
	BaseMiddleware
}

func (a *AuthenticationMiddleware) Execute(req *Request) {
	if req.AuthToken == "" {
		fmt.Println("Authentication failed: No auth token provided")
		return
	}
	fmt.Println("Authentication successful for user:", req.UserID)
	a.BaseMiddleware.Execute(req) // Pass to the next middleware
}

// ValidationMiddleware checks if the request data is valid
type ValidationMiddleware struct {
	BaseMiddleware
}

func (v *ValidationMiddleware) Execute(req *Request) {
	if req.Data == "" {
		fmt.Println("Validation failed: No data provided")
		return
	}
	fmt.Println("Validation successful for user:", req.UserID)
	v.BaseMiddleware.Execute(req) // Pass to the next middleware
}

// DemoChainOfResponsibilityPattern
func DemoChainOfResponsibilityPattern() {
	// Create a request
	req := &Request{
		UserID:    "12345",
		AuthToken: "valid_token",
		Data:      "Sample request data",
	}

	// Create middleware chain
	logging := &LoggingMiddleware{}
	authentication := &AuthenticationMiddleware{}
	validation := &ValidationMiddleware{}

	// Set up the chain
	logging.SetNext(authentication)
	authentication.SetNext(validation)

	// Execute the middleware chain
	logging.Execute(req)
	fmt.Println("Request processing completed.")
	
	fmt.Println("Chain of Responsibility Pattern Demo completed")
}