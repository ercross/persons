package main

import (
	"fmt"
	"sort"
)

type Person struct {
	ID         string `json:"id"`
	PersonName string `json:"personName"`
	Salary     Salary `json:"salary"`
}

type Salary struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

type Persons struct {
	api  ExchangeRateAPI
	Data []Person `json:"data"`
}

type ExchangeRateAPI interface {

	// GetExchangeRate gets exchange rate from API.
	// pair should be formatted as BASE_CURRENCY-QUOTE_CURRENCY e.g., EUR-USD
	GetExchangeRate(pair string) (float64, error)
}

type MockExchangeRateAPI struct {
	currencyPairToExchangeRate map[string]float64
}

func NewMockExchangeRateAPI() *MockExchangeRateAPI {
	return &MockExchangeRateAPI{
		currencyPairToExchangeRate: map[string]float64{
			"USD-USD": 1.0,
			"EUR-USD": 1.1,
			"GBP-USD": 1.25,
			"JPY-USD": 0.0087,
			"AUD-USD": 0.65,
			"CAD-USD": 0.75,
			"INR-USD": 0.012,
			"NZD-USD": 0.61,
			"CHF-USD": 1.08,
			"SGD-USD": 0.73,
			"BRL-USD": 0.19,
			"ZAR-USD": 0.052,
		},
	}
}

func (api *MockExchangeRateAPI) GetExchangeRate(pair string) (float64, error) {
	rate, exists := api.currencyPairToExchangeRate[pair]
	if !exists {
		return 0, fmt.Errorf("exchange rate not found for pair: %s", pair)
	}
	return rate, nil
}

// ConvertSalariesToDollar converts all salaries in the Persons object to USD equivalent
func (p *Persons) ConvertSalariesToDollar() (*Persons, error) {
	for i := range p.Data {
		if p.Data[i].Salary.Currency != "USD" {
			pair := p.Data[i].Salary.Currency + "-USD"
			rate, err := p.api.GetExchangeRate(pair)
			if err != nil {
				return nil, err
			}
			p.Data[i].Salary.Value *= rate
			p.Data[i].Salary.Currency = "USD"
		}
	}
	return p, nil
}

// SortBySalaryAsc sorts the persons by salary in ascending order (assumes all salaries are in USD)
func (p *Persons) SortBySalaryAsc() {
	sort.Slice(p.Data, func(i, j int) bool {
		return p.Data[i].Salary.Value < p.Data[j].Salary.Value
	})
}

// SortBySalaryDesc sorts the persons by salary in descending order (assumes all salaries are in USD)
func (p *Persons) SortBySalaryDesc() {
	sort.Slice(p.Data, func(i, j int) bool {
		return p.Data[i].Salary.Value > p.Data[j].Salary.Value
	})
}

// FilterBySalaryCriteria filters persons whose salaries meet the specified criteria (assumes all salaries are in USD)
func (p *Persons) FilterBySalaryCriteria(minSalary, maxSalary float64) []Person {
	var filteredPersons []Person
	for _, person := range p.Data {
		if person.Salary.Value >= minSalary && person.Salary.Value <= maxSalary {
			filteredPersons = append(filteredPersons, person)
		}
	}
	return filteredPersons
}
