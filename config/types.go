package config

import "fmt"

type PostgreSQLCreds struct {
	Host   string
	Port   uint16
	User   string
	Pass   string
	DbName string
}

type ContainerConfig struct {
	Name        string
	Image       string
	Environment map[string]string
	Command     []string
	Ports       map[uint16]uint16
}

type EndpointWithREST struct {
	// Usually used for health-check
	CoreREST string `toml:"core_rest"`
	Endpoint string `toml:"endpoint"`
}

type Network struct {
	ArtifactsRepository string `toml:"artifacts_repository"`
	GenesisURL          string `toml:"genesis_url"`
	// This is used when We deploy a patch to the mainnet
	BinaryVersionOverride string `toml:"binary_version_override"`

	DataNodesREST  []string           `toml:"data_nodes_rest"`
	RPCPeers       []EndpointWithREST `toml:"rpc_peers"`
	Seeds          []string           `toml:"seeds"`
	BootstrapPeers []EndpointWithREST `toml:"bootstrap_peers"`
}

func (n Network) Validate() error {
	if len(n.DataNodesREST) == 0 {
		return fmt.Errorf("no data nodes rest endpoints")
	}

	if len(n.BootstrapPeers) == 0 {
		return fmt.Errorf("no bootstrap peers")
	}

	if len(n.RPCPeers) == 0 {
		return fmt.Errorf("no rpc peers")
	}

	if len(n.Seeds) == 0 {
		return fmt.Errorf("no seeds")
	}

	if len(n.GenesisURL) == 0 {
		return fmt.Errorf("no genesis url")
	}

	if len(n.ArtifactsRepository) == 0 {
		return fmt.Errorf("empty artifacts repository")
	}

	return nil
}
