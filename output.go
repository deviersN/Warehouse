package main

import (
	"fmt"
)

func printTour(turn int) {
	fmt.Println("Tour", turn)
}

func printMove(dest Point, name string) {
	fmt.Println(name, "GO", dest)
}

func printActionPackage(transp string, action string, colis string, color string) {
	fmt.Println(transp, action, colis, color)
}

func printTruckStatus(action string, load int, loadIn int) {
	fmt.Println("camion", action, load, loadIn)
}

func printError(index int) {
	msg := [12]string {"Error: warehouse is not big enough.",
		"Error: number of turns must be between 10 and 100 000.",
		"Error: package out of the warehouse.",// SCP-Package has breached confinement. Iniciating facility-wide lockdown.
		"Error: package is wrong color.",// U just go Jim Crow lawed
		"Error: packages overlapping.",// Apartheid intersifies, must separate.
		"Error: transpalette out of the warehouse.",
		"Error: transpalette and package overlapping.",
		"Error: transpalettes overlapping.",
		"Warning: truck's load is very low.",
		"Error: truck's load is too low to hold any package.",
		"Error: truck's load point is in the center of the room or out of the warehouse."}

	fmt.Println(msg[index])
}
