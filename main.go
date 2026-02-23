package main

import (
	"fmt"
	"time"
)

// Order management
type OrderManager struct {
	OrderStore    []Order
	CustomerStore []CustomerInfo
}

type Item struct {
	ProductId string
	Name      string
	Price     float64
	Quantity  int
}

type CustomerInfo struct {
	CustomerId string
	FullName   string
	Email      string
}
type Order struct {
	OrderId     string
	CustomerId  string
	OrderDate   time.Time
	Items       []Item
	TotalAmount float64
}

func (total *OrderManager) Details(custId string) {

	CustomerLifetimeValue := 0.0
	found := false
	for index := range total.OrderStore {
		if total.OrderStore[index].CustomerId == custId {
			found = true
			CustomerLifetimeValue += total.OrderStore[index].CalculateTotal()
		}
	}
	if !found {
		fmt.Printf("customer id : %s does not exist\n", custId)
		return
	}
	fmt.Printf("customer id %s TotalAmount is %.2f \n", custId, CustomerLifetimeValue)
}
func (orders *Order) CalculateTotal() float64 {
	if orders == nil {
		fmt.Printf("they are no order ")
		return 0.0
	}
	orders.TotalAmount = 0
	for _, value := range orders.Items {
		orders.TotalAmount += (value.Price * float64(value.Quantity))
	}
	return orders.TotalAmount
}
func main() {
	// testing order management bluePrint

	var Customer1 = CustomerInfo{
		CustomerId: "01",
		FullName:   "Olayiwola Salawu",
		Email:      "salawuolayiwola181@gmail.com",
	}
	var Customer2 = CustomerInfo{
		CustomerId: "02",
		FullName:   "Divine Ositadima",
		Email:      "divineositadima@gmail.com",
	}
	var Customer3 = CustomerInfo{
		CustomerId: "03",
		FullName:   "Abisola Adenike",
		Email:      "abisolaadenike33@gmail.com",
	}
	var Order1 = Order{
		OrderId:    "01",
		CustomerId: "01",
		Items:      []Item{Item{ProductId: "111", Name: "Spagetti", Price: 500.0, Quantity: 10}, Item{ProductId: "333", Name: "Sardine", Price: 200.0, Quantity: 5}, Item{ProductId: "222", Name: "Chocolate Bread", Price: 300.0, Quantity: 3}},
	}
	var Order2 = Order{
		OrderId:    "03",
		CustomerId: "01",
		Items:      []Item{Item{ProductId: "404", Name: "Afang", Price: 5000.0, Quantity: 10}, Item{ProductId: "333", Name: "Pounded Yam", Price: 2000.0, Quantity: 3}, Item{ProductId: "232", Name: "Turkey", Price: 2009.0, Quantity: 2}},
	}
	var Order3 = Order{
		OrderId:    "04",
		CustomerId: "02",
		Items:      []Item{Item{ProductId: "111", Name: "Spagetti", Price: 500.0, Quantity: 10}, Item{ProductId: "333", Name: "Sardine", Price: 200.0, Quantity: 5}, Item{ProductId: "222", Name: "Chocolate Bread", Price: 300.0, Quantity: 3}},
	}
	var Order4 = Order{
		OrderId:    "06",
		CustomerId: "03",
		Items:      []Item{Item{ProductId: "121", Name: "Garri", Price: 100.0, Quantity: 100}, Item{ProductId: "434", Name: "GroundNut", Price: 50.0, Quantity: 6}, Item{ProductId: "137", Name: "Milk", Price: 100.0, Quantity: 4}},
	}
	var Sales = OrderManager{
		OrderStore:    []Order{Order1, Order2, Order3, Order4},
		CustomerStore: []CustomerInfo{Customer1, Customer2, Customer3},
	}
	Sales.Details("09")
	fmt.Println(Order1.CalculateTotal())
}
