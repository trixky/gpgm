package core

// type SelectionMethod int64

// const (
// 	Random SelectionMethod = iota
// )

// type MutationMethod int64

// const (
// 	Linear MutationMethod = iota
// 	Logarithmic
// )

type Options struct {
	MaxGeneration        int  `json:"max_generation"`
	TimeLimitSeconds     int  `json:"time_limit_seconds"`
	MaxCycle             int  `json:"max_cycle"`
	MaxDepth             int  `json:"max_depth"`
	NEntry               int  `json:"n_entry"`
	HistoryPartMaxLength int  `json:"history_part_max_length"`
	HistoryKeyMaxLength  int  `json:"history_key_max_length"`
	PopulationSize       int  `json:"population_size"`
	UseElitism           bool `json:"use_elitism"`
	ElitismAmount        int  `json:"elitism_amount"`
	RandomCut            bool `json:"random_cut"`
	MaxCut               int  `json:"max_cut"`
	// MutationMethod   MutationMethod
	// SelectionMethod  SelectionMethod
}
