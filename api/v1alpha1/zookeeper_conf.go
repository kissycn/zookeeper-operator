package v1alpha1

type LogLevels string

const (
	ERROR LogLevels = "ERROR"
	WARN  LogLevels = "WARN"
	INFO  LogLevels = "INFO"
	DEBUG LogLevels = "DEBUG"
	TRACE LogLevels = "TRACE"
)

// zookeeper conf params detail reference this:
// https://github.com/apache/zookeeper/blob/master/zookeeper-docs/src/main/resources/markdown/zookeeperAdmin.md
type ZookeeperConf struct {
	// DataDir Dedicated data log directory
	//
	// +kubebuilder:default:="/opt/dtweave/zookeeper/data"
	DataDir string `json:"dataDir,omitempty"`
	// DataLogDir Dedicated data log directory
	//
	// +kubebuilder:default:="/opt/dtweave/zookeeper/data_log"
	DataLogDir string `json:"dataLogDir,omitempty"`

	// TickTime is the length of a single tick, which is the basic time unit used
	// by Zookeeper, as measured in milliseconds
	//
	// +kubebuilder:default:=2000
	TickTime int32 `json:"tickTime,omitempty"`
	// InitLimit is the amount of time, in ticks, to allow followers to connect
	// and sync to a leader.
	//
	// +kubebuilder:default:=10
	InitLimit int32 `json:"initLimit,omitempty"`
	// SyncLimit is the amount of time, in ticks, to allow followers to sync with
	// Zookeeper.
	//
	// +kubebuilder:default:=2
	SyncLimit int32 `json:"syncLimit,omitempty"`

	// MaxClientCnxns Limits the number of concurrent connections that a single client, identified
	// by IP address, may make to a single member of the ZooKeeper ensemble.
	//
	// The default value is 60
	// +kubebuilder:default:=60
	MaxClientCnxns int32 `json:"maxClientCnxns,omitempty"`
	// The maximum session timeout in milliseconds that the server will allow the
	// client to negotiate.
	//
	// The default value is 40000
	// +kubebuilder:default:=40000
	MaxSessionTimeout int64 `json:"maxSessionTimeout,omitempty"`
	// Autopurge Automatic purging of the snapshots and corresponding transaction logs was introduced in version 3.4.0
	// and can be enabled via the following configuration parameters autopurge.snapRetainCount and autopurge.purgeInterva
	//
	Autopurge Autopurge `json:"autopurge,omitempty"`

	// LogLevel  Log level for the ZooKeeper server. ERROR by default
	//
	// +kubebuilder:validation:Enum=TRACE;DEBUG;INFO;WARN;ERROR
	// +kubebuilder:default:=ERROR
	LogLevel LogLevels `json:"logLevel,omitempty"`
	// JvmFlags  Default JVM flags for the ZooKeeper process
	//
	// +kubebuilder:default:=""
	JvmFlags string `json:"jvmFlags,omitempty"`

	// Auth is zookeeper auth
	Auth ZookeeperAuth `json:"auth,omitempty"`

	// AdditionalConfig key-value map of additional zookeeper configuration parameters
	// +optional
	AdditionalConfig map[string]string `json:"additionalConfig,omitempty"`

	// ExistingCfgConfigmap
	//
	// The name of an existing ConfigMap with your custom configuration for ZooKeeper zoo.cfg file
	// Notice：If set this value operator will replace all other conf item
	ExistingCfgConfigmap string `json:"existingCfgConfigmap,omitempty"`

	// extraEnvVars
	// extraEnvVarsCM
	// extraEnvVarsSecret
	// command
	// args
}

type Autopurge struct {
	// SnapRetainCount The most recent snapshots amount (and corresponding transaction logs) to retain
	//
	// +kubebuilder:default:=3
	SnapRetainCount int32 `json:"snapRetainCount,omitempty"`
	// PurgeInterval The time interval (in hours) for which the purge task has to be triggered
	//
	// +kubebuilder:default:=0
	PurgeInterval int32 `json:"purgeInterval,omitempty"`
}

type ZookeeperAuth struct {
	Client ClientAuth `json:"client,omitempty"`
	Quorum QuorumAuth `json:"quorum,omitempty"`
}

type ClientAuth struct {
	// Enabled ZooKeeper client-server authentication. It uses SASL/Digest-MD5
	//
	// +kubebuilder:default:=false
	Enabled bool `json:"enabled,omitempty"`
	// ClientUser User that will use ZooKeeper clients to auth
	//
	ClientUser string `json:"clientUser,omitempty"`
	// ClientPassword Password that will use ZooKeeper clients to auth
	//
	ClientPassword string `json:"clientPassword,omitempty"`
	// ServerUsers Comma, semicolon or whitespace separated list of user to be created
	//
	ServerUsers string `json:"serverUsers,omitempty"`
	// ServerPasswords Comma, semicolon or whitespace separated list of passwords to assign to users when created
	//
	ServerPasswords string `json:"serverPasswords,omitempty"`
	// ExistingSecret Use existing secret (ignores previous passwords)
	//
	ExistingSecret string `json:"existingSecret,omitempty"`
}

type QuorumAuth struct {
	// Enabled ZooKeeper client-server authentication. It uses SASL/Digest-MD5
	//
	// +kubebuilder:default:=false
	Enabled bool `json:"enabled,omitempty"`
	// ClientUser User that will use ZooKeeper clients to auth
	//
	LearnerUser string `json:"learnerUser,omitempty"`
	// ClientPassword Password that will use ZooKeeper clients to auth
	//
	LearnerPassword string `json:"learnerPassword,omitempty"`
	// ServerUsers Comma, semicolon or whitespace separated list of user to be created
	//
	ServerUsers string `json:"serverUsers,omitempty"`
	// ServerPasswords Comma, semicolon or whitespace separated list of passwords to assign to users when created
	//
	ServerPasswords string `json:"serverPasswords,omitempty"`
	// ExistingSecret Use existing secret (ignores previous passwords)
	//
	ExistingSecret string `json:"existingSecret,omitempty"`
}
