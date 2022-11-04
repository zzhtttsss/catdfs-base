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
	MasterCheckTime             = "chunk.checkTime"
	ChunkWaitingTime            = "chunk.waitingTime"
	ChunkDieTime                = "chunk.dieTime"
	ReplicaNum                  = "chunk.replicaNum"
	ChunkDeadChunkCheckTime     = "chunk.deadChunkCheckTime"
	ChunkDeadChunkCopyThreshold = "chunk.deadChunkCopyThreshold"
	DirectoryCheckTime          = "master.directoryCheckTime"
	CleanupTime                 = "master.cleanupTime"

	// chunk server config
	ChunkHeartbeatReconnectCount = "chunk.heartbeat.reconnectCount"
	ChunkHeartbeatSendTime       = "chunk.heartbeat.sendTime"
	ChunkStoragePath             = "chunk.storage.path"
	ChunkCheckTime               = "chunk.storage.checkTime"
	ChunkDeadTime                = "chunk.storage.chunkDeadTime"

	// etcd config
	EtcdEndPoint = "etcd.endPoint"
)

// Network const value
const (
	TCP     = "tcp"
	LocalIP = "127.0.0.1"
)

// DataNode status
const (
	Cold    = 0
	Alive   = 1
	Waiting = 2
)

// FutureSendChunks status
const (
	WaitToInform = 0
	WaitToSend   = 1
)

// SendType in ChunkSendInfo
const (
	CopySendType   = 0
	MoveSendType   = 1
	AddSendType    = 2
	DeleteSendType = 3
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
	OperationRegister       = "Register"
	OperationHeartbeat      = "Heartbeat"
	OperationAdd            = "Add"
	OperationGet            = "Get"
	OperationRemove         = "Remove"
	OperationMove           = "Move"
	OperationRename         = "Rename"
	OperationList           = "List"
	OperationMkdir          = "Mkdir"
	OperationStat           = "Stat"
	OperationAllocateChunks = "AllocateChunks"
	OperationExpand         = "Expand"
	OperationDegrade        = "Degrade"
	OperationTreeCheck      = "TreeCheck"
	OperationDataCheck      = "DataCheck"
)

// operation stage
const (
	CheckArgs       = 1
	GetDataNodes    = 2
	UnlockDic       = 3
	Degrade2Waiting = 1
	Degrade2Dead    = 2
)

const (
	IsFile4AddFile = true
	IsFile4AddDic  = false
)
