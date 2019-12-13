package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	//"strings"
	"os"

	"github.com/CreateShine/ChristmasBudget/budgetapi"
	"github.com/CreateShine/ChristmasBudget/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
)

const (
	createBudget      = "Create a New Budget"
	viewAndEditBudget = "View and Edit Budget"
)

var budgetsService *budgetapi.BudgetsService

func main() {
	db, err := db.ConnectDatabase("budgets_db.config")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	budgetsService = budgetapi.NewService(db)
	//	fmt.Println(budgetsService)

	//Prompt user to choose whether to create a new budget or copy old
	for {

		fmt.Println("Welcome to ChristmasBudget!")

		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{
				createBudget,
				viewAndEditBudget,
			},
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		switch result {
		case createBudget:
			err := addBudgetPrompt()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
			}

		case viewAndEditBudget:
			err := viewBudgetPrompt()
			if err != nil {
				fmt.Println("No Budgets to View.", err)
			}

		}

		time.Sleep(3000 * time.Millisecond)

	}
}

func addBudgetPrompt() error {
	namePrompt := promptui.Prompt{
		Label: "Name of Budget",
	}

	budgetName, err := namePrompt.Run()
	if err != nil {
		return err
	}

	budgetTotal, err := promptForNumber("Budget Total")
	if err != nil {
		return err
	}

	fmt.Println("here5.", budgetsService)
	newBudget, err := budgetsService.CreateBudget(budgetName, budgetTotal)
	if err != nil {
		return err
	}
	fmt.Println("here6")
	editGroupPrompt(newBudget)
	fmt.Println("here7")

	//budgetsService.addBudgetPrompt(budgetName, budgetTotal)
	fmt.Println("Well done you created: ", budgetName)
	return nil
}

func promptForNumber(label string) (float64, error) {
	numberPrompt := promptui.Prompt{
		Label: label,
	}
	numberStr, err := numberPrompt.Run()
	if err != nil {
		return 0, err
	}
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func editBudgetPrompt() error {
	availableBudgetsToEdit, err := budgetsService.ListBudgets()
	if err != nil {
		return err
	}
	if len(availableBudgetsToEdit) == 0 {
		return errors.New("Need to Create Budget")
	}

	budgetsService.ListBudgets()
	time.Sleep(5000 * time.Millisecond)
	return nil
}

//Returns a list of budgets available
func viewBudgetPrompt() error {
	availableBudgets, err := budgetsService.ListBudgets()
	if err != nil {
		return err
	}
	fmt.Println("here3")
	if len(availableBudgets) == 0 {
		return errors.New("No budgets created")
	}

	/*for _, name := range availableBudgets {
		fmt.Println(budgetapi.Budget.Name)
	}

	fmt.Println("Here2")
	fmt.Println(availableBudgets)
	//Loop through available budgets and print them
	for Name := range availableBudgets {
		fmt.Println(Name)
	}

	fmt.Println("here5")*/

	var options []string
	for _, budget := range availableBudgets {
		options = append(options, budget.Name)
	}

	selectBudgetPrompt := promptui.Select{
		Label: "Select Budget",
		Items: options,
	}

	chosenIndex, _, err := selectBudgetPrompt.Run()
	if err != nil {
		return err
	}
	chosenBudget := availableBudgets[chosenIndex]

	fmt.Println("Name:", chosenBudget.Name, ", Total Price: $", chosenBudget.TotalPrice, ", Groups:")

	editGroupPrompt(chosenBudget)

	return nil
}

//The Below Code enables you to Edit a group from the default list
func editGroupPrompt(newBudget *budgetapi.Budget) error {
	for {

		availableGroups := newBudget.Groups

		var options []string
		for _, group := range availableGroups {
			groupOption := (group.Name + " $" + strconv.FormatFloat(group.GroupPrice, 'g', 10, 64))
			options = append(options, groupOption)
		}

		selectGroupPrompt := promptui.Select{
			Label: "Select Group",
			Items: options,
		}

		chosenIndex, _, err := selectGroupPrompt.Run()
		if err != nil {
			return err
		}
		//Makes sure Done has not been chosen. Checks to see if the last option was not chosen.

		if chosenIndex == len(options)-1 {
			return nil
		}
		chosenGroup := availableGroups[chosenIndex]

		groupTotal, err := promptForNumber("Group Budget Total")
		if err != nil {
			return err
		}

		chosenGroup.GroupPrice = groupTotal

		total := sumGroupTotals(newBudget)

		//need to call budgetService

		if total > newBudget.TotalPrice {
			fmt.Println("Group Totals Exceed Budget Total of", newBudget.TotalPrice)
		}

		time.Sleep(500 * time.Millisecond)
	}
}

//Calculate a Budget Total Price and Compare it to the set price

func sumGroupTotals(b *budgetapi.Budget) float64 {

	totalx := 0.0
	for _, group := range b.Groups {
		totalx += group.GroupPrice

	}
	return totalx

}

/*chooseGroupPrompt := promptui.Select{
	Label: "Which Groups Would You Like to Add?",
	Items: budgetapi.Groups,
}
_, result, err := chooseGroupPrompt.Run()
if err != nil {
	fmt.Printf("Prompt failed %v\n", err)

}*/
