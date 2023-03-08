package instance

import (
	"math"
	"math/rand"
)

type Exon struct {
	Value uint16 `json:"value"`
}

// # EXPERIMENTAL #
// Cross generates a child by cross overing another one
func (e *Exon) Cross(ee *Exon) (child Exon) {
	if e.Value > ee.Value {
		child.Value = e.Value - uint16(rand.Intn(int(e.Value-ee.Value)))
	} else {
		child.Value = e.Value + uint16(rand.Intn(int(ee.Value-e.Value)))
	}

	return
}

// Mutate generates a child by mutation
func (e *Exon) Mutate(max int) (child Exon) {
	mutation := uint16(rand.Intn(max))

	if rand.Intn(2) == 0 {
		// Mutate by decrementing
		if mutation > e.Value {
			child.Value = 0
		} else {
			child.Value = e.Value - mutation
		}
	} else {
		// Mutate by incrementing
		if math.MaxUint16-mutation < e.Value {
			child.Value = math.MaxUint16
		} else {
			child.Value = e.Value + mutation
		}
	}

	return
}
