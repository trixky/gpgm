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
	MaxGeneration    int
	TimeLimitSeconds int
	MaxCycle         int
	PopulationSize   int
	// MutationMethod   MutationMethod
	// SelectionMethod  SelectionMethod
}
