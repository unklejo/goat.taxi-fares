# Taxi Fare Calculator

## User Requirements

This project is a taxi fare calculator that reads distance and time records, calculates the fare based on given rules, and outputs the fare along with the processed records.
Below are the user requirements for this project:
### (i) Overview

1. The base fare is 400 yen for up to 1 km.
2. Up to 10 km, 40 yen is added every 400 meters.
3. Over 10km, 40 yen is added every 350 meters.

This taxi is equipped with the following two meters. Only one of the most recent real values is
recorded on these meters.

- Distance Meter
- Fare Meter

### (ii) Input Format

Distance meter records are sent line by line for standard input in the following format.

00:00:00.000 0.0
00:01:00.123 480.9
00:02:00.125 1141.2
00:03:00.100 1800.8

The specifications of the distance meter are as follows.

- Space-separated first half is elapsed time (Max 99:99:99.999), second half is mileage.(the unit is meters, Max: 99999999.9)
- It keeps only the latest values.
- It calculates and creates output of the mileage per minute, but an error of less than 1
second occurs.

### (iii) Error Definition

Error occurs under the following conditions.

- Not under the format of, hh:mm:ss.fff<SPACE>xxxxxxxx.f<LF>, but under an improper
format.
- Blank Line
- When the past time has been sent.
- The interval between records is more than 5 minutes apart.
- When there are less than two lines of data.
- When the total mileage is 0.0m.

### (iv) Output

Display the current fare as an integer on the fare meter (standard output).
12345

On the next lines, display all of the input data with mileage difference compared to previous data, and order it from highest to lowest

00:02:00.125 1141.2 660.3
00:03:00.100 1800.8 659.6
00:01:00.123 480.9 480.9
00:00:00.000 0.0 0.0

Standard output displays nothing for incorrect inputs that do not meet specifications, the exit code ends with a value other than 0.

## Build Requirements

- Go 1.22 or higher

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

## Running The Tests

To validate the functionality of the program, run the following command in the terminal:
```
go test -v ./...
```
This command will run all tests in the project and display detailed information about each test executed.

## Check the test coverage

To view the coverage of the projects run the following command in the terminal:
```
go test -cover ./...
```
