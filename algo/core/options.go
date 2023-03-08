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
	MaxGeneration    int `json:"max_generation"`
	TimeLimitSeconds int `json:"time_limit_seconds"`
	MaxCycle         int `json:"max_cycle"`
	PopulationSize   int `json:"population_size"`
	// MutationMethod   MutationMethod
	// SelectionMethod  SelectionMethod
}
