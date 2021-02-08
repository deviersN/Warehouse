package main

import (
	"os"
)

func main() {
	ret := true
	argv := os.Args
    var warehouse Warehouse

//	fmt.Println(argv)
	if len(argv) >= 2 {
		ret, warehouse = dataReader(argv[1])
//		fmt.Println(ret)
		if (ret == false) {
			return
		}
		ret = dataChecker(warehouse)
		if (ret == false) {
			return
		}
	}
}