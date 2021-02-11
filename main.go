package main

import (
	"fmt"
	"os"
)

func main() {
	ret := true
	argv := os.Args
	var warehouse Warehouse
	var view [][]int8

//	fmt.Println(argv)
	if len(argv) == 2 {
		ret, warehouse = dataReader(argv[1])
		if (ret == false) {
			return
		}
		ret = dataChecker(warehouse)
		fmt.Println(warehouse)
		if (ret == false) {
			return
		}
		view = buildMap(warehouse)
		fmt.Println(view)
	}
}