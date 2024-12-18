# Persons

Persons manages a list of persons, each with a salary in various currencies. The program provides functionality to:

- Convert all salaries to USD using exchange rates.
- Sort the list of persons by their salaries in ascending or descending order.
- Filter the list of persons based on a salary range.

---

## Features

### 1. Salary Conversion
All salaries are converted from their original currency to USD using the exchange rates provided by the `ExchangeRateAPI`.

### 2. Sorting
- **Ascending Order**: Sorts persons by their salaries in ascending order.
- **Descending Order**: Sorts persons by their salaries in descending order.

### 3. Filtering
Filters persons based on a specified salary range in USD.

---

## Installation

1. Clone the repository:
   `git clone https://github.com/ercross/manipulator.git`

2. Navigate to the project directory:
   `cd persons-project`

3. Run the application:
   `go run main.go`

4. Run tests:
   `go test ./...`
