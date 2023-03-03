package parser

import (
	"testing"
)

func TestSimple(t *testing.T) {
	res, err := ParseSimulationFile("euro:10\nachat_materiel:(euro:8):(materiel:1):10\nrealisation_produit:(materiel:1):(produit:1):30\nlivraison:(produit:1):(client_content:1):20\noptimize:(time;client_content)")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if res.Stock["euro"] != 10 {
		t.Errorf("Failed to set Stock")
	}
	if res.Stock["materiel"] != 0 {
		t.Errorf("Failed to set Stock from Inputs")
	}
	if res.Stock["produit"] != 0 {
		t.Errorf("Failed to set Stock from Outputs")
	}
	if res.Stock["client_content"] != 0 {
		t.Errorf("Failed to set Stock from Outputs")
	}
	if len(res.Processes) != 3 {
		t.Errorf("Failed to add all Processes")
	}
	if len(res.Optimize) != 2 {
		t.Errorf("Failed to add all Optimize")
	}
}

func TestPomme(t *testing.T) {
	_, err := ParseSimulationFile("#\n#  krpsim tarte aux pommes\n#\nfour:10\neuro:10000\n#\nbuy_pomme:(euro:100):(pomme:700):200\nbuy_citron:(euro:100):(citron:400):200\nbuy_oeuf:(euro:100):(oeuf:100):200\nbuy_farine:(euro:100):(farine:800):200\nbuy_beurre:(euro:100):(beurre:2000):200\nbuy_lait:(euro:100):(lait:2000):200\n#\nseparation_oeuf:(oeuf:1):(jaune_oeuf:1;blanc_oeuf:1):2\nreunion_oeuf:(jaune_oeuf:1;blanc_oeuf:1):(oeuf:1):1\ndo_pate_sablee:(oeuf:5;farine:100;beurre:4;lait:5):(pate_sablee:300;blanc_oeuf:3):300\ndo_pate_feuilletee:(oeuf:3;farine:200;beurre:10;lait:2):(pate_feuilletee:100):800\ndo_tarte_citron:(pate_feuilletee:100;citron:50;blanc_oeuf:5;four:1):(tarte_citron:5;four:1):60\ndo_tarte_pomme:(pate_sablee:100;pomme:30;four:1):(tarte_pomme:8;four:1):50\ndo_flan:(jaune_oeuf:10;lait:4;four:1):(flan:5;four:1):300\ndo_boite:(tarte_citron:3;tarte_pomme:7;flan:1;euro:30):(boite:1):1\nvente_boite:(boite:100):(euro:55000):30\nvente_tarte_pomme:(tarte_pomme:10):(euro:100):30\nvente_tarte_citron:(tarte_citron:10):(euro:200):30\nvente_flan:(flan:10):(euro:300):30\n#do_benef:(euro:1):(benefice:1):0\n#\n#\n#optimize:(benefice)\noptimize:(euro)\n#")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestNoProcess(t *testing.T) {
	_, err := ParseSimulationFile("euro:10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestNoOptimize(t *testing.T) {
	_, err := ParseSimulationFile("euro:10\nachat_materiel:(euro:8):(materiel:1):10")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestNoStock(t *testing.T) {
	_, err := ParseSimulationFile("achat_materiel:(euro:8):(materiel:1):10\noptimize:(time)")
	if err != nil {
		t.Errorf("Expected error: %v", err)
	}
}

func TestInvalidStock1(t *testing.T) {
	_, err := ParseSimulationFile("thing:-10\nmake:(thing:8):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidStock2(t *testing.T) {
	_, err := ParseSimulationFile("thing:a\nmake:(thing:8):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidStock3(t *testing.T) {
	_, err := ParseSimulationFile("thing:\nmake:(thing:8):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidStock4(t *testing.T) {
	_, err := ParseSimulationFile("thing:12:\nmake:(thing:8):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidStock5(t *testing.T) {
	_, err := ParseSimulationFile("thing:12a\nmake:(thing:8):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestValidProcessNoInputs(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake::(stuff:1):10\noptimize:(time)")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestValidProcessNoOutputs(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8)::10\noptimize:(time)")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestInvalidProcessInput1(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing;8):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessInput2(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8;):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessInput3(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessInput4(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessInput5(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:a):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessInput6(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(:8):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessInput7(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(stuff:-8):(stuff:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessOutput1(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(stuff):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessOutput2(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(stuff:):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessOutput3(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(stuff:a):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessOutput4(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(stuff:1;):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessOutput5(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(:1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessOutput6(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(stuff;1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessOutput7(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(stuff:-1):10\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessDelay1(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(stuff:1):\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessDelay2(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(stuff:1):a\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessDelay3(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(stuff:1):-12\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInvalidProcessDelay4(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(stuff:1):1:\noptimize:(time)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}

func TestInexistingOptimizeProduct(t *testing.T) {
	_, err := ParseSimulationFile("thing:12\nmake:(thing:8):(stuff:1):1:\noptimize:(that)")
	if err == nil {
		t.Errorf("Expected error but got nothing")
	}
}
