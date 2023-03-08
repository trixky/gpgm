package simulation

type ExpectedStock struct {
	Name            string `json:"name"`
	Quantity        int    `json:"quantity"`
	RemainingCycles int    `json:"remaining_cycles"`
}
