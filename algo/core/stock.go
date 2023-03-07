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

func (s Stock) Remove(name string, quantity int) {
	s[name] = s[name] - quantity
}

func (s Stock) Insert(name string, quantity int) {
	s[name] = quantity
}

func (s Stock) Add(name string, quantity int) {
	s[name] = s[name] + quantity
}

func (s Stock) Exists(name string) bool {
	_, exists := s[name]
	return exists
}

func (s Stock) Get(name string) int {
	v, ok := s[name]
	if !ok {
		return 0
	}
	return v
}
