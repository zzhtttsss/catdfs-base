package common

// Config key string
const (
	// common config
	MasterAddr = "master.addr"
	MasterPort = "master.port"
	ChunkPort  = "chunk.port"
	ClientPort = "client.port"

	// master config
	ChunkWaitTime      = "chunk.waitTime"
	ChunkDieTime       = "chunk.dieTime"
	ReplicaNum         = "chunk.replicaNum"
	SMAddr             = "shadowMaster.addr"
	SMPort             = "shadowMaster.port"
	SMFsimageFlushTime = "shadowMaster.fsimageFlushTime"

	// chunk server config
	ChunkHeartbeatReconnectCount = "chunk.heartbeat.reconnectCount"
	ChunkHeartbeatSendTime       = "chunk.heartbeat.sendTime"
	ChunkStoragePath             = "chunk.storage.path"
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

// Persistence String
const (
	LogFileName       = "log/edits.txt"
	DirectoryFileName = "log/fsimage.txt"
	NormalTimeFormat  = "2006-01-02 15:04:05"
	LogFileTimeFormat = "2006-01-02.15.04.05"
)

// common error code
const (
	Connect2ClientFailed      = 1001
	Connect2MasterFailed      = 1002
	Connect2ChunkServerFailed = 1003
	JsonMarshalFailed         = 1099
	JsonUnmarshalFailed       = 1100
)

// client error code
const (
	ClientRPCServerFailed = 2001
)

// master error code
const (
	MasterRPCServerFailed             = 3001
	MasterRegisterFailed              = 3002
	MasterHeartbeatFailed             = 3003
	MasterCheckArgs4AddFailed         = 3004
	MasterGetDataNodes4AddFailed      = 3005
	MasterUnlockDic4AddFailed         = 3006
	MasterReleaseLease4AddFailed      = 3007
	MasterSendOperationFailed         = 3008
	MasterFinishOperationFailed       = 3009
	MasterCheckAndMkdirFailed         = 3010
	MasterCheckAndMoveFailed          = 3011
	MasterCheckAndRemoveFailed        = 3012
	MasterCheckAndListFailed          = 3013
	MasterCheckAndStatFailed          = 3014
	MasterCheckAndRenameFailed        = 3015
	MasterCheckAndGetFailed           = 3016
	MasterGetDataNodes4GetFailed      = 3017
	MasterReleaseLease4GetFailed      = 3018
	ShadowMasterRPCServerFailed       = 3049
	ShadowMasterSendOperationFailed   = 3050
	ShadowMasterFinishOperationFailed = 3051
)

// chunkserver error code
const (
	ChunkServerRPCServerFailed     = 4001
	ChunkServerRegisterFailed      = 4002
	ChunkServerHeartbeatFailed     = 4003
	ChunkServerTransferChunkFailed = 4004
)

// delimiter
const (
	AddressDelimiter = ":"
	ChunkIdDelimiter = "_"
	MinusOneString   = "-1"
)

// map key
const (
	ChunkIdString    = "chunkId"
	AddressString    = "address"
	ChunkIndexString = "ChunkIndex"
)

// const number
const (
	KB         = 1024
	MB         = 1024 * KB
	GB         = 1024 * MB
	ChunkMBNum = 64
	ChunkSize  = ChunkMBNum * MB
	DirSize    = 0
)
