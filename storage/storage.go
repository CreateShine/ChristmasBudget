package storage

import (
	"encoding/json"
	"io/ioutil"

	"gitlab.com/parallellearning/lessons/lesson-08/Sunshine-ChristmasBudget/budgetcli/budgetapi"
)

const filename = "budgets.json"

func Load() error {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var savedBudgets []*budgetapi.Budget
	err = json.Unmarshal(fileContents, &savedBudgets)
	if err != nil {
		return err
	}

	budgetapi.SetBudgets(savedBudgets)

	return nil
}

func Save() error {
	budgetsList := budgetapi.ListBudgets()

	budgetsListBytes, err := json.MarshalIndent(budgetsList, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, budgetsListBytes, 0775)
	if err != nil {
		return err
	}

	return nil
}
