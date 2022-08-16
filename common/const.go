package common

// Config key string
const (
	MasterAddr = "master.addr"
	MasterPort = "master.port"
)

// Network const value
const (
	TCP     = "tcp"
	LocalIP = "127.0.0.1"
)

// Datanode status
const (
	Died    = 0
	Alive   = 1
	Waiting = 2
)
