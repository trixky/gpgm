package checker

import (
	"testing"
)

func TestSimple(t *testing.T) {
	res, err := CheckOutput(
		"euro:10\nachat_materiel:(euro:8):(materiel:1):10\nrealisation_produit:(materiel:1):(produit:1):30\nlivraison:(produit:1):(client_content:1):20\noptimize:(time;client_content)",
		"0: achat_materiel:1\n10: realisation_produit:1\n40: livraison:1\nstock: euro:2;materiel:0;produit:0;client_content:1",
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if res != true {
		t.Errorf("Expected result to be true but got false")
	}
}
