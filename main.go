package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/CreateShine/ChristmasBudget/budgetapi"
	"github.com/CreateShine/ChristmasBudget/storage"
	"github.com/manifoldco/promptui"
)

const (
	createBudget      = "Create a New Budget"
	viewAndEditBudget = "View and Edit Budget"
)

func main() {
	err := storage.Load()
	if err != nil {
		fmt.Println("Error Loading Arcades from file", err)
	}
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

		//If create new then have user fill in new budget details
		//Prompt User to Choose BudgetGroups
		//Promt User to Choose budget amounts for group
		//Prompt User to Input Names for Each Budget Group
		/*Display print of budget and say "Great - you are done!"
		and option to edit or return to main screen*/

		//If copy then prompt user "Type the New Name"
		//Allow user to edit

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

	newBudget := &budgetapi.Budget{
		Name:       budgetName,
		TotalPrice: budgetTotal,
		Groups:     nil,
	}
	budgetapi.CreateBudget(newBudget)
	editGroupPrompt(newBudget)

	err = storage.Save()
	if err != nil {
		return err
	}

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
	availableBudgetsToEdit := budgetapi.ListBudgets()

	if len(availableBudgetsToEdit) == 0 {
		return errors.New("Need to Create Budget")
	}

	budgetapi.ListBudgetsToEdit()
	time.Sleep(5000 * time.Millisecond)
	return nil
}

//Returns a list of budgets available
func viewBudgetPrompt() error {
	availableBudgets := budgetapi.ListBudgets()

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

	//Should I add a function here that looks at which groups are greater than 0 and returns the names of these?
	//fmt.Println("Name:", chosenBudget.Name)
	fmt.Println("Name:", chosenBudget.Name, ", Total Price: $", chosenBudget.TotalPrice, ", Groups:")

	editGroupPrompt(chosenBudget)
	err = storage.Save()
	if err != nil {
		return err
	}
	//time.Sleep(5000 * time.Millisecond)
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

		// var budgetGroupAmount []string
		// for _, budget := range availableGroups {
		// 	if availableGroups.GroupPrice > 0 {
		// 		options = append(budgetGroupAmount, arcade.Name)
		// 	}
		// }

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
