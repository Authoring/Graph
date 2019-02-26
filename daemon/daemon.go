package daemon

// Daemon defines the daemon
type Daemon struct {
	Port int
}

// InitDaemon init a new daemon object
func InitDaemon() *Daemon {
	return &Daemon{
		Port: 7475,
	}
}
