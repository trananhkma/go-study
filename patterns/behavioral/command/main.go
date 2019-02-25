package main

import "fmt"

// Command
type Order struct {
	meal []string
}

type Worker interface {
	execute()
}

type Waiter struct {
	Order
}

func (w *Waiter) execute() {
	fmt.Println("This is Waiter")
	fmt.Println("Get order from Customer")
	fmt.Println("Send order to Cooker: ", w.Order.meal)
}

// Receiver
type Cooker struct {
	Order
}

func (c *Cooker) execute() {
	fmt.Println("This is Cooker")
	fmt.Println("Get order from Waiter")
	fmt.Println("Cooking meal ", c.Order.meal)
}

// Invoker
type Customer struct{}

func (c *Customer) mail_order(meal []string) {
    order := Order{meal: meal}
	waiter := Waiter{order}
	cooker := Cooker{order}
	execute_order(&waiter, &cooker)
}

func execute_order(workers ...Worker) {
	for _, w := range workers {
		w.execute()
	}
}

func main() {
	customer := new(Customer)
	meal := []string{"Salat", "Soup", "Rice"}
	customer.mail_order(meal)
}
