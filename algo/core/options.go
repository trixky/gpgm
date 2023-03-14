package core

type SelectionMethod int64

const (
	RandomSelection SelectionMethod = iota
	TournamentSelection
)

type MutationMethod int64

const (
	LinearMutation MutationMethod = iota
	LogarithmicMutation
)

type Options struct {
	MaxGeneration    int `json:"max_generation"`
	TimeLimitSeconds int `json:"time_limit_seconds"`
	MaxCycle         int `json:"max_cycle"`
	MaxDepth         int `json:"max_depth"`
	// Population
	PopulationSize        int             `json:"population_size"`
	ElitismAmount         int             `json:"elitism_amount"`
	SelectionMethod       SelectionMethod `json:"selection_method"`
	TournamentSize        int             `json:"tournament_size"`
	TournamentProbability float64         `json:"tournament_probability"`
	CrossoverNewInstances int             `json:"crossover_new_instances"`
	// Mutation
	MutationChance float64        `json:"mutation_chance"`
	MutationMethod MutationMethod `json:"mutation_method"`
	// Genetic
	NEntry               int  `json:"n_entry"`
	HistoryPartMaxLength int  `json:"history_part_max_length"`
	HistoryKeyMaxLength  int  `json:"history_key_max_length"`
	RandomCut            bool `json:"random_cut"`
	MaxCut               int  `json:"max_cut"`
}
