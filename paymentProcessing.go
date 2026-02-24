package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type PaymentMethod interface {
	ProcessPayment(amount float64) error
	RefundPayment(transactionId string, amount float64) error
	Authorize(account RequestModel) ResponseModel
	Capture(transactionId string) ResponseModel
}
type OrderDetail struct {
	IsCompleted  bool
	Status       string
	ProccessedAt time.Time
}
type RequestModel struct {
	IncomingAmount float64
	RequestId      string
	SecureToken    string
}

type ResponseModel struct {
	Success       bool
	TransactionId string
	ErrorMessage  error
	Processedat   time.Time
}
type CheckOut struct {
	GateWay PaymentMethod
}

// Dependency Injection.
type PaymentProcessor struct{}

func (pp PaymentProcessor) Authorize(account RequestModel) ResponseModel {
	response := ResponseModel{}
	if account.IncomingAmount <= 0 {
		response.ErrorMessage = errors.New("Can not withdraw amount lower than or equal to zero.\n")
		return response
	}
	if account.SecureToken == "" {
		response.ErrorMessage = errors.New("token cant be empty.\n")
		return response
	}
	if account.SecureToken == "succes_token" {
		response.Success = true
		response.Processedat = time.Now()
		response.TransactionId = "MOCK-TXN-98765"
		return response
	} else {
		response.Success = false
		response.Processedat = time.Now()
		return response
	}
}

func (pp PaymentProcessor) Capture(transactionId string) ResponseModel {
	response := ResponseModel{}
	status := 1 + rand.Intn(10-1+1)
	if transactionId == "" && status >= 2 {
		response.ErrorMessage = errors.New("cannot access transactionid {empty field} \n")
		return response
	}
	if transactionId == "MOCK-TXN-98765" {
		fmt.Printf("Mock Bank: Successfully captured funds for transaction...\n")
		time.Sleep(200 * time.Millisecond)
		response.Success = true
		response.Processedat = time.Now()
		response.TransactionId = "MOCK-TXN-98765"
		return response
	} else {
		response.Success = false
		response.Processedat = time.Now()
		response.ErrorMessage = errors.New("Bank TimeOut\n")
		return response
	}
}

func (pp PaymentProcessor) RefundPayment(transactionId string, amount float64) error {
	if amount <= 0 {
		return errors.New("cannot refund negative and zero amount ")
	}
	fmt.Printf("Refunded %.2f to %s \n", amount, transactionId)
	return nil
}

func (pp PaymentProcessor) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return errors.New("cannot refund negative and zero amount ")
	}
	fmt.Printf("processed amount: %.2f \n", amount)
	return nil
}

func (cc CheckOut) CompleteCheckout(orders Order) (OrderDetail, error) {

	// check  if incoming order is empty or has no item
	if orders.Items == nil || orders.TotalAmount == 0.0 {
		return OrderDetail{Status: "FAILED"}, errors.New("Cannot CheckOut an empty catalogue")
	}
	// we dont want to checkout one order more than once
	// i could do this by iterating through the store and if the orderId as occured before we stop the function.
	// Or i could check the incoming order status if its empty then it is a new order and then i change it to pending else i stop
	// doing the first.
	// (I could'nt access the OrderStore from this method so i would try the second approach)
	if orders.PaymentStatus == "" {
		orders.PaymentStatus = "Pending"
	} else {
		return OrderDetail{Status: "FAILED"}, errors.New("Order is already existing and order status: " + orders.PaymentStatus)
	}
	// Calculate the total amount
	amount := orders.CalculateTotal()
	secureToken := "succes_token"
	Request := RequestModel{
		IncomingAmount: amount,
		SecureToken:    secureToken,
		RequestId:      orders.OrderId,
	}
	// parse all this to the gateway
	InitResponse := cc.GateWay.Authorize(Request)
	Capture := cc.GateWay.Capture(InitResponse.TransactionId)
	if Capture.Success == true {
		Processed := cc.GateWay.ProcessPayment(amount)
		if Processed == nil {
			orders.PaymentRefrence = Capture.TransactionId
			orders.PaymentStatus = "PAID"
			orders.OrderDate = Capture.Processedat
			orders.TotalAmount = amount
			OrderManager.OrderStore = append(orders)
			return OrderDetail{IsCompleted: true, Status: "Paid", ProccessedAt: Capture.Processedat}, nil
		} else {
			return OrderDetail{IsCompleted: false, Status: "FAILED", ProccessedAt: Capture.Processedat}, errors.New("Could not complete the payment processes")
		}
	} else {
		return OrderDetail{IsCompleted: false, Status: "FAILED", ProccessedAt: Capture.Processedat}, errors.New("Could not complete the payment processes")
	}

}
