package budgetapi

import (
	"fmt"
)

//Naming of variables
var (
	budgets []*Budget
)

type Budget struct {
	Name       string
	TotalPrice float64
	Groups     []*Group
}

type Group struct {
	Name       string
	People     []string
	GroupPrice float64
}

func SetBudgets(a []*Budget) {
	budgets = a
}

//Creation of a new Christmas Budget

//This function allows the user to create a new budget and then adds it to the list of budgets
func CreateBudget(budget *Budget) {

	budget.Groups = []*Group{
		&Group{
			Name:       "Inlaws - His",
			People:     nil,
			GroupPrice: 0,
		},
		&Group{
			Name:       "Inlaws - Hers",
			People:     nil,
			GroupPrice: 0,
		},
		&Group{
			Name:       "Personal Family",
			People:     nil,
			GroupPrice: 0,
		},
		&Group{
			Name:       "Done",
			People:     nil,
			GroupPrice: 0,
		},
	}
	budgets = append(budgets, budget)

}

//This function displays a list of budgets (so user can then copy or edit)
func ListBudgets() []*Budget {
	return budgets
}

func (G *Group) PrintName() {
	fmt.Println(G.Name)
}

//This group allows you to edit a budget.
//First, list the budgets you could choose from to edit. And allow to select
func ListBudgetsToEdit() {
	//fmt.Println(budgets)
	//return budgets

	/*editPrompt := promptui.Select{
		Label: "Choose Which Budget to Edit",
		Items: budgets,
	}
	_, result, err := editPrompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	return
		//fmt.Printf("You choose %q\n", result)*/
}

/*func ListGroups() []Group {
	return arcades
}*/

//This Function lists the Groups that were chosen for user's budget
/*func ListChosenGroups() []*BudgetGroup {
	//Here add functionality that makes it list the groups
	return &BudgetGroup
}*/

//Budget needs method called AddBudgetGroup
//Should this string below be the actual name of Groups
/*func (b *Budget) AddBudgetGroup() string {
	groups := []BudgetGroup{"Group1", "Group2", "Group3", "Group4"}
	addGroupPrompt := promptui.Select{
		Label: "Which Group Would You Like to Add?",
		Items: groups,
	}
	_, result, err := addGroupPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

	}
	return result
}*/

//BudgetGroup needs method called Add Person
//func (p *BudgetGroup)

//Budget needs method Verify that price is right - make sure each summed is less  than Total budget - return bool
/*func (*Budget budgets) CalculateTotalPrice(numberOfGames float64) float64 {
	return numberOfGames*arcade.PricePerGame + arcade.EntryPrice
}*/
