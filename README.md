# Math Skills

A Go program that reads a file of integers (one per line) and prints basic statistics: average, median, variance and standard deviation.

## Features

- Reads integers from a text file (one number per line)
- Computes four statistics:

| Statistic          | Description                                      |
|--------------------|--------------------------------------------------|
| Average            | Sum of all numbers divided by their count        |
| Median             | Middle value of the sorted numbers               |
| Variance           | Population variance (divides by N)               |
| Standard Deviation | Square root of the variance                      |

- Results are rounded to the nearest integer

## Requirements

- Go 1.22.2 or later

## Usage

```bash
go run . <file.txt>
```

Example:

```bash
go run . example.txt
```

## Input format

- One integer per line
- Negative numbers are allowed
- Empty lines are ignored

Example (`example.txt`):

```
50
10
30
20
40
```

## Output

```
Average: 30
Median: 30
Variance: 200
Standard Deviation: 14
```

The numbers do not need to be sorted in the file: the program sorts them before computing the median.

## How it works

1. The file is read and split line by line into integers
2. The numbers are sorted (bubble sort)
3. Each statistic is computed from the sorted list
4. Each result is rounded to the nearest integer and printed

## Errors

- If no file (or more than one) is given, the program prints `need a file with numbers`
- If the file contains no numbers, the program prints `need numbers in file`
- If a line is not a valid integer, the program prints `not a number`
- If the file cannot be read, the error from Go is printed as is

## Tests

```bash
go test ./...
```

## Project structure

```
.
├── example.txt
├── go.mod
├── main.go
├── README.md
└── tools
    ├── mathtools.go
    ├── stats_test.go
    ├── utils.go
    └── utils_test.go
```

## Author

swannsavary
