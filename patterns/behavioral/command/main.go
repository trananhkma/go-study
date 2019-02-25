package main

import "fmt"

// Command
type Order struct {
	cook Cook
	meal []string
}

func (o *Order) execute() {
	fmt.Println("Ordering ", o.meal)
	o.cook.prepare(o.meal)
}

// Receiver
type Cook struct{}

func (c *Cook) prepare(meal []string) {
	fmt.Println("Preparing meal: ", meal)
}

// Invoker
type Customer struct {
	order Order
}

func (c *Customer) mail_order(o *Order) {
	o.execute()
}

func main() {
	c1 := new(Customer)
	meals := []string{"Fish"}
    meals = append(meals, "Cat")
    meals = append(meals, "Dog")
	
	order1 := Order{meal: meals}
    
    c2 := new(Customer)
    order2 := Order{meal: []string{"Chicken"}}
    
    
    c2.mail_order(&order2)
	c1.mail_order(&order1)
}
