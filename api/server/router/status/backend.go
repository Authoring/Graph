package status

// Backend defines the status backend
type Backend interface {
	Status() (bool, error)
}
