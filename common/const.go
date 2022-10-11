package common

// Config key string
const (
	// common config
	MasterPort     = "master.rpcPort"
	MasterRaftPort = "master.raftPort"
	MasterRaftDir  = "master.raftDir"
	ChunkPort      = "chunk.port"
	ClientPort     = "client.port"

	// master config
	MasterCheckTime = "chunk.checkTime"
	ChunkDieTime    = "chunk.dieTime"
	ReplicaNum      = "chunk.replicaNum"

	// chunk server config
	ChunkHeartbeatReconnectCount = "chunk.heartbeat.reconnectCount"
	ChunkHeartbeatSendTime       = "chunk.heartbeat.sendTime"
	ChunkStoragePath             = "chunk.storage.path"

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
	MasterRPCServerFailed        = 3001
	MasterRegisterFailed         = 3002
	MasterHeartbeatFailed        = 3003
	MasterCheckArgs4AddFailed    = 3004
	MasterGetDataNodes4AddFailed = 3005
	MasterUnlockDic4AddFailed    = 3006
	MasterReleaseLease4AddFailed = 3007
	MasterSendOperationFailed    = 3008
	MasterFinishOperationFailed  = 3009
	MasterCheckAndMkdirFailed    = 3010
	MasterCheckAndMoveFailed     = 3011
	MasterCheckAndRemoveFailed   = 3012
	MasterCheckAndListFailed     = 3013
	MasterCheckAndStatFailed     = 3014
	MasterCheckAndRenameFailed   = 3015
	MasterCheckAndGetFailed      = 3016
	MasterGetDataNodes4GetFailed = 3017
	MasterReleaseLease4GetFailed = 3018
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
	OperationRegister  = "Register"
	OperationHeartbeat = "Heartbeat"
	OperationAdd       = "Add"
	OperationGet       = "Get"
	OperationRemove    = "Remove"
	OperationMove      = "Move"
	OperationRename    = "Rename"
	OperationList      = "List"
	OperationMkdir     = "Mkdir"
	OperationStat      = "Stat"
	OperationFinish    = "Finish"
)

// operation stage
const (
	CheckArgs    = 1
	GetDataNodes = 2
	ReleaseLease = 3
	UnlockDic    = 4
)

const (
	IsFile4AddFile = true
	IsFile4AddDic  = false
)
