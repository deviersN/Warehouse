package main

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
    entrepot Entrepot
    colis []Colis
    transp []Transpalette
    camion Camion
}