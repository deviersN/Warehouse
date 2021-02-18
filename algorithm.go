package main

import (

)

func core(warehouse Warehouse, view [][]int8) {
	turn := 1
	printMap(view)
	for gameOn(warehouse, turn) == true {
		printTour(turn)
		for v := range warehouse.transp {
//			fmt.Println(v, warehouse.transp[v].target)
			if warehouse.transp[v].target.x == -1 { // no target
				warehouse.transp[v] = lockOnTarget(warehouse.transp[v], warehouse)
			}
			warehouse, view = moveTransp(warehouse, view)
			printTruckStatus("WAITING", warehouse.camion.load, warehouse.camion.loadIn)
		}
		turn = turn + 1
		printMap(view)
	}
}

func pickPackage(colis []Colis, transp Transpalette) (id int) {
	id = -1
	for _, v := range colis {
		if compareCoords(v.Point, transp.target) == true {
			printActionPackage(transp.name, "TAKE", v.name, myStrCapitalize(v.color))
			id = v.id
		}
	}
	return
}

func dropper(s []Colis, index int) ([]Colis) {
    return append(s[:index], s[index+1:]...)
}

func dropPackage(warehouse Warehouse, transp Transpalette) (Warehouse) {
	for i := range warehouse.colis {
		if warehouse.colis[i].id == transp.loaded {
			printActionPackage(transp.name, "LEAVE", warehouse.colis[i].name, myStrCapitalize(warehouse.colis[i].color))
			warehouse.camion.loadIn += weightWatcher(warehouse.colis[i].color)
			warehouse.colis = dropper(warehouse.colis, i)
			break
		}
	}
	return warehouse
}

func moveTransp(warehouse Warehouse, view [][]int8) (Warehouse, [][]int8) {
	for v := range warehouse.transp {
		if beside(warehouse.transp[v].Point, warehouse.transp[v].target) {
			if compareCoords(warehouse.transp[v].target, warehouse.camion.Point) == true {
				warehouse = dropPackage(warehouse, warehouse.transp[v])
				warehouse.transp[v].target = Point{-1, -1}
				// leave
			} else {
				id := pickPackage(warehouse.colis, warehouse.transp[v])
				warehouse.transp[v].loaded = id
				view[warehouse.transp[v].target.y][warehouse.transp[v].target.x] = 0
				warehouse.transp[v].target = warehouse.camion.Point
				// take
			}
		} else {
			if warehouse.transp[v].y != warehouse.transp[v].target.y {
				if warehouse.transp[v].y > warehouse.transp[v].target.y { //Go up
					view[warehouse.transp[v].y][warehouse.transp[v].x] = 0
					view[warehouse.transp[v].y - 1][warehouse.transp[v].x] = 2
					warehouse.transp[v].y -= 1
				} else {//Go down
					view[warehouse.transp[v].y][warehouse.transp[v].x] = 0
					view[warehouse.transp[v].y + 1][warehouse.transp[v].x] = 2
					warehouse.transp[v].y += 1
				}
			} else if warehouse.transp[v].x != warehouse.transp[v].target.x {
				if warehouse.transp[v].x > warehouse.transp[v].target.x { //Go left
					view[warehouse.transp[v].y][warehouse.transp[v].x] = 0
					view[warehouse.transp[v].y][warehouse.transp[v].x-1] = 2
					warehouse.transp[v].x -= 1
				} else { //Go right
					view[warehouse.transp[v].y][warehouse.transp[v].x] = 0
					view[warehouse.transp[v].y][warehouse.transp[v].x + 1] = 2
					warehouse.transp[v].x += 1
				}
			}
			printMove(warehouse.transp[v].Point, warehouse.transp[v].name)
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
//		fmt.Println(x, y, dist)
		if dist < closest {
			closest = dist
			transp.target.x = warehouse.colis[v].x
			transp.target.y = warehouse.colis[v].y
//			fmt.Println(transp)
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