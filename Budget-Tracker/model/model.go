package model

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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
	NextID       int
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
		ID:       bt.NextID,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     tType,
	}

	bt.transactions = append(bt.transactions, newTransaction)
	bt.NextID++
}

func (bt BudgetTracker) DisplayTransactions() {

	fmt.Println("ID\tAmount\tCategory\tDate\tType")

	for _, transaction := range bt.transactions {
		fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n", transaction.ID, transaction.Amount, transaction.Category, transaction.Date.Format("2006-01-02"), transaction.Type)
	}
}

func (bt BudgetTracker) CalculateTotalAmount(tType string) float64 {

	var total float64

	for _, transaction := range bt.transactions {
		if transaction.Type == tType {
			total = transaction.Amount + total
		}
	}
	return total

}

func (bt BudgetTracker) SaveToCSV(filename string) error {

	file, err := os.Create(filename) // create file
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file) // creating a new CSV file
	defer writer.Flush()

	writer.Write([]string{"ID", "Amount", "Category", "Date", "Type"}) // write the CSV header

	// Write Data
	for _, transaction := range bt.transactions {

		record := []string{
			strconv.Itoa(transaction.ID),
			fmt.Sprintf("%.2f", transaction.Amount),
			transaction.Category,
			transaction.Date.Format("2006-01-02"),
			transaction.Type,
		}

		writer.Write(record)

	}
	fmt.Println("Transactions saved to", filename)
	return nil

}
