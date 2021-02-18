package main

import (
	"testing"
)

func TestEntrepot(t *testing.T) {
	var warehouse Warehouse
	var entrepot Entrepot = Entrepot{Point{5, 5}, 15}
	data := "resources/instructions.txt"
	_, warehouse = dataReader(data)
	if warehouse.entrepot != entrepot {
		t.Errorf("L'entrepot lu n'est pas bon")
	}

}

func TestCamion(t *testing.T) {
	var warehouse Warehouse
	var camion Camion = Camion{Point{3, 4}, 4000, 5, 0}
	data := "resources/instructions.txt"
	_, warehouse = dataReader(data)
	if warehouse.camion != camion {
		t.Errorf("Le camion lu n'est pas bon")
	}

}

func TestColis(t *testing.T) {
	var warehouse Warehouse
	colis := [4]Colis{Colis{"colis_a_livrer", Point{2,1}, "green", 1},
		Colis{"paquet", Point{2,2}, "BLUE", 2},
		Colis{"deadpool", Point{0,3}, "yellow", 3},
		Colis{"colère_DU_dragon", Point{4,1}, "green", 4}}
	data := "resources/instructions.txt"
	_, warehouse = dataReader(data)
	for i, el := range warehouse.colis {
		if el != colis[i] {
			t.Errorf("Le colis %s n'est pas bon", el.name)
		}
	}
}

func TestTranspalette(t *testing.T) {
	var warehouse Warehouse
	trans := [1]Transpalette{{"transpalette_1", Point{0,0}, Point{-1,-1}, 0}}
	data := "resources/instructions.txt"
	_, warehouse = dataReader(data)
	for i, el := range warehouse.transp {
		if el != trans[i] {
			t.Errorf("Le transpalette %s n'est pas bon", el.name)
		}
	}
}

func TestIdentificatorEntrepot(t *testing.T) {
	s := "6 5 4000"
	tmp := identificator(s)
	if tmp != 1 {
		t.Errorf("Identificator was incorrect said : %d", tmp)
	}
}

func TestIdentificatorColis(t *testing.T) {
	s := "colis_privé 5 1 yellow"
	tmp := identificator(s)
	if tmp != 2 {
		t.Errorf("Identificator was incorrect said : %d", tmp)
	}
}

func TestIdentificatorTranspalette(t *testing.T) {
	s := "transpalette_3 0 0"
	tmp := identificator(s)
	if tmp != 3 {
		t.Errorf("Identificator was incorrect said : %d", tmp)
	}
}

func TestIdentificatorCamion(t *testing.T) {
	s := "2 2 4000 6"
	tmp := identificator(s)
	if tmp != 4 {
		t.Errorf("Identificator was incorrect said : %d", tmp)
	}
}

func TestInvalidEntryEmptyFile(t *testing.T)  {
	invalidFile := "invalid_data_empty.txt"
	var warehouse Warehouse
	_, warehouse = dataReader(invalidFile)
	if dataChecker(warehouse) != false {
		t.Errorf("The file was invalid, datareader should return false")
	}
}

func TestInvalidEntryNoTurn(t *testing.T)  {
	invalidFile := "invalid_data_noturn.txt"
	var warehouse Warehouse
	_, warehouse = dataReader(invalidFile)
	if dataChecker(warehouse) != false {
		t.Errorf("The file was invalid, datareader should return false")
	}
}

func TestInvalidEntryNoFirstLine(t *testing.T)  {
	invalidFile := "invalid_data_nofirstline.txt"
	var warehouse Warehouse
	_, warehouse = dataReader(invalidFile)
	if dataChecker(warehouse) != false {
		t.Errorf("The file was invalid, datareader should return false")
	}
}

func TestInvalidEntryNoColis(t *testing.T)  {
	invalidFile := "invalid_data_nocolis.txt"
	var warehouse Warehouse
	_, warehouse = dataReader(invalidFile)
	if dataChecker(warehouse) != false {
		t.Errorf("The file was invalid, datareader should return false")
	}
}