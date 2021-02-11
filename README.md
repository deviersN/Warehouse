# Warehouse
This project emulates the internal management of a warehouse, where packages must be moved to trucks to be shipped using forklifts. The point of the exercise is to take in account the weight and distance of each package to optimize the trips and fill the truck.


## Compilation and execution

This project realised in Go is compiled as following :
$ go build

Genering an executable called "Warehouse"

To run it, enter on the terminal :
$ ./Warehouse [entry.txt]

"entry.txt" in a file containing the specs of the warehouse.

 - The first line gives out its size and te time allowed to the program.
"5 5 1000" would mean that the warehouse is 5 slots wide and long, and that the user has 1000 moves to empty it.

 - The second line and to forth specifies the name of the packages, their coordinates and color which defines their
 weight. Yellow for 100kg, Green for 200kg and Blue for 500kg.
"momzspaghetti 2 1 green" is a 200kg package called momzspaghetti located at coordinates 2 - 1.

 - The lines after list the packages determine the forklifts names and spawning locations.
"forklift_1 0 0"

 - Finally, the truck's loading point, maximum load and reload timeout appear on the last line.
"3 4 4000 5"


## Project organisation


### Arguments handling

The file given as argument is read by the datareader.go file.
The nature and content of each line will be verified to make sure that no invalid entry will be input in the program.
Once the format has been verified, all the information is stored in the project's data structures.

Then the datachecker.go file is being charged of verifying the plausability of the input data : are the package's coordinates within the walls of the warehouse, is the delivery point of the truck alongside the walls ?

Once the checking is done, the true algorithm may begin.