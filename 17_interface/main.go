package main

import "fmt"

type paymeter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymeter
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGateway := razorpay{}
	// razorpayPaymentGateway.pay(amount)

	// stripePaymentGateway := stripe{}
	p.gateway.pay(amount)

}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct {
}

func (s stripe) pay(amount float32) {
	fmt.Println("Make payment using stripe", amount)
}

type paypal struct {
}

func (p paypal) pay(amount float32) {
	fmt.Println("Make payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {
	
}

type fakepayment struct {
}

func (f fakepayment) pay(amount float32) {
	fmt.Println("Make payment using fake payment for testing purpose", amount)
}

func main() {
	// stripePaymentGateway := stripe{}
	// razorpayPaymentGateway := razorpay{}
	// fakePaymentGateway := fakepayment{}
	paypalPaymentGatewy := paypal{}

	newPayment := payment{
		gateway: paypalPaymentGatewy,
	}
	newPayment.makePayment(100)
}
