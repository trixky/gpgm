package simulation

type Storage map[string]int

func NewStorage(stocks map[string]int) Storage {
	storage := make(Storage)
	for k, v := range stocks {
		storage[k] = v
	}
	return storage
}

func (s Storage) DeepCopy() Storage {
	newStorage := make(Storage)
	for k, v := range s {
		newStorage[k] = v
	}
	return newStorage
}

func (s Storage) ConsumeIfAvailable(inputs map[string]int) bool {
	for k, v := range inputs {
		stored, ok := s[k]
		if !ok || stored - v < 0 {
			return false
		}
	}
	for k, v := range inputs {
		stored, _ := s[k]
		s[k] = stored - v
	}
	return true
}

func (s Storage) Store(outputs map[string]int) {
	for k, v := range outputs {
		stored, ok := s[k]
		if !ok {
			s[k] = v
			continue
		}
		s[k] = stored + v
	}
}

func (s Storage) Get(name string) int {
	v, ok := s[name]
	if !ok {
		return 0
	}
	return v
}