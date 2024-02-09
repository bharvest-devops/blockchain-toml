package blockchain_toml

// ###################################################
// ################### config.toml ###################
// ###################################################

type CosmosConfigFile struct {
	ProxyApp               *string                `json:"proxyApp" toml:"proxy_app"`
	Moniker                *string                `json:"moniker" toml:"moniker"`
	BoolBlockSync          *bool                  `json:"boolBlockSync" toml:"block_sync"`
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
	RPC                    *CosmosRPC             `json:"rpc" toml:"rpc"`
	P2P                    *CosmosP2P             `json:"p2p" toml:"p2p"`
	Mempool                *CosmosMempool         `json:"mempool" toml:"mempool"`
	Statesync              *CosmosStatesync       `json:"statesync" toml:"statesync"`
	Blocksync              *CosmosBlocksync       `json:"blocksync" toml:"blocksync"`
	Consensus              *CosmosConsensus       `json:"consensus" toml:"consensus"`
	Storage                *CosmosStorage         `json:"storage" toml:"storage"`
	TxIndex                *CosmosTxIndex         `json:"txIndex" toml:"tx_index"`
	Instrumentation        *CosmosInstrumentation `json:"instrumentation" toml:"instrumentation"`
}

type CosmosRPC struct {
	Laddr                                *string   `json:"laddr" toml:"laddr"`
	CorsAllowedOrigins                   *[]string `json:"corsAllowedOrigins" toml:"cors_allowed_origins"`
	CorsAllowedMethods                   *[]string `json:"corsAllowedMethods" toml:"cors_allowed_methods"`
	CorsAllowedHeaders                   *[]string `json:"corsAllowedHeaders" toml:"cors_allowed_headers"`
	GrpcLaddr                            *string   `json:"grpcLaddr" toml:"grpc_laddr"`
	GrpcMaxOpenConnections               *int      `json:"grpcMaxOpenConnections" toml:"grpc_max_open_connections"`
	Unsafe                               *bool     `json:"unsafe" toml:"unsafe"`
	MaxOpenConnections                   *int32    `json:"maxOpenConnections" toml:"max_open_connections"`
	MaxSubscriptionClients               *int      `json:"maxSubscriptionClients" toml:"max_subscription_clients"`
	MaxSubscriptionsPerClient            *int      `json:"maxSubscriptionsPerClient" toml:"max_subscriptions_per_client"`
	ExperimentalSubscriptionBufferSize   *int      `json:"experimentalSubscriptionBufferSize" toml:"experimental_subscription_buffer_size"`
	ExperimentalWebsocketWriteBufferSize *int      `json:"experimentalWebsocketWriteBufferSize" toml:"experimental_websocket_write_buffer_size"`
	ExperimentalCloseOnSlowClient        *bool     `json:"experimentalCloseOnSlowClient" toml:"experimental_close_on_slow_client"`
	TimeoutBroadcastTxCommit             *string   `json:"timeoutBroadcastTxCommit" toml:"timeout_broadcast_tx_commit"`
	MaxBodyBytes                         *int      `json:"maxBodyBytes" toml:"max_body_bytes"`
	MaxHeaderBytes                       *int      `json:"maxHeaderBytes" toml:"max_header_bytes"`
	TLSCertFile                          *string   `json:"tlsCertFile" toml:"tls_cert_file"`
	TLSKeyFile                           *string   `json:"tlsKeyFile" toml:"tls_key_file"`
	PprofLaddr                           *string   `json:"pprofLaddr" toml:"pprof_laddr"`
}

type CosmosP2P struct {
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

type CosmosMempool struct {
	Version               *string `json:"version" toml:"version"`
	Recheck               *bool   `json:"recheck" toml:"recheck"`
	Broadcast             *bool   `json:"broadcast" toml:"broadcast"`
	WalDir                *string `json:"walDir" toml:"wal_dir"`
	Size                  *int    `json:"size" toml:"size"`
	MaxTxsBytes           *int    `json:"maxTxsBytes" toml:"max_txs_bytes"`
	CacheSize             *int    `json:"cacheSize" toml:"cache_size"`
	KeepInvalidTxsInCache *bool   `json:"keepInvalidTxsInCache" toml:"keep-invalid-txs-in-cache"`
	MaxTxBytes            *int    `json:"maxTxBytes" toml:"max_tx_bytes"`
	MaxBatchBytes         *int    `json:"maxBatchBytes" toml:"max_batch_bytes"`
	TTLDuration           *string `json:"ttlDuration" toml:"ttl-duration"`
	TTLNumBlocks          *int    `json:"ttlNumBlocks" toml:"ttl-num-blocks"`
}

type CosmosStatesync struct {
	Enable              *bool   `json:"enable" toml:"enable"`
	RPCServers          *string `json:"rpcServers" toml:"rpc_servers"`
	TrustHeight         *uint64 `json:"trustHeight" toml:"trust_height"`
	TrustHash           *string `json:"trustHash" toml:"trust_hash"`
	TrustPeriod         *string `json:"trustPeriod" toml:"trust_period"`
	DiscoveryTime       *string `json:"discoveryTime" toml:"discovery_time"`
	TempDir             *string `json:"tempDir" toml:"temp_dir"`
	ChunkRequestTimeout *string `json:"chunkRequestTimeout" toml:"chunk_request_timeout"`
	ChunkFetchers       *string `json:"chunkFetchers" toml:"chunk_fetchers"`
}

type CosmosBlocksync struct {
	Version *string `json:"version" toml:"version"`
}

type CosmosConsensus struct {
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

type CosmosStorage struct {
	DiscardAbciResponses *bool `json:"discardAbciResponses" toml:"discard_abci_responses"`
}

type CosmosTxIndex struct {
	Indexer  *string `json:"indexer" toml:"indexer"`
	PsqlConn *string `json:"psqlConn" toml:"psql-conn"`
}

type CosmosInstrumentation struct {
	Prometheus           *bool   `json:"prometheus" toml:"prometheus"`
	PrometheusListenAddr *string `json:"prometheusListenAddr" toml:"prometheus_listen_addr"`
	MaxOpenConnections   *int    `json:"maxOpenConnections" toml:"max_open_connections"`
	Namespace            *string `json:"namespace" toml:"namespace"`
}

// ###################################################
// #################### app.toml ####################
// ###################################################

type CosmosAppFile struct {
	MinimumGasPrices    *string        `json:"minimumGasPrices" toml:"minimum-gas-prices"`
	Pruning             *string        `json:"pruning" toml:"pruning"`
	PruningKeepRecent   *string        `json:"pruningKeepRecent" toml:"pruning-keep-recent"`
	PruningKeepEvery    *string        `json:"pruningKeepEvery" toml:"pruning-keep-every"`
	PruningInterval     *string        `json:"pruningInterval" toml:"pruning-interval"`
	HaltHeight          *uint64        `json:"haltHeight" toml:"halt-height"`
	HaltTime            *uint64        `json:"haltTime" toml:"halt-time"`
	MinRetainBlocks     *int           `json:"minRetainBlocks" toml:"min-retain-blocks"`
	InterBlockCache     *bool          `json:"interBlockCache" toml:"inter-block-cache"`
	IndexEvents         []*interface{} `json:"indexEvents" toml:"index-events"`
	IavlCacheSize       *int           `json:"iavlCacheSize" toml:"iavl-cache-size"`
	IavlDisableFastnode *bool          `json:"iavlDisableFastnode" toml:"iavl-disable-fastnode"`
	Telemetry           *Telemetry     `json:"telemetry" toml:"telemetry"`
	API                 *API           `json:"api" toml:"api"`
	Rosetta             *Rosetta       `json:"rosetta" toml:"rosetta"`
	Grpc                *Grpc          `json:"grpc" toml:"grpc"`
	GrpcWeb             *GrpcWeb       `json:"grpcWeb" toml:"grpc-web"`
	StateSync           *StateSync     `json:"stateSync" toml:"state-sync"`
	Store               *Store         `json:"store" toml:"store"`
}

type Telemetry struct {
	ServiceName             *string        `json:"serviceName" toml:"service-name"`
	Enabled                 *bool          `json:"enabled" toml:"enabled"`
	EnableHostname          *bool          `json:"enableHostname" toml:"enable-hostname"`
	EnableHostnameLabel     *bool          `json:"enableHostnameLabel" toml:"enable-hostname-label"`
	EnableServiceLabel      *bool          `json:"enableServiceLabel" toml:"enable-service-label"`
	PrometheusRetentionTime *int           `json:"prometheusRetentionTime" toml:"prometheus-retention-time"`
	GlobalLabels            []*interface{} `json:"globalLabels" toml:"global-labels"`
}

type API struct {
	Enable             *bool   `json:"enable" toml:"enable"`
	Swagger            *bool   `json:"swagger" toml:"swagger"`
	Address            *string `json:"address" toml:"address"`
	MaxOpenConnections *int32  `json:"maxOpenConnections" toml:"max-open-connections"`
	RPCReadTimeout     *int    `json:"rpcReadTimeout" toml:"rpc-read-timeout"`
	RPCWriteTimeout    *int    `json:"rpcWriteTimeout" toml:"rpc-write-timeout"`
	RPCMaxBodyBytes    *int    `json:"rpcMaxBodyBytes" toml:"rpc-max-body-bytes"`
	EnabledUnsafeCors  *bool   `json:"enabledUnsafeCors" toml:"enabled-unsafe-cors"`
}

type Rosetta struct {
	Enable     *bool   `json:"enable" toml:"enable"`
	Address    *string `json:"address" toml:"address"`
	Blockchain *string `json:"blockchain" toml:"blockchain"`
	Network    *string `json:"network" toml:"network"`
	Retries    *int    `json:"retries" toml:"retries"`
	Offline    *bool   `json:"offline" toml:"offline"`
}

type Grpc struct {
	Enable  *bool   `json:"enable" toml:"enable"`
	Address *string `json:"address" toml:"address"`
}

type GrpcWeb struct {
	Enable           *bool   `json:"enable" toml:"enable"`
	Address          *string `json:"address" toml:"address"`
	EnableUnsafeCors *bool   `json:"enableUnsafeCors" toml:"enable-unsafe-cors"`
}

type StateSync struct {
	SnapshotInterval   *int `json:"snapshotInterval" toml:"snapshot-Interval"`
	SnapshotKeepRecent *int `json:"snapshotKeepRecent" toml:"snapshot-keep-recent"`
}

type Store struct {
	Streamers []*interface{} `json:"streamers" toml:"streamers"`
}
