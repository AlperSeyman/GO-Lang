package model

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string
}

type BudgetTracker struct {
	transactions []Transaction
	nextID       int
}

type FinancialRecord interface {
	GetAmount() float64
	GetType() string
}

func (t Transaction) GetAmount() float64 {
	return t.Amount
}

func (t Transaction) GetType() string {
	return t.Type
}

func (bt *BudgetTracker) AddTransaction(amount float64, category string, tType string) {

	newTransaction := Transaction{
		ID:       bt.nextID,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     tType,
	}

	bt.transactions = append(bt.transactions, newTransaction)
	bt.nextID++
}

func (bt BudgetTracker) DisplayTransactions() {

	fmt.Println("ID\tAmount\tCategory\tDate\tType")

	for _, transaction := range bt.transactions {
		fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n", transaction.ID, transaction.Amount, transaction.Category, transaction.Date, transaction.Type)
	}
}
