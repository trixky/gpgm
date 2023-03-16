package core

type Stock map[string]int

func NewStorage(stocks map[string]int) Stock {
	storage := make(Stock)
	for k, v := range stocks {
		storage[k] = v
	}
	return storage
}

func (s Stock) DeepCopy() Stock {
	newStorage := make(Stock)
	for k, v := range s {
		newStorage[k] = v
	}
	return newStorage
}

// RemoveResource removes a quantity of resource
func (s Stock) RemoveResource(resource string, quantity int) {
	// WARNING: no negative checker
	s[resource] = s[resource] - quantity
}

// RemoveResource set a quantity of resource
func (s Stock) SetResource(resource string, quantity int) {
	s[resource] = quantity
}

// AddResource add a quantity of resource
func (s Stock) AddResource(resource string, quantity int) {
	// WARNING: no overflow checker
	s[resource] = s[resource] + quantity
}

// HaveResource checks if it has a resource
func (s Stock) HaveResource(resource string) bool {
	_, exists := s[resource]

	return exists
}

// GetResource get a the quantity of a resource
func (s Stock) GetResource(resource string) int {
	if quantity, ok := s[resource]; !ok {
		return 0
	} else {
		return quantity
	}
}
