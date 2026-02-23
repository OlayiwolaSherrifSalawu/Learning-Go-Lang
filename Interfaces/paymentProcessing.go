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
type OrderDetail struct{
	IsCompleted bool
	Status string
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


func (cc CheckOut) CompleteCheckout() (OrderDetail, error){
	
}
