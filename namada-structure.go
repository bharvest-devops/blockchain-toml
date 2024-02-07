package blockchain_toml

type NamadaConfigFile struct {
	WasmDir *string       `json:"wasmDir" toml:"wasm_dir"`
	Ledger  *NamadaLedger `json:"ledger" toml:"ledger"`
}

type NamadaLedger struct {
	GenesisTime    *string               `json:"genesisTime" toml:"genesis_time"`
	ChainID        *string               `json:"chainID" toml:"chain_id"`
	Shell          *NamadaShell          `json:"shell" toml:"shell"`
	Cometbft       *NamadaCometbft       `json:"cometbft" toml:"cometbft"`
	EthereumBridge *NamadaEthereumBridge `json:"ethereumBridge" toml:"ethereum_bridge"`
}

type NamadaShell struct {
	BaseDir                    *string `json:"baseDir" toml:"base_dir"`
	StorageReadPastHeightLimit *uint64 `json:"storageReadPastHeightLimit" toml:"storage_read_past_height_limit"`
	DbDir                      *string `json:"dbDir" toml:"db_dir"`
	CometbftDir                *string `json:"cometbftDir" toml:"cometbft_dir"`
	TendermintMode             *string `json:"tendermintMode" toml:"tendermint_mode"`
}

type NamadaRPC struct {
	Laddr                     *string   `json:"laddr" toml:"laddr"`
	CorsAllowedOrigins        *[]string `json:"corsAllowedOrigins" toml:"cors_allowed_origins"`
	CorsAllowedMethods        *[]string `json:"corsAllowedMethods" toml:"cors_allowed_methods"`
	CorsAllowedHeaders        *[]string `json:"corsAllowedHeaders" toml:"cors_allowed_headers"`
	GrpcLaddr                 *string   `json:"grpcLaddr" toml:"grpc_laddr"`
	GrpcMaxOpenConnections    *int32    `json:"grpcMaxOpenConnections" toml:"grpc_max_open_connections"`
	Unsafe                    *bool     `json:"unsafe" toml:"unsafe"`
	MaxOpenConnections        *int32    `json:"maxOpenConnections" toml:"max_open_connections"`
	MaxSubscriptionClients    *int32    `json:"maxSubscriptionClients" toml:"max_subscription_clients"`
	MaxSubscriptionsPerClient *int32    `json:"maxSubscriptionsPerClient" toml:"max_subscriptions_per_client"`
	TimeoutBroadcastTxCommit  *string   `json:"timeoutBroadcastTxCommit" toml:"timeout_broadcast_tx_commit"`
	MaxBodyBytes              *int      `json:"maxBodyBytes" toml:"max_body_bytes"`
	MaxHeaderBytes            *int      `json:"maxHeaderBytes" toml:"max_header_bytes"`
	TLSCertFile               *string   `json:"tlsCertFile" toml:"tls_cert_file"`
	TLSKeyFile                *string   `json:"tlsKeyFile" toml:"tls_key_file"`
	PprofLaddr                *string   `json:"pprofLaddr" toml:"pprof_laddr"`
}

type NamadaP2P struct {
	Laddr                        *string `json:"laddr" toml:"laddr"`
	ExternalAddress              *string `json:"externalAddress" toml:"external_address"`
	Seeds                        *string `json:"seeds" toml:"seeds"`
	PersistentPeers              *string `json:"persistentPeers" toml:"persistent_peers"`
	Upnp                         *bool   `json:"upnp" toml:"upnp"`
	AddrBookFile                 *string `json:"addrBookFile" toml:"addr_book_file"`
	AddrBookStrict               *bool   `json:"addrBookStrict" toml:"addr_book_strict"`
	MaxNumInboundPeers           *int32  `json:"maxNumInboundPeers" toml:"max_num_inbound_peers"`
	MaxNumOutboundPeers          *int32  `json:"maxNumOutboundPeers" toml:"max_num_outbound_peers"`
	UnconditionalPeerIds         *string `json:"unconditionalPeerIds" toml:"unconditional_peer_ids"`
	PersistentPeersMaxDialPeriod *string `json:"persistentPeersMaxDialPeriod" toml:"persistent_peers_max_dial_period"`
	FlushThrottleTimeout         *string `json:"flushThrottleTimeout" toml:"flush_throttle_timeout"`
	MaxPacketMsgPayloadSize      *int    `json:"maxPacketMsgPayloadSize" toml:"max_packet_msg_payload_size"`
	SendRate                     *int    `json:"sendRate" toml:"send_rate"`
	RecvRate                     *int    `json:"recvRate" toml:"recv_rate"`
	Pex                          *bool   `json:"pex" toml:"pex"`
	SeedMode                     *bool   `json:"seedMode" toml:"seed_mode"`
	PrivatePeerIds               *string `json:"privatePeerIds" toml:"private_peer_ids"`
	AllowDuplicateIP             *bool   `json:"allowDuplicateIP" toml:"allow_duplicate_ip"`
	HandshakeTimeout             *string `json:"handshakeTimeout" toml:"handshake_timeout"`
	DialTimeout                  *string `json:"dialTimeout" toml:"dial_timeout"`
}

type NamadaMempool struct {
	Recheck               *bool   `json:"recheck" toml:"recheck"`
	Broadcast             *bool   `json:"broadcast" toml:"broadcast"`
	WalDir                *string `json:"walDir" toml:"wal_dir"`
	Size                  *int    `json:"size" toml:"size"`
	MaxTxsBytes           *int    `json:"maxTxsBytes" toml:"max_txs_bytes"`
	CacheSize             *int    `json:"cacheSize" toml:"cache_size"`
	KeepInvalidTxsInCache *bool   `json:"keepInvalidTxsInCache" toml:"keep-invalid-txs-in-cache"`
	MaxTxBytes            *int    `json:"maxTxBytes" toml:"max_tx_bytes"`
	MaxBatchBytes         *int    `json:"maxBatchBytes" toml:"max_batch_bytes"`
}

type NamadaConsensus struct {
	WalFile                     *string `json:"walFile" toml:"wal_file"`
	TimeoutPropose              *string `json:"timeoutPropose" toml:"timeout_propose"`
	TimeoutProposeDelta         *string `json:"timeoutProposeDelta" toml:"timeout_propose_delta"`
	TimeoutPrevote              *string `json:"timeoutPrevote" toml:"timeout_prevote"`
	TimeoutPrevoteDelta         *string `json:"timeoutPrevoteDelta" toml:"timeout_prevote_delta"`
	TimeoutPrecommit            *string `json:"timeoutPrecommit" toml:"timeout_precommit"`
	TimeoutPrecommitDelta       *string `json:"timeoutPrecommitDelta" toml:"timeout_precommit_delta"`
	TimeoutCommit               *string `json:"timeoutCommit" toml:"timeout_commit"`
	DoubleSignCheckHeight       *uint64 `json:"doubleSignCheckHeight" toml:"double_sign_check_height"`
	SkipTimeoutCommit           *bool   `json:"skipTimeoutCommit" toml:"skip_timeout_commit"`
	CreateEmptyBlocks           *bool   `json:"createEmptyBlocks" toml:"create_empty_blocks"`
	CreateEmptyBlocksInterval   *string `json:"createEmptyBlocksInterval" toml:"create_empty_blocks_interval"`
	PeerGossipSleepDuration     *string `json:"peerGossipSleepDuration" toml:"peer_gossip_sleep_duration"`
	PeerQueryMaj23SleepDuration *string `json:"peerQueryMaj23SleepDuration" toml:"peer_query_maj23_sleep_duration"`
}

type NamadaStorage struct {
	DiscardAbciResponses *bool `json:"discardAbciResponses" toml:"discard_abci_responses"`
}

type NamadaTxIndex struct {
	Indexer *string `json:"indexer" toml:"indexer"`
}

type NamadaInstrumentation struct {
	Prometheus           *bool   `json:"prometheus" toml:"prometheus"`
	PrometheusListenAddr *string `json:"prometheusListenAddr" toml:"prometheus_listen_addr"`
	MaxOpenConnections   *int    `json:"maxOpenConnections" toml:"max_open_connections"`
	Namespace            *string `json:"namespace" toml:"namespace"`
}

type NamadaStatesync struct {
	Enable        *bool   `json:"enable" toml:"enable"`
	RPCServers    *string `json:"rpcServers" toml:"rpc_servers"`
	TrustHeight   *uint64 `json:"trustHeight" toml:"trust_height"`
	TrustHash     *string `json:"trustHash" toml:"trust_hash"`
	TrustPeriod   *string `json:"trustPeriod" toml:"trust_period"`
	DiscoveryTime *string `json:"discoveryTime" toml:"discovery_time"`
	TempDir       *string `json:"tempDir" toml:"temp_dir"`
}

type NamadaFastsync struct {
	Version *string `json:"version" toml:"version"`
}
type NamadaCometbft struct {
	ProxyApp               *string                `json:"proxyApp" toml:"proxy_app"`
	Moniker                *string                `json:"moniker" toml:"moniker"`
	FastSync               *bool                  `json:"fastSync" toml:"fast_sync"`
	DbBackend              *string                `json:"dbBackend" toml:"db_backend"`
	DbDir                  *string                `json:"dbDir" toml:"db_dir"`
	LogLevel               *string                `json:"logLevel" toml:"log_level"`
	LogFormat              *string                `json:"logFormat" toml:"log_format"`
	GenesisFile            *string                `json:"genesisFile" toml:"genesis_file"`
	PrivValidatorKeyFile   *string                `json:"privValidatorKeyFile" toml:"priv_validator_key_file"`
	PrivValidatorStateFile *string                `json:"privValidatorStateFile" toml:"priv_validator_state_file"`
	PrivValidatorLaddr     *string                `json:"privValidatorLaddr" toml:"priv_validator_laddr"`
	NodeKeyFile            *string                `json:"nodeKeyFile" toml:"node_key_file"`
	Abci                   *string                `json:"abci" toml:"abci"`
	FilterPeers            *bool                  `json:"filterPeers" toml:"filter_peers"`
	RPC                    *NamadaRPC             `json:"rpc" toml:"rpc"`
	P2P                    *NamadaP2P             `json:"p2p" toml:"p2p"`
	Mempool                *NamadaMempool         `json:"mempool" toml:"mempool"`
	Consensus              *NamadaConsensus       `json:"consensus" toml:"consensus"`
	Storage                *NamadaStorage         `json:"storage" toml:"storage"`
	TxIndex                *NamadaTxIndex         `json:"txIndex" toml:"tx_index"`
	Instrumentation        *NamadaInstrumentation `json:"instrumentation" toml:"instrumentation"`
	Statesync              *NamadaStatesync       `json:"statesync" toml:"statesync"`
	Fastsync               *NamadaFastsync        `json:"fastsync" toml:"fastsync"`
}

type NamadaEthereumBridge struct {
	Mode              *string `json:"mode" toml:"mode"`
	OracleRPCEndpoint *string `json:"oracleRpcEndpoint" toml:"oracle_rpc_endpoint"`
	ChannelBufferSize *int    `json:"channelBufferSize" toml:"channel_buffer_size"`
}
