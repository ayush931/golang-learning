package main

import (
	"fmt"
	"time"
)

// struct embedding
type customer struct {
	name string
	phone string
}

// Order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) getAmount() float32 {
	return o.amount
}

func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
		createdAt: time.Now(),
	}

	return &myOrder
}

func main() {
	// If u don't set any field, default value 0 be assigned.
	// int => 0, float => 0, bool => false, string => ""
	currentOrder := order{
		id:     "123abc",
		amount: 56.34,
		status: "Running",
	}

	currentOrder.createdAt = time.Now()

	fmt.Println("Order struct", currentOrder)

	fmt.Println(currentOrder.amount)

	myOrder := order{
		id: "123",
		amount: 47,
		status: "pending",
		createdAt: time.Now(),
	}

	myOrder.status = "delivered"
	myOrder.changeStatus("confirmed")

	fmt.Println(myOrder.status)
	fmt.Println(myOrder)

	fmt.Println(myOrder.getAmount())

	anotherOrder := newOrder("1", 54.32, "received")

	fmt.Println(anotherOrder)

	language := struct{
		name string
		isGood bool
	} {"golang", true}

	fmt.Println(language)

	// struct embedding

	newCustomer := customer{
		name: "ayush",
		phone: "3243456554",
	}

	newCustomer.name = "robin"

	myOrder1 := order{
		id: "1",
		amount: 30.34,
		status: "dispatch",
		customer: newCustomer,
	}

	fmt.Println(myOrder1)
}
