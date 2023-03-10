package instance

import (
	"strconv"

	"github.com/trixky/krpsim/algo/core"
)

type Dependences struct {
	InputsProcesses [][]int // [inputs (in order)][processes (in order | can be skipped)]
}

type Gene struct {
	History map[string]Dependences
	Process *core.Process
}

func (g *Gene) Init(process *core.Process, processes []core.Process) {
	const history_max_length = 3

	g.History = map[string]Dependences{}
	g.Process = process

	g.GetParentKeys(history_max_length, "", process, processes)
}

func (g *Gene) GetParentKeys(depth int, child_key string, process *core.Process, processes []core.Process) {
	depth--

	for _, process_parent := range process.Parents {
		key := child_key + "." + strconv.Itoa(process_parent)
		g.History[key] = Dependences{}

		if depth > 0 {
			g.GetParentKeys(depth, key, &processes[process_parent], processes)
		}
	}
}

func (g *Gene) Mutate(process_max uint16, process_shift int, quantity_shift int, activation_chance int) {

}

// // Mutate generates a child by mutation
// func (g *Gene) Mutate(process_max uint16, process_shift int, quantity_shift int, activation_chance int) (child Gene) {
// 	var shift uint16

// 	// ---------------------- ProcessId
// 	shift = uint16(rand.Intn(process_shift))

// 	if rand.Intn(2) == 0 {
// 		// Substract
// 		if shift > g.ProcessId {
// 			// Prevent overflow
// 			child.ProcessId = 0
// 		} else {
// 			child.ProcessId = g.ProcessId - shift
// 		}
// 	} else {
// 		// Addition
// 		if process_max-shift < g.ProcessId {
// 			// Prevent overflow
// 			child.ProcessId = process_max - 1
// 		} else {
// 			child.ProcessId = g.ProcessId + shift
// 		}
// 	}

// 	// ---------------------- MinQuantity
// 	shift = uint16(rand.Intn(quantity_shift))

// 	if rand.Intn(2) == 0 {
// 		// Substract
// 		if quantity_shift > int(g.MinQuantity) {
// 			// Prevent overflow
// 			child.MinQuantity = 0
// 		} else {
// 			child.MinQuantity = g.MinQuantity - shift
// 		}
// 	} else {
// 		// Addition
// 		if shift > g.MaxQuantity-shift || g.MaxQuantity-shift < g.MinQuantity {
// 			// Prevent overflow
// 			child.MinQuantity = g.MaxQuantity
// 		} else {
// 			child.MinQuantity = g.MinQuantity + shift
// 		}
// 	}

// 	// ---------------------- MaxQuantity
// 	shift = uint16(rand.Intn(int(quantity_shift)))

// 	if rand.Intn(2) == 0 {
// 		// Substract
// 		if shift > g.MaxQuantity-shift || g.MaxQuantity-shift < g.MinQuantity {
// 			// Prevent overflow
// 			child.MaxQuantity = g.MinQuantity
// 		} else {
// 			child.MaxQuantity = g.MaxQuantity - shift
// 		}
// 	} else {
// 		// Addition
// 		// math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16
// 		// math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16
// 		// math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16
// 		// math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16
// 		// math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16math.MaxUint16
// 		if 3-shift < g.MaxQuantity {
// 			// Prevent overflow
// 			child.MaxQuantity = 3 - shift
// 		} else {
// 			child.MaxQuantity = g.MaxQuantity - shift
// 		}
// 	}

// 	// ---------------------- MinQuantityActive
// 	if rand.Intn(activation_chance) == 0 {
// 		child.MinQuantityActive = !g.MinQuantityActive
// 	}

// 	// ---------------------- MaxQuantityActive
// 	if rand.Intn(activation_chance) == 0 {
// 		child.MaxQuantityActive = !g.MaxQuantityActive
// 	}

// 	return
// }
