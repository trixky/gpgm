package instance

import (
	"math/rand"
)

type Gene struct {
	// Process
	ProcessId uint16
	// Minimum quantity
	MinQuantity       uint16
	MinQuantityActive bool
	// Maximum quantity
	MaxQuantity       uint16
	MaxQuantityActive bool
}

// Mutate generates a child by mutation
func (g *Gene) Mutate(process_max uint16, process_shift int, quantity_shift int, activation_chance int) (child Gene) {
	var shift uint16

	// ---------------------- ProcessId
	shift = uint16(rand.Intn(process_shift))

	if rand.Intn(2) == 0 {
		// Substract
		if shift > g.ProcessId {
			// Prevent overflow
			child.ProcessId = 0
		} else {
			child.ProcessId = g.ProcessId - shift
		}
	} else {
		// Addition
		if process_max-shift < g.ProcessId {
			// Prevent overflow
			child.ProcessId = process_max - 1
		} else {
			child.ProcessId = g.ProcessId + shift
		}
	}

	// ---------------------- MinQuantity
	shift = uint16(rand.Intn(quantity_shift))

	if rand.Intn(2) == 0 {
		// Substract
		if quantity_shift > int(g.MinQuantity) {
			// Prevent overflow
			child.MinQuantity = 0
		} else {
			child.MinQuantity = g.MinQuantity - shift
		}
	} else {
		// Addition
		if shift > g.MaxQuantity-shift || g.MaxQuantity-shift < g.MinQuantity {
			// Prevent overflow
			child.MinQuantity = g.MaxQuantity
		} else {
			child.MinQuantity = g.MinQuantity + shift
		}
	}

	// ---------------------- MaxQuantity
	shift = uint16(rand.Intn(int(quantity_shift)))

	if rand.Intn(2) == 0 {
		// Substract
		if shift > g.MaxQuantity-shift || g.MaxQuantity-shift < g.MinQuantity {
			// Prevent overflow
			child.MaxQuantity = g.MinQuantity
		} else {
			child.MaxQuantity = g.MaxQuantity - shift
		}
	} else {
		// Addition
		// math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16
		// math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16
		// math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16
		// math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16
		// math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16
		if 3-shift < g.MaxQuantity {
			// Prevent overflow
			child.MaxQuantity = 3 - shift
		} else {
			child.MaxQuantity = g.MaxQuantity - shift
		}
	}

	// ---------------------- MinQuantityActive
	if rand.Intn(activation_chance) == 0 {
		child.MinQuantityActive = !g.MinQuantityActive
	}

	// ---------------------- MaxQuantityActive
	if rand.Intn(activation_chance) == 0 {
		child.MaxQuantityActive = !g.MaxQuantityActive
	}

	return
}
