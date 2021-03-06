# Warehouse

This project emulates the internal management of a warehouse, where packages must be moved to trucks to be shipped using
forklifts. The point of the exercise is to take in account the weight and distance of each package to optimize the trips
and fill the truck.


## Compilation and execution

This project realised in Go is compiled as following :
$ go build

Genering an executable called "Warehouse"

To run it, enter on the terminal :
$ ./Warehouse ./resources/[entry.txt]

"entry.txt" in a file containing the specs of the warehouse.

- The first line gives out its size and the time allowed to the program.
  "5 5 1000" would mean that the warehouse is 5 slots wide and long, and that the user has 1000 moves to empty it.

- The second line and to forth specifies the name of the packages, their coordinates and color which defines their
  weight. Yellow for 100kg, Green for 200kg and Blue for 500kg.
  "momzspaghetti 2 1 green" is a 200kg package called momzspaghetti located at coordinates 2 - 1.

- The lines after list the packages determine the forklifts names and spawning locations.
  "forklift_1 0 0"

- Finally, the truck's loading point, maximum load and reload timeout appear on the last line.
  "3 4 4000 5" represents a truck whose loading coordinates are on 3-4, can wear up to 4 tons of packages and comes back 5 turns after being sent.


## Docker : compilation and execution

In order to deploy and execute the project in Docker containers, simply follow some steps listed below.

If, unfortunately, you must have "administrator" rights to perform these commands, I invite you to add your user to the
docker group of your machine.

```sh
      sudo usermod -a -G docker <USER_NAME>
```

The first thing you will need to do is to create the network in which your container will evolve. You must be in the
folder where the application is located, in our case « Warehouse ».

```sh
      docker network create warehouse-network
```

Execute the following command to build the Docker container.

```sh
      docker-compose build
```

Then to finish, execute a last command to launch the Docker container.

```sh
      docker-compose up
```

Make sure that an instruction file named **instructions.txt** is present in the [resources](./resources) folder so that
the program launches correctly.

To check that the docker container is launched correctly, you can run the following command:

```sh
      docker ps
```

To stop the running container, you can run the following command:

```sh
      docker-compose stop
```

To delete the stopped container, you can run the following command:

```sh
      docker container prune
```

To delete the dangling images, you can run the following command:

```sh
      docker image prune
```

To delete a specific image, simply list the images and delete the image using its name or identifier:

```sh
      docker image ls
      docker rmi <IMAGE_NAME / IMAGE_ID>
```


## Project organisation

### Main sections

main.go : Declares the basic data structures, calls the various sections of the program.

datareader.go : Reads the file given as argument, identificates each input to store it in the data structure.

datachecker.go : Checks the validity of the data stored

mapgenerator.go : Generates a map representing the placement of the items in the warehouse.

algorithm.go : Actual process of moving the packages around the warehouse.


### Secondary files

output.go : Regroups the prints all around the project

tools.go : Regroups a few basic functions

structure.go : Regroups the structures used in the data storing unit.


## Technical documentation

### Arguments handling

The file given as argument is read by the datareader.go file. The nature and content of each line will be verified to
make sure that no invalid entry will be input in the program. Once the format has been verified, all the information is
stored in the project's data structures.

Then the datachecker.go file is being charged of verifying the plausability of the input data : are the package's
coordinates within the walls of the warehouse, is the delivery point of the truck alongside the walls ?

Once the checking is done, the true algorithm may begin.


### Structures used

Two main structures are used in the process of the program : the map, and the data storage.

The map is a simple bidimentional int8 arrayshowing the disposition of the packages, forklifts and truck's loading point.

The data storage is a complex structure of structures integrating all the data stored in the instruction file given as argument, plus several points needed for the management of the different objects.
There is :
- The Warehouse object containing all the other structures
      - The Entrepot structure with the size of the warehouse and the number of turns allowed to the process.
      - An array of Colis structures holding the data of the different packages: their name, coordinates, color which defines the weight and an indicator.
      - An array of Transpalette structures holding the data of the different forklifts: their name, coordinates, coordinates of their target and the index of the potential package loaded in.
      - The Camion structure representing the truck: it holds its loading coordinates, the maximum load that can be held by the truck, the time the truck takes to come back empty in number of turns and the sum of the weight of the packages already inside the truck.


### Core algorithm

As long as any of the end-of-game conditions are met, the program goes on.
First, we check the forklifts objective. If they are already aiming for a target - whether it is a package or the truck's loading zone - they will move 1 step towards it and announce it by printing its move.
However, if no target has been set yet the forklift will find the closest package and lock onto it in order to go pick it up. To do that, it determines its distance to each package using Pythagoras' theorem.


### Unit Testing

In order to test the validity of the parser, we made some unit testing with variable cases of wrong and good datas. Those testing will check in depth the functions and what they return. 