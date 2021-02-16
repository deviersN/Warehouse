package main

import (
	"fmt"
)

func core(warehouse Warehouse, view [][]int8) {
	turn := 1
	for gameOn(warehouse, turn) == true {
		for v := range warehouse.transp {
			fmt.Println(v, warehouse.transp[v].target)
			if warehouse.transp[v].target.x == -1 { // no target
				warehouse.transp[v] = lockOnTarget(warehouse.transp[v], warehouse)
			}
			warehouse, view = moveTransp(warehouse, view)
		}
		turn = turn + 1
		printMap(view)
	}
}

func pickPackage(colis []Colis, loc Point) (id int) {
	id = -1
	for _, v := range colis {
		if compareCoords(v.Point, loc) == true {
			id = v.id
		}
	}
	return
}

func moveTransp(warehouse Warehouse, view [][]int8) (Warehouse, [][]int8) {
	for v := range warehouse.transp {
		if beside(warehouse.transp[v].Point, warehouse.transp[v].target) {
			if compareCoords(warehouse.transp[v].target, warehouse.camion.Point) == true {
				// leave
			} else {
				id := pickPackage(warehouse.colis, warehouse.transp[v].target)
				warehouse.transp[v].loaded = id
				view[warehouse.transp[v].target.y][warehouse.transp[v].target.x] = 0
				warehouse.transp[v].target = warehouse.camion.Point
				// take
			}
		} else {
			if warehouse.transp[v].y != warehouse.transp[v].target.y {
				if warehouse.transp[v].y > warehouse.transp[v].target.y {
					view[warehouse.transp[v].y][warehouse.transp[v].x] = 0
					view[warehouse.transp[v].y - 1][warehouse.transp[v].x] = 2
					warehouse.transp[v].y -= 1
					//Go up
				} else {
					view[warehouse.transp[v].y][warehouse.transp[v].x] = 0
					view[warehouse.transp[v].y + 1][warehouse.transp[v].x] = 2
					warehouse.transp[v].y += 1
					//Go down
				}
			} else if warehouse.transp[v].x != warehouse.transp[v].target.x {
				if warehouse.transp[v].x > warehouse.transp[v].target.x {
					view[warehouse.transp[v].y][warehouse.transp[v].x - 1] = 0
					view[warehouse.transp[v].y][warehouse.transp[v].x] = 2
					warehouse.transp[v].x += 1
					//Go left
				} else {
					view[warehouse.transp[v].y][warehouse.transp[v].x] = 0
					view[warehouse.transp[v].y][warehouse.transp[v].x + 1] = 2
					warehouse.transp[v].x += 1
					//Go right
				}
			}
		}
	}
	return warehouse, view
}

func beside(src Point, dest Point) (ret bool) {
	ret = false
	x := absolute(src.x - dest.x)
	y := absolute(src.y - dest.y)
	if x + y == 1 {
		ret = true
	}
	return
}

func lockOnTarget(transp Transpalette, warehouse Warehouse) (Transpalette) {
	closest := 1000000
	for v := range warehouse.colis {
		x := absolute(transp.x - warehouse.colis[v].x)
		y := absolute(transp.y - warehouse.colis[v].y)
		dist := pythagoras(float64(x), float64(y))
		if dist < closest {
			closest = dist
			transp.target.x = x
			transp.target.y = y
			fmt.Println(transp)
		}
	}
	return transp
}

func gameOn(warehouse Warehouse, turn int) (ret bool) {
	ret = true
	if warehouse.entrepot.turns - turn <= 0 {
		ret = false
	}
	if len(warehouse.colis) <= 0 {
		ret = false
	}
	return
}