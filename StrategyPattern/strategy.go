/*
Strategy Pattern: (also known as Policy Pattern)
Its the behavioral design pattern that is used when there are multiple algorithms for a specific task.
And clinet decides the actual implementation to be used at runtime.
It defines a family of algorithms, encapsulates each one, and makes them interchangeable.
The strategy pattern lets the algorithm vary independently from clients that use it.

Pros:
1. The application can switch between strategies at runtime.
2. Strategy enables the client to choose the algorithm that best suits its needs without using conditional statements.
The Data structure implementation is completly encapsulated in the strategy. 
So, implementation can be changed without affecting the client/context code/classes.
By encapsulating the algorithm in separate classes, new algorithms can be implemented with the same interface.

Cons:
1. The application must be aware of all the strategies available, to choose the right one for right situation/solutions.
*/

package strategy

import (
	"fmt"
)

// struct can be assigned to an interface any type that implements the interface methods in runtime
// interface in go has properties that it can accept any type or value that implements the methods defined in the interface
// it can accept one value at a time, and it can be used to define a contract that any type must implement
type IDBConnection interface { 
	Connect()
}

// Compatible struct that implements the IDBConnection interface to accept any type of database connection in runtime
type DBCOnnection struct {
	Db IDBConnection
}

// Receiver function for struct DBCOnnection
func (con DBCOnnection) DbConnect() {
	if con.Db == nil {
		fmt.Println("No database connection provided")
		return
	}
	con.Db.Connect()
}

// Implementation of IDBConnection interface for MySQL (First behavior)
type MySQLConnection struct{
	ConnectionString string
}

// Receiver function for MySQLConnection struct
func (con MySQLConnection) Connect() {
	fmt.Println("Connecting to MySQL database with connection string:", con.ConnectionString)
}

// Second behavior implementation of IDBConnection interface for PostgreSQL
type PostgreSQLConnection struct {
	ConnectionString string
}

// Receiver function for PostgreSQLConnection struct
func (con PostgreSQLConnection) Connect() {
	fmt.Println("Connecting to PostgreSQL database with connection string:", con.ConnectionString)
}

// Third behavior implementation of IDBConnection interface for MongoDB
type MongoDBConnection struct {
	ConnectionString string
}

// Receiver function for MongoDBConnection struct
func (con MongoDBConnection) Connect() {
	fmt.Println("Connecting to MongoDB database with connection string:", con.ConnectionString)
}

func DemoStrategyPattern() {

	var dbConnection DBCOnnection // Declare a variable of type DBCOnnection that can hold any database connection

	// Create instances of different database connections
	mysql := MySQLConnection{ConnectionString: "This a MySQL connection"}
	postgres := PostgreSQLConnection{ConnectionString: "This is a PostgreSQL connection"}
	mongo := MongoDBConnection{ConnectionString: "This is a MongoDB connection"}
	
	// Connect to DBCOnnection
	dbConnection = DBCOnnection{Db: mysql}
	dbConnection.DbConnect()
	
	// Connect to PostgreSQL 
	dbConnection = DBCOnnection{Db: postgres}
	dbConnection.DbConnect()

	// Connect to MongoDB
	dbConnection = DBCOnnection{Db: mongo}
	dbConnection.DbConnect()
}