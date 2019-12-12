package budgetapi

import (
	"database/sql"
	"fmt"
)

const (
	insertBudgetQuery = "INSERT INTO budgets (id, budget_name, budget_price) VALUES (?, ?,?);"

	selectBudgetsQuery = "SELECT id, budget_name, budget_price FROM budgets"

	groupNames = "Inlaws-His, Inlaws-Hers, Personal Family"
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
	ID         float64
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
func (b *BudgetsService) CreateBudget(budgetName string, budgetPrice float64, groups string) error {
	_, err := b.db.Exec(insertBudgetQuery, budgetName, budgetPrice, groupNames)

	//Do a query to find out what the budget ID is and then insert this ID into each of the groups
	/*Groups = []*Group{
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
	}*/
	//budgets = append(budgets, budget)
	if err != nil {
		return err
	}

	return nil
}

//This function displays a list of budgets (so user can then copy or edit)
func (b *BudgetsService) ListBudgets() ([]*Budget, error) {
	rows, err := b.db.Query(selectBudgetsQuery)
	if err != nil {
		return nil, err
	}
	var budgets []*Budget
	for rows.Next() {
		var budget *Budget

		err := rows.Scan(
			&budget.ID,
			&budget.Name,
			&budget.TotalPrice,
			&budget.Groups,
		)
		if err != nil {
			return nil, err
		}

		budgets = append(budgets, budget)
	}

	return budgets, nil
}

func (G *Group) PrintName() {
	fmt.Println(G.Name)
}
