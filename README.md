# Taxi Fare Calculator

This project is a taxi fare calculator that reads distance and time records, calculates the fare based on given rules, and outputs the fare along with the processed records.

## Requirements

- Go 1.22 or higher

## Running The Tests

To validate the functionality of the program, run the following command in the terminal:
```
go test -v ./...
```
This command will run all tests in the project and display detailed information about each test executed.

## Building and Running the Program Manually

To build the program and run it manually, follow these steps:

1. **Prepare Input File:**
   Ensure you have an `input.txt` file with appropriate distance and time records. You can use the provided file in the project's root folder.

2. **Build the Program:**
   Run the following command in the root directory of the project:
```
go build -o taxi-fares cmd/main.go
```
This command compiles the program and creates an executable named `taxi-fares`.

3. **Run the Program:**
Execute the following command in the terminal to run the program:
```
./taxi-fares < fare-input.txt

```
Replace `input.txt` with the name of your input file if it has a different name. This command feeds the content of the input file into the program and displays the output.
