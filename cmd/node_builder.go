package cmd

import (
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"

	"github.com/onflow/flow-go/crypto"
	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module"
	"github.com/onflow/flow-go/module/local"
	"github.com/onflow/flow-go/network"
	"github.com/onflow/flow-go/network/p2p"
	"github.com/onflow/flow-go/state/protocol"
	"github.com/onflow/flow-go/state/protocol/events"
)

// NodeBuilder declares the initialization methods needed to bootstrap up a Flow node
type NodeBuilder interface {

	// BaseFlags reads the command line arguments common to all nodes
	BaseFlags()

	// ExtraFlags reads the node specific command line arguments and adds it to the FlagSet
	ExtraFlags(f func(*pflag.FlagSet)) NodeBuilder

	// ParseAndPrintFlags parses all the command line arguments
	ParseAndPrintFlags()

	// Initialize performs all the initialization needed at the very start of a node
	Initialize() NodeBuilder

	// PrintBuildVersionDetails prints the node software build version
	PrintBuildVersionDetails()

	// EnqueueNetworkInit enqueues the default network component
	EnqueueNetworkInit()

	// EnqueueMetricsServerInit enqueues the metrics component
	EnqueueMetricsServerInit()

	// Enqueues the Tracer component
	EnqueueTracer()

	// Module enables setting up dependencies of the engine with the builder context
	Module(name string, f func(builder NodeBuilder) error) NodeBuilder

	// Component adds a new component to the node that conforms to the ReadyDone
	// interface.
	//
	// When the node is run, this component will be started with `Ready`. When the
	// node is stopped, we will wait for the component to exit gracefully with
	// `Done`.
	Component(name string, f func(NodeBuilder) (module.ReadyDoneAware, error)) NodeBuilder

	// MustNot asserts that the given error must not occur.
	// If the error is nil, returns a nil log event (which acts as a no-op).
	// If the error is not nil, returns a fatal log event containing the error.
	MustNot(err error) *zerolog.Event

	// Run initiates all common components (logger, database, protocol state etc.)
	// then starts each component. It also sets up a channel to gracefully shut
	// down each component if a SIGINT is received.
	Run()

	// PreInit registers a new PreInit function.
	// PreInit functions run before the protocol state is initialized or any other modules or components are initialized
	PreInit(f func(node NodeBuilder)) NodeBuilder

	// PostInit registers a new PreInit function.
	// PostInit functions run after the protocol state has been initialized but before any other modules or components
	// are initialized
	PostInit(f func(node NodeBuilder)) NodeBuilder

	// RegisterBadgerMetrics registers all badger related metrics
	RegisterBadgerMetrics()

	// getters
	Config() BaseConfig
	NodeID() flow.Identifier
	Logger() zerolog.Logger
	Me() *local.Local
	Tracer() module.Tracer
	MetricsRegisterer() prometheus.Registerer
	Metrics() Metrics
	DB() *badger.DB
	Storage() Storage
	ProtocolEvents() *events.Distributor
	ProtocolState() protocol.State
	Middleware() *p2p.Middleware
	Network() *p2p.Network
	MsgValidators() []network.MessageValidator
	FvmOptions() []fvm.Option
	NetworkKey() crypto.PrivateKey
	RootBlock() *flow.Block
	RootQC() *flow.QuorumCertificate
	RootSeal() *flow.Seal
	RootChainID() flow.ChainID

	// setters
	SetMsgValidators(validators []network.MessageValidator)
	SetMe(me *local.Local)
	SetMiddleware(m *p2p.Middleware)
	SetNetwork(n *p2p.Network)
}

// BaseConfig is the general config for the NodeBuilder
type BaseConfig struct {
	nodeIDHex             string
	bindAddr              string
	NodeRole              string
	timeout               time.Duration
	datadir               string
	level                 string
	metricsPort           uint
	BootstrapDir          string
	peerUpdateInterval    time.Duration
	unicastMessageTimeout time.Duration
	profilerEnabled       bool
	profilerDir           string
	profilerInterval      time.Duration
	profilerDuration      time.Duration
	tracerEnabled         bool
}