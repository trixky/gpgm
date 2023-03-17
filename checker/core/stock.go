package core

type Stock map[string]int

func NewStorage(stocks map[string]int) Stock {
	storage := make(Stock)
	for k, v := range stocks {
		storage[k] = v
	}
	return storage
}

// Reset removes all resources
func (s *Stock) Reset() {
	for resource := range *s {
		delete(*s, resource)
	}
}

// DeepCopy make a deep copy
func (s Stock) DeepCopy() *Stock {
	new_stock := make(Stock)

	for resource, quantity := range s {
		//
		new_stock[resource] = quantity
	}

	return &new_stock
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

// ResourceExists checks if a resource exists without looking at it's quantity
func (s Stock) ResourceExists(resource string) bool {
	_, exists := s[resource]

	return exists
}

// HaveResource checks if it has a resource
func (s Stock) HaveResource(resource string) bool {
	quantity, ok := s[resource]

	return ok && quantity > 0
}

// GetResource get a the quantity of a resource
func (s Stock) GetResource(resource string) int {
	if quantity, ok := s[resource]; !ok {
		return 0
	} else {
		return quantity
	}
}
