package common

// Config key string
const (
	// common config
	MasterAddr = "master.addr"
	MasterPort = "master.port"

	// master config
	ChunkWaitTime = "chunk.waitTime"
	ChunkDieTime  = "chunk.dieTime"
	ReplicaNum    = "replicaNum"

	// chunk server config
	ChunkHeartbeatReconnectCount = "chunk.heartbeat.reconnectCount"
	ChunkHeartbeatSendTime       = "chunk.heartbeat.sendTime"
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
	MasterRPCServerFailed        = 3001
	MasterRegisterFailed         = 3002
	MasterHeartbeatFailed        = 3003
	MasterCheckArgs4AddFailed    = 3004
	MasterGetDataNodes4AddFailed = 3005
)

// chunkserver error code
const (
	ChunkServerRPCServerFailed = 4001
	ChunkServerRegisterFailed  = 4002
	ChunkServerHeartbeatFailed = 4003
)
