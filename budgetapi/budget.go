package budgetapi

import (
	"database/sql"
	"fmt"
)

//Connection to Database
type BudgetsService struct {
	db *sql.DB
}

func NewService(db *sql.DB) *BudgetsService {
	return &BudgetsService{
		db: db,
	}
}

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
