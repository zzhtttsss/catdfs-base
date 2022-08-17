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

// common error code
const (
	Connect2ClientFailed      = 1001
	Connect2MasterFailed      = 1002
	Connect2ChunkServerFailed = 1003
)

// client error code
const (
	ClientRPCServerFailed = 2001
)

// master error code
const (
	MasterRPCServerFailed = 3001
	MasterRegisterFailed  = 3002
)

// chunkserver error code
const (
	ChunkServerRPCServerFailed = 4001
)
