package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

type Point struct {
    x int
    y int
}

type Entrepot struct {
    Point
    turns int
}

type Colis struct {
    name string
    Point
    color string
}

type Transpalette struct {
    name string
    Point
}

type Camion struct {
    Point
    load int
    cooldown int
}

type Warehouse struct {
    Entrepot
    colis []Colis
    transp []Transpalette
    Camion
}

func dataReader(filename string) (bool) {
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		text := scanner.Text()
        fmt.Println(text)

        lineType := identificator(text)
        fmt.Println("id: ", lineType)

        dataStorer(text)
        //	fmt.Sscanf(text, "%dx%dx%d", &nb1, &nb2, &nb3)
	}

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
        return false
    }
    return true
}

/*
tp values:
0 Empty line
1 Entrepot
2 Colis
3 Transpalette
4 Camion
5 Erreur
*/
/*
isNum value: 0 for num, 1 for alphanum
*/
func identificator(line string) (tp int) {
    tp = 0 
    words := 1
    var isNum [4]int
    for c, v := range line {
        if v != ' ' && (v < '0' || v > '9') {
            isNum[words-1] = 1
        }
        if (v == ' ' && c != 0 && line[c - 1] != ' ') {
            words ++
        }
    }
    if words == 3 && compareArray(isNum, [4]int{0, 0, 0, 0}, words) {
        tp = 1 //Entrepot
    } else if words == 4 && compareArray(isNum, [4]int{1, 0, 0, 1}, words) {
        tp = 2 //Colis
    } else if words == 3 && compareArray(isNum, [4]int{1, 0, 0, 0}, words) {
        tp = 3 //Transpalette
    } else if words == 4 && compareArray(isNum, [4]int{0, 0, 0, 0}, words) {
        tp = 4 //Camion
    } else if len(line) == 0 {
        tp = 0 //Empty line
    } else {
        tp = 5 //Error
    }
    return
}

func compareArray(array_1 [4]int, array_2 [4]int, size int) (ret bool) {
    ret = true
    for i := 0; i < size; i++ {
        if array_1[i] != array_2[i] {
            ret = false
        }
    }
    return
}

func dataStorer(line string) (bool) {
    var w, h, t int

    fmt.Sscanf(line, "%d %d %d", &w, &h, &t)
    fmt.Println("->", w, h, t)
    return true
}