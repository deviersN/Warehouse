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
	ret += checkTranspalette(warehouse, size)
	ret += checkCamion(warehouse, size)

	switch ret {
	case 0:
		return true
	default:
		return false
	}
}

func checkCamion(warehouse Warehouse, size Point) (ret int) {
	ret = 0
	var truck Point = warehouse.camion.Point
	if warehouse.camion.load <= 500 {
		printError(8)
	} else if warehouse.camion.load < 100 {
		printError(9)
		ret = 1
	}
	if truck.x != 0 && truck.y != 0 &&
		truck.x != size.x - 1 && truck.y != size.y - 1 {
		printError(10)
		ret = 1
	}
	return
}

func checkTranspalette(warehouse Warehouse, size Point) (ret int) {
	ret = 0
	var overlap = []Point{}

	for _, v := range warehouse.transp {
		if v.x < 0 || v.x >= size.x || v.y < 0 || v.y >= size.y {
			printError(5)
			ret = 1
		}
		for _, v2 := range warehouse.colis {
			if (v.x == v2.x && v.y == v2.y) {
				printError(6)
				ret = 1
			}
		}
		for _, v2 := range overlap {
			if (v.x == v2.x && v.y == v2.y) {
				printError(7)
				ret = 1
			}
		}
		overlap = append(overlap, v.Point)
	}
	return
}

func checkColis(warehouse Warehouse, size Point) (ret int) {
	ret = 0
	var overlap = []Point{}

	for _, v := range warehouse.colis {
		if v.x < 0 || v.x >= size.x || v.y < 0 || v.y >= size.y {
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
	return
}
