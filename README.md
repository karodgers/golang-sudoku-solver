Sudoku Solver

Language: Go

This program is designed to solve valid sudoku puzzles. A valid sudoku puzzle is one that has only one possible solution.

How to Use

Step 1: Installation

Ensure you have Go installed on your system. If not, download and install it from Go's official website.

Step 2: Clone the Repository

Clone this repository to your local machine using the following command:

bash

git clone <repository_url>

Step 3: Navigate to the directory

Navigate to the directory where you cloned the repository. 

Step 4: Run the Program

To run the program, use the following command:

go run . <sudoku_string_1> <sudoku_string_2> ... <sudoku_string_9>

Replace <sudoku_string_X> with the strings representing the rows of the sudoku puzzle. Use periods . to represent empty cells.

Examples
Example 1: Valid Sudoku

bash

go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"

Output:

3 9 6 2 4 5 7 8 1
1 7 8 3 6 9 5 2 4
5 2 4 8 1 7 3 9 6
2 8 7 9 5 1 6 4 3
9 3 1 4 8 6 2 7 5
4 6 5 7 2 3 9 1 8
7 1 2 6 3 8 4 5 9
6 5 9 1 7 4 8 3 2
8 4 3 5 9 2 1 6 7

Example 2: Invalid Inputs

go run . 1 2 3 4

Output:

Error

go run .

Output:

Error

bash

go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"

Output:

Error

Note

Ensure that each sudoku string has exactly 9 characters, representing each cell in the row. Use periods . to represent empty cells.