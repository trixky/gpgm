package core

type Product struct {
	Name     string
	Quantity int
}

type Process struct {
	Name    string
	Inputs  map[string]int
	Outputs map[string]int
	Delay   int
}
