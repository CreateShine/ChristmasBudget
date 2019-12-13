package budgetapi

import (
	"database/sql"
	"fmt"
	_ "strings"
)

const (
	insertBudgetQuery = "INSERT INTO budgets (budget_name, budget_price) VALUES (?,?);"

	insertGroupQuery = "INSERT INTO groups (group_name, group_price, budget_id, people) VALUES (?,?,?,?);"

	selectBudgetsQuery = "SELECT id, budget_name, budget_price FROM budgets"

	selectGroupsQuery = "SELECT group_name WHERE budget_id = ?"

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
	ID         int
	Name       string
	TotalPrice float64
	Groups     []*Group
}

type Group struct {
	ID         int
	Name       string
	GroupPrice float64
	People     []string
	BudgetID   int
}

func SetBudgets(a []*Budget) {
	budgets = a
}

//Creation of a new Christmas Budget
//This function allows the user to create a new budget and then adds it to the list of budgets
func (b *BudgetsService) CreateBudget(budgetName string, budgetPrice float64) (*Budget, error) {
	result, err := b.db.Exec(insertBudgetQuery, budgetName, budgetPrice)

	newBudgetID64, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	newBudgetID := int(newBudgetID64)

	group1, err := b.createGroup("Inlaws-His", newBudgetID)
	if err != nil {
		return nil, err
	}
	group2, err := b.createGroup("Inlaws-Hers", newBudgetID)
	if err != nil {
		return nil, err
	}
	group3, err := b.createGroup("PersonalFamily", newBudgetID)

	if err != nil {
		return nil, err
	}

	budget := &Budget{
		ID:         newBudgetID,
		Name:       budgetName,
		TotalPrice: budgetPrice,
		Groups: []*Group{
			group1, group2, group3,
		},
	}

	return budget, nil
}

func (b *BudgetsService) createGroup(groupName string, newBudgetID int) (*Group, error) {
	result, err := b.db.Exec(insertGroupQuery, groupName, 0, newBudgetID, "")

	if err != nil {
		return nil, err
	}

	newGroupID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	group := &Group{
		Name:     groupName,
		BudgetID: newBudgetID,
		ID:       int(newGroupID),
	}
	return group, nil
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

//function to pack groups together and then to unpack them.
/*func GroupCompiler() {
	groups := []string{"Inlaws-His", "Inlaws-Hers", "Personal Family"}

	groupsCSVStr := strings.Join(groups, ",")

	fmt.Println(groupsCSVStr)

	hairdressersFromCSV := strings.Split(groupsCSVStr, ",")

	fmt.Println(groupsCSVStr)

}
*/
