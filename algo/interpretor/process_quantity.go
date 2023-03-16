package interpretor

import "github.com/trixky/krpsim/algo/core"

type ProcessQuantity struct {
	Process  *core.Process `json:"process"`
	Quantity int           `json:"quantity"`
}

type ProcessQuantities struct {
	Stack []ProcessQuantity `json:"process_quantity"`
}

// DeepCopy creates a deep copy of itself
func (pqs *ProcessQuantities) DeepCopy() (copy *ProcessQuantities) {
	// Initialize the copy
	copy = &ProcessQuantities{
		Stack: make([]ProcessQuantity, len(pqs.Stack)),
	}

	for process_quantity_index, process_quantity := range pqs.Stack {
		// For each initial process quantity
		// Add an copy to the process quantities copy
		copy.Stack[process_quantity_index] = ProcessQuantity{
			Process:  process_quantity.Process,
			Quantity: process_quantity.Quantity,
		}
	}

	return
}

// Concatenate concatenates its stack with another one
func (pqs *ProcessQuantities) Concatenate(second_pqs ProcessQuantities) {
	if len(second_pqs.Stack) > 0 {
		// If the given process quantities is not empty
		pqs.Stack = append(pqs.Stack, second_pqs.Stack...)
	}
}

// Push push a process quantity in its stack
func (pqs *ProcessQuantities) Push(process_quantities *ProcessQuantity) {
	pqs.Stack = append(pqs.Stack, *process_quantities)
}
