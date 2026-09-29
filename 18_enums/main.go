package main

import "fmt"

// enumerated types

// type MyType string -> customtype

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
)

type OrderStatusAnother string

const (
	ReceivedAnother  OrderStatusAnother = "received"
	ConfirmedAnother OrderStatusAnother = "confirm"
	PreparedAnother  OrderStatusAnother = "prepared"
)

func changeOrderStatus(status OrderStatusAnother) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changeOrderStatus(ReceivedAnother)
}
