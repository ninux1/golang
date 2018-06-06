package singleton

//Singleton interface
type Singleton interface {
	AddOne()
}

type singleton struct {
	count int
}

var instance *singleton

// GetInstance returns an instance
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() {
	s.count = s.count + 1
}
