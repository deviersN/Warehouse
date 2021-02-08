package main

import (
	"fmt"
)

func dataChecker(warehouse Warehouse) (bool) {
	var ret int = 0
	var size Point = warehouse.entrepot.Point

	fmt.Println(size)
	ret += checkEntrepot(warehouse)
	ret += checkColis(warehouse, size)

	switch ret {
	case 0:
		return true
	default:
		return false
	}
}

func checkColis(warehouse Warehouse, size Point) (ret int) {
	ret = 0
	var overlap = []Point{}

	for _, v := range warehouse.colis {
		if v.x < 0 || v.x >= size.x {
			printError(2)
			ret = 1
		}
		capColor := myStrCapitalize(v.color)
		if capColor != "GREEN" && capColor != "YELLOW" && capColor != "BLUE" {
			printError(3)
			ret = 1
		}
		for _, v2 := range overlap {
			if (v.x == v2.x && v.y == v2.y) {
				printError(4)
				ret = 1
			}
		}
		overlap = append(overlap, v.Point)
	}
	return
}

func checkEntrepot(warehouse Warehouse) (ret int) {
	ret = 0
	if warehouse.entrepot.x < 2 || warehouse.entrepot.y < 2 {
		printError(0)
		ret = 1
	}
	if warehouse.entrepot.turns <= 10 || warehouse.entrepot.turns >= 100000 {
		printError(1)
		ret = 1
	}
	return ret
}

func printError(index int) {
	msg := [5]string {"Error: warehouse is not big enough.",
		"Error: number of turns must be between 10 and 100 000.",
		"Error: package out of the warehouse.",// SCP-Package has breached confineemnt. Iniciating facility-wide lockdown.
		"Error: package is wrong color.",// U just go Jim Crow lawed //Apartheid intersifies, must separate.
		"Error: packages overlapping."}

	fmt.Println(msg[index])
}
