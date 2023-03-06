package instance

import (
	"math"
	"math/rand"
)

type Gene struct {
	Value uint16
}

// # EXPERIMENTAL #
// Cross generates a child by cross overing another one
func (g *Gene) Cross(gg *Gene) (child Gene) {
	if g.Value > gg.Value {
		child.Value = g.Value - uint16(rand.Intn(int(g.Value-gg.Value)))
	} else {
		child.Value = g.Value + uint16(rand.Intn(int(gg.Value-g.Value)))
	}

	return
}

// Mutate generates a child by mutation
func (g *Gene) Mutate(max int) (child Gene) {
	mutation := uint16(rand.Intn(max))

	if rand.Intn(2) == 0 {
		// Mutate by decrementing
		if mutation > g.Value {
			child.Value = 0
		} else {
			child.Value = g.Value - mutation
		}
	} else {
		// Mutate by incrementing
		if math.MaxUint16-mutation < g.Value {
			child.Value = math.MaxUint16
		} else {
			child.Value = g.Value + mutation
		}
	}

	return
}
