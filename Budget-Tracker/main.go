package main

import (
	"fmt"

	"github.com/AlperSeyman/budget-tracker/model"
)

func main() {

	bt := &model.BudgetTracker{
		NextID: 1,
	}

	for {
		fmt.Println("\n--- Personal Budget Tracker---")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. Display Transaction")
		fmt.Println("3. Show Total Income")
		fmt.Println("4. Show Total Expenses")
		fmt.Println("5. Save Transaction to CSV")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:

			fmt.Print("Enter Amount: ")
			var amount float64
			fmt.Scanln(&amount)

			fmt.Print("Enter Category: ")
			var category string
			fmt.Scanln(&category)

			fmt.Print("Enter type (Income/Expense): ")
			var tType string
			fmt.Scanln(&tType)

			bt.AddTransaction(amount, category, tType)
			fmt.Println("Transaction Added!")

		case 2:

			bt.DisplayTransactions()

		case 3:

			fmt.Printf("Total Income: %.2f\n", bt.CalculateTotalAmount("income"))

		case 4:

			fmt.Printf("Total Expenses: %.2f\n", bt.CalculateTotalAmount("expense"))

		case 5:

			fmt.Println("Enter filename (e.g. transactions.csv): ")
			var filename string
			fmt.Scanln(&filename)
			if err := bt.SaveToCSV(filename); err != nil {
				fmt.Println("Error saving transactions :", err)
			}

		case 6:

			fmt.Println("Exiting....")
			return

		default:
			fmt.Println("Invalid Choice! Try Again")

		}
	}
}
