package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

type Warehouse struct {
    width int
    height int
    timeLeft int
}

type isFileComplete struct {
    warehouse bool
    colis int
    transpalette int
    camion bool
}

func dataReader(filename string) bool{
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    text := scanner.Text()
    dataStorer(text)
/*
    for scanner.Scan() {
		text := scanner.Text()
        fmt.Println(text)
        dataStorer(text)
        //	fmt.Sscanf(text, "%dx%dx%d", &nb1, &nb2, &nb3)
	}*/

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
        return false
    }
    return true
}

func dataStorer(line string) bool {
    var w, h, t int

    fmt.Sscanf(line, "%d %d %d", &w, &h, &t)
    fmt.Println("->", w, h, t)
    return true
}