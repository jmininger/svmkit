package agave

import (
	"fmt"
)

// GeyserPlugin provides native support for YellowstoneGrpc geyser plugins through structured configuration.
// For other geyser plugins, use GenericPluginConfig to specify a JSON config string that must contain
// a top-level "libpath" field pointing to the plugin's shared library based on where it will be installed
// on the host machine.

type GeyserPlugin struct {
	// All these options are mutally exclusive
	YellowstoneGrpcConfig *YellowstoneGrpcConfig `pulumi:"yellowstoneGrpcConfig,optional"`
	GenericPluginConfig   *string                `pulumi:"genericPluginConfig,optional"`
}

func (g *GeyserPlugin) Check() error {
	if g.YellowstoneGrpcConfig != nil && g.GenericPluginConfig != nil {
		return fmt.Errorf("only one of YellowstoneGrpcConfig or GenericPluginConfig can be specified")
	}

	if g.YellowstoneGrpcConfig == nil && g.GenericPluginConfig == nil {
		return fmt.Errorf("either YellowstoneGrpcConfig or GenericPluginConfig must be specified")
	}

	return nil
}

type GeyserPlugin struct {
	YellowstoneGrpcConfig *YellowstoneGrpcConfig `pulumi:"yellowstoneGrpcConfig,optional"`
	GenericPluginConfig   *string                `pulumi:"genericPluginConfig,optional"`
}

func (g *GeyserPlugin) Check() error {
	if g.YellowstoneGrpcConfig != nil && g.GenericPluginConfig != nil {
		return fmt.Errorf("only one of YellowstoneGrpcConfig or GenericPluginConfig can be specified")
	}

	if g.YellowstoneGrpcConfig == nil && g.GenericPluginConfig == nil {
		return fmt.Errorf("either YellowstoneGrpcConfig or GenericPluginConfig must be specified")
	}

	return nil
}

type YellowstoneGrpcConfig struct {
	LibPath          string                `json:"libpath" pulumi:"libpath"`
	Log              GrpcConfigLog         `json:"log" pulumi:"log"`
	Tokio            GrpcConfigTokio       `json:"tokio" pulumi:"tokio"`
	Grpc             GrpcConfigGrpc        `json:"grpc" pulumi:"grpc"`
	Prometheus       *GrpcConfigPrometheus `json:"prometheus,omitempty" pulumi:"prometheus,optional"`
	DebugClientsHttp bool                  `json:"debug_clients_http" pulumi:"debugClientsHttp"`
}

type GrpcConfigLog struct {
	Level string `json:"level" pulumi:"level"`
}

type GrpcConfigTokio struct {
	WorkerThreads *int  `json:"worker_threads,omitempty" pulumi:"workerThreads,optional"`
	Affinity      []int `json:"affinity,omitempty" pulumi:"affinity,optional"`
}

type GrpcConfigGrpc struct {
	Address                       string                    `json:"address" pulumi:"address"`
	TlsConfig                     *GrpcConfigGrpcServerTls  `json:"tls_config,omitempty" pulumi:"tlsConfig,optional"`
	Compression                   GrpcConfigGrpcCompression `json:"compression" pulumi:"compression"`
	MaxDecodingMessageSize        int                       `json:"max_decoding_message_size" pulumi:"maxDecodingMessageSize"`
	SnapshotPluginChannelCapacity *int                      `json:"snapshot_plugin_channel_capacity,omitempty" pulumi:"snapshotPluginChannelCapacity,optional"`
	SnapshotClientChannelCapacity int                       `json:"snapshot_client_channel_capacity" pulumi:"snapshotClientChannelCapacity"`
	ChannelCapacity               int                       `json:"channel_capacity" pulumi:"channelCapacity"`
	UnaryConcurrencyLimit         int                       `json:"unary_concurrency_limit" pulumi:"unaryConcurrencyLimit"`
	UnaryDisabled                 bool                      `json:"unary_disabled" pulumi:"unaryDisabled"`
	FilterLimits                  GrpcConfigFilterLimits    `json:"filter_limits" pulumi:"filterLimits"`
	XToken                        *string                   `json:"x_token,omitempty" pulumi:"xToken,optional"`
	FilterNameSizeLimit           int                       `json:"filter_name_size_limit" pulumi:"filterNameSizeLimit"`
	FilterNamesSizeLimit          int                       `json:"filter_names_size_limit" pulumi:"filterNamesSizeLimit"`
	FilterNamesCleanupInterval    string                    `json:"filter_names_cleanup_interval" pulumi:"filterNamesCleanupInterval"`
	ReplayStoredSlots             uint64                    `json:"replay_stored_slots" pulumi:"replayStoredSlots"`
	ServerHttp2AdaptiveWindow     *bool                     `json:"server_http2_adaptive_window,omitempty" pulumi:"serverHttp2AdaptiveWindow,optional"`
	ServerHttp2KeepaliveInterval  *string                   `json:"server_http2_keepalive_interval,omitempty" pulumi:"serverHttp2KeepaliveInterval,optional"`
	ServerHttp2KeepaliveTimeout   *string                   `json:"server_http2_keepalive_timeout,omitempty" pulumi:"serverHttp2KeepaliveTimeout,optional"`
	ServerInitialConnWindowSize   *uint32                   `json:"server_initial_connection_window_size,omitempty" pulumi:"serverInitialConnectionWindowSize,optional"`
	ServerInitialStreamWindowSize *uint32                   `json:"server_initial_stream_window_size,omitempty" pulumi:"serverInitialStreamWindowSize,optional"`
}

type GrpcConfigGrpcServerTls struct {
	CertPath string `json:"cert_path" pulumi:"certPath"`
	KeyPath  string `json:"key_path" pulumi:"keyPath"`
}

type GrpcConfigGrpcCompression struct {
	Accept []string `json:"accept" pulumi:"accept,optional"`
	Send   []string `json:"send" pulumi:"send,optional"`
}

type GrpcConfigPrometheus struct {
	Address string `json:"address" pulumi:"address"`
}

type GrpcConfigFilterLimits struct {
	Accounts           GrpcConfigFilterLimitsAccounts     `json:"accounts" pulumi:"accounts"`
	Slots              GrpcConfigFilterLimitsSlots        `json:"slots" pulumi:"slots"`
	Transactions       GrpcConfigFilterLimitsTransactions `json:"transactions" pulumi:"transactions"`
	TransactionsStatus GrpcConfigFilterLimitsTransactions `json:"transactions_status" pulumi:"transactionsStatus"`
	Blocks             GrpcConfigFilterLimitsBlocks       `json:"blocks" pulumi:"blocks"`
	BlocksMeta         GrpcConfigFilterLimitsBlocksMeta   `json:"blocks_meta" pulumi:"blocksMeta"`
	Entries            GrpcConfigFilterLimitsEntries      `json:"entries" pulumi:"entries"`
}

type GrpcConfigFilterLimitsAccounts struct {
	Max           int      `json:"max" pulumi:"max"`
	Any           bool     `json:"any" pulumi:"any"`
	AccountMax    int      `json:"account_max" pulumi:"accountMax"`
	AccountReject []string `json:"account_reject" pulumi:"accountReject,optional"`
	OwnerMax      int      `json:"owner_max" pulumi:"ownerMax"`
	OwnerReject   []string `json:"owner_reject" pulumi:"ownerReject,optional"`
	DataSliceMax  int      `json:"data_slice_max" pulumi:"dataSliceMax"`
}

type GrpcConfigFilterLimitsSlots struct {
	Max int `json:"max" pulumi:"max"`
}

type GrpcConfigFilterLimitsTransactions struct {
	Max                  int      `json:"max" pulumi:"max"`
	Any                  bool     `json:"any" pulumi:"any"`
	AccountIncludeMax    int      `json:"account_include_max" pulumi:"accountIncludeMax"`
	AccountIncludeReject []string `json:"account_include_reject" pulumi:"accountIncludeReject,optional"`
	AccountExcludeMax    int      `json:"account_exclude_max" pulumi:"accountExcludeMax"`
	AccountRequiredMax   int      `json:"account_required_max" pulumi:"accountRequiredMax"`
}

type GrpcConfigFilterLimitsBlocks struct {
	Max                  int      `json:"max" pulumi:"max"`
	AccountIncludeMax    int      `json:"account_include_max" pulumi:"accountIncludeMax"`
	AccountIncludeAny    bool     `json:"account_include_any" pulumi:"accountIncludeAny"`
	AccountIncludeReject []string `json:"account_include_reject" pulumi:"accountIncludeReject,optional"`
	IncludeTransactions  bool     `json:"include_transactions" pulumi:"includeTransactions"`
	IncludeAccounts      bool     `json:"include_accounts" pulumi:"includeAccounts"`
	IncludeEntries       bool     `json:"include_entries" pulumi:"includeEntries"`
}

type GrpcConfigFilterLimitsBlocksMeta struct {
	Max int `json:"max" pulumi:"max"`
}

type GrpcConfigFilterLimitsEntries struct {
	Max int `json:"max" pulumi:"max"`
}
