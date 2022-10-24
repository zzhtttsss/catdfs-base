package common

// Config key string
const (
	// common config
	MasterPort         = "master.rpcPort"
	MasterRaftPort     = "master.raftPort"
	MasterRaftDir      = "master.raftDir"
	ChunkPort          = "chunk.port"
	ClientPort         = "client.port"
	ChunkHeartbeatTime = "chunk.heartbeatTime"

	// master config
	MasterCheckTime  = "chunk.checkTime"
	ChunkWaitingTime = "chunk.waitingTime"
	ChunkDieTime     = "chunk.dieTime"
	ReplicaNum       = "chunk.replicaNum"

	// chunk server config
	ChunkHeartbeatReconnectCount = "chunk.heartbeat.reconnectCount"
	ChunkHeartbeatSendTime       = "chunk.heartbeat.sendTime"
	ChunkStoragePath             = "chunk.storage.path"
	ChunkDeadChunkCheckTime      = "chunk.deadChunkCheckTime"
	ChunkDeadChunkCopyThreshold  = "chunk.deadChunkCopyThreshold"

	// etcd config
	EtcdEndPoint = "etcd.endPoint"
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
	LogDBName         = "logs.dat"
	StableDBName      = "stable.dat"
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
	MasterRPCServerFailed = iota + 3001
	MasterRegisterFailed
	MasterHeartbeatFailed
	MasterCheckArgs4AddFailed
	MasterGetDataNodes4AddFailed
	MasterUnlockDic4AddFailed
	MasterSendOperationFailed
	MasterFinishOperationFailed
	MasterCheckAndMkdirFailed
	MasterCheckAndMoveFailed
	MasterCheckAndRemoveFailed
	MasterCheckAndListFailed
	MasterCheckAndStatFailed
	MasterCheckAndRenameFailed
	MasterCheckAndGetFailed
	MasterGetDataNodes4GetFailed
	MasterShrinkFailed
	MasterExpandFailed
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
	AddressDelimiter  = ":"
	ChunkIdDelimiter  = "_"
	MinusOneString    = "-1"
	DollarDelimiter   = "$"
	SnapshotDelimiter = "$$$$$$\n"
)

// map keys
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

// etcd keys
const (
	TinyDFSPrefix     = "/tinyDFS/"
	FollowerKeyPrefix = TinyDFSPrefix + "follower/"
	LeaderAddressKey  = TinyDFSPrefix + "leader-address"
)

// operation type
const (
	OperationRegister   = "Register"
	OperationHeartbeat  = "Heartbeat"
	OperationAdd        = "Add"
	OperationGet        = "Get"
	OperationRemove     = "Remove"
	OperationMove       = "Move"
	OperationRename     = "Rename"
	OperationList       = "List"
	OperationMkdir      = "Mkdir"
	OperationStat       = "Stat"
	OperationShrink     = "Shrink"
	OperationExpand     = "Expand"
	OperationDeregister = "Deregister"
)

// operation stage
const (
	CheckArgs    = 1
	GetDataNodes = 2
	UnlockDic    = 3
)

const (
	IsFile4AddFile = true
	IsFile4AddDic  = false
)
