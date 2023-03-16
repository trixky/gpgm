package instance

import (
	"math/rand"

	"github.com/trixky/krpsim/algo/core"
)

type EntryGene struct {
	Process_ids []int
}

// DeepCopy make a deep copy of itself
func (eg *EntryGene) DeepCopy() *EntryGene {
	return &EntryGene{
		Process_ids: append(make([]int, 0, len(eg.Process_ids)), eg.Process_ids...),
	}
}

// Shuffle shuffles its process ids
func (eg *EntryGene) Shuffle() {
	// Initializes the new process ids array
	dest := make([]int, len(eg.Process_ids))
	// Generate a random array of index
	perm := rand.Perm(len(eg.Process_ids))

	for random_index, random := range perm {
		// For each random index
		// Use it to extract a random process id
		dest[random_index] = eg.Process_ids[random]
	}

	eg.Process_ids = dest
}

// RandomCut cut processes randomly when is possible
func (eg *EntryGene) RandomCut(luck int) {
	// WARNING: luck at one cut all processes except the last one

	if luck > 0 {
		// If the luck is stricly positive
		for i := len(eg.Process_ids); i > 1; i-- {
			// For each process except the last one

			if rand.Intn(luck) == 0 {
				// Cut the first process id randomly regarding the luck
				eg.Process_ids = eg.Process_ids[1:]
			}
		}
	}
}

// CutN cut n processes
func (eg *EntryGene) CutN(n uint) {
	if n > 0 && n < uint(len(eg.Process_ids)) {
		// If n is stricly positive
		// Cut the process ids
		eg.Process_ids = eg.Process_ids[:n]
	}
}

// CutN cut a random n processes
func (eg *EntryGene) CutRandomN() {
	if length := len(eg.Process_ids); length > 0 {
		// If at least one process id is present
		// Get the random n
		n := rand.Intn(length) + 1
		// Cut the process ids
		eg.Process_ids = eg.Process_ids[:n]
	}
}

// InitProcesses initalizes the processes
func (eg *EntryGene) InitProcesses(processes []core.Process, optimize map[string]bool) {
	for resource := range optimize { // "time" is not used
		// For each resource to optimize
		for process_index, process := range processes {
			// For each process
			for output := range process.Outputs {
				// For each output of the process
				if output == resource {
					// If the output corresponds to the resource
					if rand.Intn(2) == 0 { // Randomization
						// Add it at the beginning
						eg.Process_ids = append([]int{process_index}, eg.Process_ids...)
					} else {
						// Add it at the end
						eg.Process_ids = append(eg.Process_ids, process_index)
					}
				}
			}
		}
	}
}

// Init initalizes the processes
func (eg *EntryGene) Init(processes []core.Process, optimize map[string]bool, options *core.Options) {
	// Initializes processes
	eg.InitProcesses(processes, optimize)
	eg.Shuffle()

	if options.RandomCut {
		// If random option is active
		// Cut randomly processes
		eg.CutRandomN()
	}

	// Cut processes regarding max
	// Remember is ignored if equal 0
	eg.CutN(uint(options.MaxCut))
}

// ContainProcessid checks if it contain a specific process id
func (eg *EntryGene) ContainProcessid(process_id int) bool {
	for _, eg_process_id := range eg.Process_ids {
		// For each process id
		if eg_process_id == process_id {
			// If the process id correspond to the searched one
			return true
		}
	}

	return false
}

// InsertProcessIdRandomPosition Inserts a process id at a random position
func (eg *EntryGene) InsertProcessIdRandomPosition(process_id int) {
	if process_ids_length := len(eg.Process_ids); process_ids_length > 0 {
		// If it contains at least one process id

		// Copy the process ids
		process_ids_copy := make([]int, process_ids_length)
		copy(process_ids_copy, eg.Process_ids)

		// Get a random position

		random_position := rand.Intn(process_ids_length)

		// Add the new process id to the first part of the process ids
		// cutted by the random position
		new_process_ids := append(process_ids_copy[:random_position], process_id)
		// Add the last part of the process ids
		eg.Process_ids = append(new_process_ids, eg.Process_ids[random_position:]...)
	} else {
		// Else it's empty
		eg.Process_ids = []int{process_id}
	}
}

// RemoveProcessIdRandomPosition Removes a process id at a random position
func (eg *EntryGene) RemoveProcessIdRandomPosition() {
	if process_ids_length := len(eg.Process_ids); process_ids_length > 0 {
		// If it contains at least one process id

		// Get a random position
		random_position := 0
		if process_ids_length > 1 {
			random_position = rand.Intn(process_ids_length)
		}

		// Add the new process id to the first part of the process ids
		// cutted by the random position
		eg.Process_ids = append(eg.Process_ids[:random_position], eg.Process_ids[random_position+1:]...)
	}
}

// Mutate mutates according to a pourcentage
func (eg *EntryGene) Mutate(egeg *EntryGene, options *core.Options) *EntryGene {
	mutated_entry_gene := EntryGene{}

	// Use the pourcentage of chance
	if rand.Intn(1000) < int(options.MutationChance*1000) {
		// Mutation

		random := rand.Intn(9)
		if random == 0 { // 1/9 chance
			// Mutate fully as the new entry gene
			mutated_entry_gene = *egeg.DeepCopy()
		} else { // 8/9 chance
			// Mutate randomly as the new entry gene

			// ------------------------ New process ids
			new_process_ids := []int{}

			for _, egeg_process_id := range egeg.Process_ids {
				if !eg.ContainProcessid(egeg_process_id) {
					new_process_ids = append(new_process_ids, egeg_process_id)
				}
			}

			// ------------------------ Mutation
			mutated_entry_gene = *eg.DeepCopy()

			if diff := len(new_process_ids); diff > 0 {
				// Add randomly new process ids
				for _, new_process_id := range new_process_ids {
					mutated_entry_gene.InsertProcessIdRandomPosition(new_process_id)
				}

				// Remove randomly new/inital process ids
				for i := rand.Intn(len(new_process_ids)); i >= 0; i-- {
					mutated_entry_gene.RemoveProcessIdRandomPosition()
				}
			}
		}
	} else {
		// No mutation
		mutated_entry_gene = *eg.DeepCopy()
	}

	return &mutated_entry_gene
}
