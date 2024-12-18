package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	persons, err := loadPersons("persons.json")
	if err != nil {
		panic(err)
	}
	persons.api = NewMockExchangeRateAPI()
	convertedPersons, err := persons.ConvertSalariesToDollar()
	if err != nil {
		log.Fatalf("Error converting salaries to USD: %v", err)
	}

	fmt.Println("\nSorting by ascending salary (in USD):")
	convertedPersons.SortBySalaryAsc()
	for _, person := range convertedPersons.Data {
		fmt.Printf("%+v\n", person)
	}

	fmt.Println("\nSorting by descending salary (in USD):")
	convertedPersons.SortBySalaryDesc()
	for _, person := range convertedPersons.Data {
		fmt.Printf("%+v\n", person)
	}

	fmt.Println("\nFiltering persons with salary between 20 and 50 USD:")
	filtered := convertedPersons.FilterBySalaryCriteria(20, 50)
	for _, person := range filtered {
		fmt.Printf("%+v\n", person)
	}

}

// loadPersons load Persons.Data from filename.
// Note that filename must point to a json file in the local filesystem
func loadPersons(filename string) (*Persons, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Unmarshal JSON data
	var persons Persons
	err = json.Unmarshal(bytes, &persons)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return &persons, nil
}
