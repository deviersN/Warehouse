package main

import (
    "fmt"
	"os"
)

func main() {
	ret := true
	argv := os.Args

	fmt.Println(argv)
	if len(argv) >= 2 {
		ret = dataReader(argv[1])
	}
	fmt.Println(ret)
}