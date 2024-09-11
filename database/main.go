package main

import (
	"fmt"
	"log"
)

type User struct {
	ID   int
	Name string
}

// Concrete Implementations - is a specific class or struct that implements an interface or abstract class. It provides the actual behavior and functionality for the methods defined in the interface or abstract class.
// type MockDatastore struct { /*...*/ }
// type SQLDatastore struct { /*...*/ }
// type FileDatastore struct { /*...*/ }

type MockDatastore struct {
	Users map[int]User
}

// Crate a Value Receiver "md MockDatastore" for GetUser method
// If the method is defined with a value receiver, can call it on both Pointer and Value Instance
func (md MockDatastore) GetUser(id int) (User, error) {
	user, exist := md.Users[id]
	if !exist {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil
}

// Create a Pointer Receiver "md *MockDatastore" for SaveUser method
// If the method is defined with a pointer receiver, need a Pointer Instance "&MockDatastore{}"
func (md *MockDatastore) SaveUser(u User) error {
	_, exist := md.Users[u.ID]
	if exist {
		return fmt.Errorf("User %v already exist", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

// Flexible to use different data storage mechanisms
// Interface defines what mdthods are availabe, Concrete Implementations define how those methods actually perform in their tasks.
type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

// Service Layer
// Decoupling Data Access and Business logic
// it doesn't need to know how data is stored or retrieve from where. So, it use interface

type Service struct {
	ds Datastore
}

func (s *Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s *Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	u1 := User{2, "Micheal"}
	u2 := User{1, "James Doe"}

	// "&MockDatastore" creat a new MockDatastore instance and returns a pointer to it. This pointer is store in the variablae "db"
	// Pointer Instance
	// called Poitner Dereferncing
	db := &MockDatastore{
		Users: make(map[int]User),
	}

	// Value Instance
	// db := &MockDatastore{
	// 	Users: make(map[int]User),
	// }

	serv := &Service{ds: db}
	err := serv.SaveUser(u1)
	err1 := serv.SaveUser(u2)
	if err != nil || err1 != nil {
		log.Fatalf("Error %s %s", err, err1)
	}

	u1Returned, err := serv.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	// log.Println("Input ", u1)
	log.Println("Output ", u1Returned)
	log.Println("All Users ", db.Users)

}
