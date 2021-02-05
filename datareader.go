package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func dataReader(filename string) (bool) {
    var warehouse Warehouse

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

        warehouse = dataStorer(warehouse, lineType, text)
        //	fmt.Sscanf(text, "%dx%dx%d", &nb1, &nb2, &nb3)
	}
    fmt.Println(warehouse)
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

func dataStorer(warehouse Warehouse, lineType int, line string) (Warehouse) {
    switch lineType {
    case 1:
        fmt.Sscanf(line, "%d %d %d", &warehouse.entrepot.x, &warehouse.entrepot.y, &warehouse.entrepot.turns)
    case 2:
        var colis Colis
        fmt.Sscanf(line, "%s %d %d %s", &colis.name, &colis.x, &colis.y, &colis.color)
        warehouse.colis = append(warehouse.colis, colis)
    case 3:
        var transpalette Transpalette
        fmt.Sscanf(line, "%s %d %d", &transpalette.name, &transpalette.x, &transpalette.y)
        warehouse.transp = append(warehouse.transp, transpalette)
    case 4:
        fmt.Sscanf(line, "%d %d %d %d", &warehouse.camion.x, &warehouse.camion.y, &warehouse.camion.load, &warehouse.camion.cooldown)
    default:
        fmt.Println("Oulalah no")
    }
    return warehouse
}