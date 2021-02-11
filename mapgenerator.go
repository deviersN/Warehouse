package main

import (
//	"fmt"
)

/*
** 0: empty case
** 1: truck's loading point
** 2: transpalette
** 3: colis
*/
func buildMap(warehouse Warehouse) ([][]int8) {
	view := make([][]int8, 5)
	for i := range view {
		view[i] = make([]int8, 5)
	}
	for _, v := range warehouse.colis {
		view[v.y][v.x] = 3
	}
	for _, v := range warehouse.transp {
		view[v.y][v.x] = 2
	}
	view[warehouse.camion.y][warehouse.camion.x] = 1
	return view
}