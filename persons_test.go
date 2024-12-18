package main

import (
	"reflect"
	"testing"
)

func TestConvertSalariesToDollar(t *testing.T) {
	// Mock data with multiple currencies
	mockData := Persons{
		Data: []Person{
			{ID: "1", PersonName: "Cadanaut 1", Salary: Salary{Value: 10.0, Currency: "USD"}},
			{ID: "2", PersonName: "Cadanaut 2", Salary: Salary{Value: 15.0, Currency: "EUR"}},
			{ID: "3", PersonName: "Cadanaut 3", Salary: Salary{Value: 2500.0, Currency: "JPY"}},
		},
		api: NewMockExchangeRateAPI(),
	}

	// Convert salaries to USD
	converted, err := mockData.ConvertSalariesToDollar()
	if err != nil {
		t.Fatalf("Error converting salaries to USD: %v", err)
	}

	// Expected values
	expectedSalaries := []float64{10.0, 16.5, 21.75}

	for i, person := range converted.Data {
		if person.Salary.Currency != "USD" {
			t.Errorf("Expected currency USD, got %s", person.Salary.Currency)
		}
		if person.Salary.Value != expectedSalaries[i] {
			t.Errorf("Expected salary %.2f, got %.2f", expectedSalaries[i], person.Salary.Value)
		}
	}
}

func TestSortBySalaryAsc(t *testing.T) {
	mockData := Persons{
		Data: []Person{
			{ID: "1", PersonName: "Cadanaut 1", Salary: Salary{Value: 50.0, Currency: "USD"}},
			{ID: "2", PersonName: "Cadanaut 2", Salary: Salary{Value: 30.0, Currency: "USD"}},
			{ID: "3", PersonName: "Cadanaut 3", Salary: Salary{Value: 70.0, Currency: "USD"}},
		},
	}

	mockData.SortBySalaryAsc()

	expectedOrder := []float64{30.0, 50.0, 70.0}
	for i, person := range mockData.Data {
		if person.Salary.Value != expectedOrder[i] {
			t.Errorf("Expected salary %.2f, got %.2f", expectedOrder[i], person.Salary.Value)
		}
	}
}

func TestSortBySalaryDesc(t *testing.T) {
	mockData := Persons{
		Data: []Person{
			{ID: "1", PersonName: "Cadanaut 1", Salary: Salary{Value: 50.0, Currency: "USD"}},
			{ID: "2", PersonName: "Cadanaut 2", Salary: Salary{Value: 30.0, Currency: "USD"}},
			{ID: "3", PersonName: "Cadanaut 3", Salary: Salary{Value: 70.0, Currency: "USD"}},
		},
	}

	mockData.SortBySalaryDesc()

	expectedOrder := []float64{70.0, 50.0, 30.0}
	for i, person := range mockData.Data {
		if person.Salary.Value != expectedOrder[i] {
			t.Errorf("Expected salary %.2f, got %.2f", expectedOrder[i], person.Salary.Value)
		}
	}
}

func TestFilterBySalaryCriteria(t *testing.T) {
	mockData := Persons{
		Data: []Person{
			{ID: "1", PersonName: "Cadanaut 1", Salary: Salary{Value: 50.0, Currency: "USD"}},
			{ID: "2", PersonName: "Cadanaut 2", Salary: Salary{Value: 30.0, Currency: "USD"}},
			{ID: "3", PersonName: "Cadanaut 3", Salary: Salary{Value: 70.0, Currency: "USD"}},
		},
	}

	filtered := mockData.FilterBySalaryCriteria(40.0, 60.0)

	expectedFiltered := []Person{
		{ID: "1", PersonName: "Cadanaut 1", Salary: Salary{Value: 50.0, Currency: "USD"}},
	}

	if len(filtered) != len(expectedFiltered) {
		t.Fatalf("Expected %d persons, got %d", len(expectedFiltered), len(filtered))
	}

	for i, person := range filtered {
		if !reflect.DeepEqual(person, expectedFiltered[i]) {
			t.Errorf("Expected %+v, got %+v", expectedFiltered[i], person)
		}
	}
}
