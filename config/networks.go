package config

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
)

const (
	NetworkNameNebula        string = "nebula"
	NetworkNameMainnet       string = "mainnet"
	NetworkNameFairground    string = "fairground"
	NetworkNameStagnet1      string = "stagnet1"
	NetworkNameDevnet1       string = "devnet1"
	NetworkMainnetMirror     string = "mainnet-mirror"
	NetworkMainnetMirrorAlt  string = "mirror"
	NetworkValidatorTestnet  string = "validator-testnet"
	NetworkValidatorsTestnet string = "validators-testnet"
)

func NetworkConfigForGivenInput(envName string, configPath string, workDir string) (*Network, error) {
	if configPath == "" {
		return networkConfigForEnvironmentName(envName)
	}

	var err error
	if strings.HasPrefix(configPath, "http") {
		configPath, err = downloadConfigFile(configPath, workDir)
		if err != nil {
			return nil, fmt.Errorf("failed to config download file: %w", err)
		}
	}

	return loadConfigFromLocalFile(configPath)
}

func loadConfigFromLocalFile(path string) (*Network, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	netCfg := &Network{}
	if err := toml.Unmarshal(data, netCfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return netCfg, nil
}

func downloadConfigFile(uri string, workDir string) (string, error) {
	outputFile := filepath.Join(workDir, "config.toml")

	out, err := os.Create(outputFile)
	if err != nil {
		return "", fmt.Errorf("failed to create the %s file: %w", outputFile, err)
	}
	defer out.Close()

	resp, err := http.Get(uri)
	if err != nil {
		return "", fmt.Errorf("failed to get response from %s: %w", uri, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid response code from %s: expected %d, got %d", uri, http.StatusOK, resp.StatusCode)
	}

	if _, err := io.Copy(out, resp.Body); err != nil {
		return "", fmt.Errorf("cannot copy content of downloaded file into output: %w", err)
	}

	return outputFile, nil
}

func networkConfigForEnvironmentName(envName string) (*Network, error) {
	switch envName {
	case NetworkNameMainnet:
		return &Mainnet, nil
	case NetworkMainnetMirror, NetworkMainnetMirrorAlt:
		return &MainnetMirror, nil
	case NetworkValidatorTestnet, NetworkValidatorsTestnet:
		return &ValidatorsTestnet, nil
	case NetworkNameFairground:
		return &Fairground, nil
	case NetworkNameStagnet1:
		return &Stagnet1, nil
	case NetworkNameDevnet1:
		return &Devnet1, nil
	}

	return nil, fmt.Errorf("unknown network name: expected one of [mainnet, fairground, stagnet1, devnet1], %s got", envName)
}

var (
	Mainnet = Network{
		ArtifactsRepository: "vegaprotocol/vega",
		// BinaryVersionOverride: "v0.75.8-fix.2",
		GenesisURL: "https://raw.githubusercontent.com/vegaprotocol/networks/master/mainnet1/genesis.json",
		DataNodesREST: []string{
			"https://api0.vega.community",
			"https://api1.vega.community",
			"https://api2.vega.community",
			"https://api3.vega.community",
		},
		RPCPeers: []EndpointWithREST{
			{CoreREST: "https://api1.vega.community", Endpoint: "api1.vega.community:26657"},
			{CoreREST: "https://api2.vega.community", Endpoint: "api2.vega.community:26657"},
			{CoreREST: "https://api3.vega.community", Endpoint: "api3.vega.community:26657"},
			{CoreREST: "https://be0.vega.community", Endpoint: "be0.vega.community:26657"},
			{CoreREST: "https://be1.vega.community", Endpoint: "be1.vega.community:26657"},
			{CoreREST: "https://be3.vega.community", Endpoint: "be3.vega.community:26657"},
		},
		Seeds: []string{
			"b0db58f5651c85385f588bd5238b42bedbe57073@13.125.55.240:26656",
			"abe207dae9367995526812d42207aeab73fd6418@18.158.4.175:26656",
			"198ecd046ebb9da0fc5a3270ee9a1aeef57a76ff@144.76.105.240:26656",
			"211e435c2162aedb6d687409d5d7f67399d198a9@65.21.60.252:26656",
			"c5b11e1d819115c4f3974d14f76269e802f3417b@34.88.191.54:26656",
			"61051c21f083ee30c835a34a0c17c5d1ceef3c62@51.178.75.45:26656",
			"b0db58f5651c85385f588bd5238b42bedbe57073@18.192.52.234:26656",
			"36a2ca7bb6a50427be2181c8ebb7f62ac62ebaf5@m2.vega.community:26656",
			"9903c02a0ff881dc369fc7daccb22c1f9680d2dd@api0.vega.community:26656",
			"9903c02a0ff881dc369fc7daccb22c1f9680d2dd@api0.vega.community:26656",
			"32d7380b195c088c0605c5d24bcf15ff1dade05f@api1.vega.community:26656",
			"4f26ec99d3cf6f0e9e973c0a5f3da87d89ec6677@api2.vega.community:26656",
			"eafacd11af53cd9fb2a14eada53485779cbee4ab@api3.vega.community:26656",
			"9de3ca2bbeb62d165d39acbbcf174e7ac3a6b7c9@be3.vega.community:26656",
		},
		BootstrapPeers: []EndpointWithREST{
			{CoreREST: "https://api0.vega.community", Endpoint: "/dns/api0.vega.community/tcp/4001/ipfs/12D3KooWAHkKJfX7rt1pAuGebP9g2BGTT5w7peFGyWd2QbpyZwaw"},
			{CoreREST: "https://api1.vega.community", Endpoint: "/dns/api1.vega.community/tcp/4001/ipfs/12D3KooWDZrusS1p2XyJDbCaWkVDCk2wJaKi6tNb4bjgSHo9yi5Q"},
			{CoreREST: "https://api2.vega.community", Endpoint: "/dns/api2.vega.community/tcp/4001/ipfs/12D3KooWEH9pQd6P7RgNEpwbRyavWcwrAdiy9etivXqQZzd7Jkrh"},
			{CoreREST: "https://api3.vega.community", Endpoint: "/dns/api3.vega.community/tcp/4001/ipfs/12D3KooWHSoYzEqSfUWEXfFbSnmRhWcP2WgZG2GRT8fzZzio5BTY"},
		},
	}

	Fairground = Network{
		ArtifactsRepository: "vegaprotocol/vega",
		GenesisURL:          "https://raw.githubusercontent.com/vegaprotocol/networks-internal/main/fairground/genesis.json",
		DataNodesREST: []string{
			"https://api.n00.testnet.vega.rocks",
			"https://api.n06.testnet.vega.rocks",
			"https://api.n07.testnet.vega.rocks",
		},
		RPCPeers: []EndpointWithREST{
			{CoreREST: "https://n00.testnet.vega.rocks/", Endpoint: "n00.testnet.vega.rocks:26657"},
			{CoreREST: "https://n01.testnet.vega.rocks/", Endpoint: "n01.testnet.vega.rocks:26657"},
			{CoreREST: "https://n02.testnet.vega.rocks/", Endpoint: "n02.testnet.vega.rocks:26657"},
			{CoreREST: "https://n03.testnet.vega.rocks/", Endpoint: "n03.testnet.vega.rocks:26657"},
			{CoreREST: "https://n04.testnet.vega.rocks/", Endpoint: "n04.testnet.vega.rocks:26657"},
			{CoreREST: "https://n05.testnet.vega.rocks/", Endpoint: "n05.testnet.vega.rocks:26657"},
			{CoreREST: "https://n06.testnet.vega.rocks/", Endpoint: "n06.testnet.vega.rocks:26657"},
			{CoreREST: "https://n07.testnet.vega.rocks/", Endpoint: "n07.testnet.vega.rocks:26657"},
		},
		Seeds: []string{
			"e1e741234f05d1067c73457c87420f68994f5acd@be.testnet.vega.rocks:26656",
			"5f6c6fbc805b2c1be8292f836ffdaeee4754695a@n00.testnet.vega.rocks:26656",
			"7274e41e52752991c0ddff1abeb094dd672e5016@n01.testnet.vega.rocks:26656",
			"90d1ab24a3fcb6eee7193f8dc3fc70f08c8dadda@n02.testnet.vega.rocks:26656",
			"fc1559208ba5a6f1e521a1d149554f1adf8782fe@n03.testnet.vega.rocks:26656",
			"d23735e467f834328cde0c371e052efbefa14791@n04.testnet.vega.rocks:26656",
			"f76cc3679e5341537d4c9abd16dd682b0d65ca84@n05.testnet.vega.rocks:26656",
			"029aeae69c4620092097de72d41caf603f54ad54@n06.testnet.vega.rocks:26656",
			"13fa1c91ca159af72a163187d72dd0c3b84df603@n07.testnet.vega.rocks:26656",
		},
		BootstrapPeers: []EndpointWithREST{
			{CoreREST: "https://n00.testnet.vega.rocks", Endpoint: "/dns/n00.testnet.vega.rocks/tcp/4001/ipfs/12D3KooWBL6aFCBT4h8dod679QZpRuwUxizY9AERGFaVoULp8D78"},
			{CoreREST: "https://n06.testnet.vega.rocks", Endpoint: "/dns/n06.testnet.vega.rocks/tcp/4001/ipfs/12D3KooWBqYJfevTSESGabA7uPxxiAPjeiYiiRQAdw3LzMNjUUjM"},
			{CoreREST: "https://n07.testnet.vega.rocks", Endpoint: "/dns/n07.testnet.vega.rocks/tcp/4001/ipfs/12D3KooWMQ5o7wRDMpjgiGhbKwB6Tg3h4Hu8XQGZxBeS7nZEQjeC"},
		},
	}

	Stagnet1 = Network{
		ArtifactsRepository: "vegaprotocol/vega",
		GenesisURL:          "https://raw.githubusercontent.com/vegaprotocol/networks-internal/main/stagnet1/genesis.json",
		DataNodesREST: []string{
			"https://api.n00.stagnet1.vega.rocks",
			"https://api.n05.stagnet1.vega.rocks",
		},
		RPCPeers: []EndpointWithREST{
			{CoreREST: "https://be.stagnet1.vega.rocks", Endpoint: "be.stagnet1.vega.rocks:26657"},
			{CoreREST: "https://n00.stagnet1.vega.rocks", Endpoint: "n00.stagnet1.vega.rocks:26657"},
			{CoreREST: "https://n01.stagnet1.vega.rocks", Endpoint: "n01.stagnet1.vega.rocks:26657"},
			{CoreREST: "https://n02.stagnet1.vega.rocks", Endpoint: "n02.stagnet1.vega.rocks:26657"},
			{CoreREST: "https://n03.stagnet1.vega.rocks", Endpoint: "n03.stagnet1.vega.rocks:26657"},
			{CoreREST: "https://n04.stagnet1.vega.rocks", Endpoint: "n04.stagnet1.vega.rocks:26657"},
			{CoreREST: "https://n05.stagnet1.vega.rocks", Endpoint: "n05.stagnet1.vega.rocks:26657"},
		},
		Seeds: []string{
			"ccdaddd1099c3338f79d732e4b4d75514d8805d9@be.stagnet1.vega.rocks:26656",
			"71a8538c31b8311034eb01d1ec2511f477e62144@n00.stagnet1.vega.rocks:26656",
			"0215d9af2bab228218d1cf4376d4fa4efc0d1dc8@n01.stagnet1.vega.rocks:26656",
			"1ea3d6b578d9f186ab0bc087ae114be1eb6862fe@n02.stagnet1.vega.rocks:26656",
			"bea26a4a9776c7d0e41a33c574e4bd5b9256d0c9@n03.stagnet1.vega.rocks:26656",
			"3b739f5847a0abff3b86b5240ba46f5b6b9673a9@n04.stagnet1.vega.rocks:26656",
			"9131eef85b12b0ac37bb1e4a934ff18ca7461edc@n05.stagnet1.vega.rocks:26656",
		},

		BootstrapPeers: []EndpointWithREST{
			{CoreREST: "https://n00.stagnet1.vega.rocks", Endpoint: "/dns/n00.stagnet1.vega.rocks/tcp/4001/ipfs/12D3KooWP8YaH8ohx44bGyXWxas6wqiCn1cFRidxR3yPrFVd3ZjP"},
			{CoreREST: "https://n05.stagnet1.vega.rocks", Endpoint: "/dns/n05.stagnet1.vega.rocks/tcp/4001/ipfs/12D3KooWAzYpdr3oeXaQ1CWmFystwi7M2pLrPLpRvD2E1XzjSjHG"},
		},
	}

	Devnet1 = Network{
		ArtifactsRepository: "vegaprotocol/vega-dev-releases",
		GenesisURL:          "https://raw.githubusercontent.com/vegaprotocol/networks-internal/main/devnet1/genesis.json",
		DataNodesREST: []string{
			"https://api.n00.devnet1.vega.rocks",
			"https://api.n06.devnet1.vega.rocks",
			"https://api.n07.devnet1.vega.rocks",
		},
		RPCPeers: []EndpointWithREST{
			{CoreREST: "https://n00.devnet1.vega.rocks", Endpoint: "n00.devnet1.vega.rocks:26657"},
			{CoreREST: "https://n01.devnet1.vega.rocks", Endpoint: "n01.devnet1.vega.rocks:26657"},
			{CoreREST: "https://n02.devnet1.vega.rocks", Endpoint: "n02.devnet1.vega.rocks:26657"},
			{CoreREST: "https://n03.devnet1.vega.rocks", Endpoint: "n03.devnet1.vega.rocks:26657"},
			{CoreREST: "https://n04.devnet1.vega.rocks", Endpoint: "n04.devnet1.vega.rocks:26657"},
			{CoreREST: "https://n05.devnet1.vega.rocks", Endpoint: "n05.devnet1.vega.rocks:26657"},
			{CoreREST: "https://n06.devnet1.vega.rocks", Endpoint: "n06.devnet1.vega.rocks:26657"},
		},
		Seeds: []string{
			"a0928bc929506560c66f5ae4fa2f73df3ed8aab8@n01.devnet1.vega.rocks:26656",
			"0c0f1575d159ed02ac05670c333593b2deb4d57e@n02.devnet1.vega.rocks:26656",
			"091cb0675d0f59305d6b72072fe423206bf17048@n03.devnet1.vega.rocks:26656",
			"e475c424a3f20313f5b0911a06b438c850b89066@n04.devnet1.vega.rocks:26656",
			"7f2b12134155929f70ef162a58a8ad5c289eacde@n05.devnet1.vega.rocks:26656",
		},
		BootstrapPeers: []EndpointWithREST{
			{CoreREST: "https://n00.devnet1.vega.rocks", Endpoint: "/dns/n00.devnet1.vega.rocks/tcp/4001/ipfs/12D3KooWBsVeEhCjG2djhpwexZWb76Afd7Nh6gUfpxNBr61KKojj"},
			{CoreREST: "https://n06.devnet1.vega.rocks", Endpoint: "/dns/n06.devnet1.vega.rocks/tcp/4001/ipfs/12D3KooWEbFqpQc2srFtrPcYK5t1e8mfouDutyzwW3XBEPhqYrLi"},
			{CoreREST: "https://n07.devnet1.vega.rocks", Endpoint: "/dns/n07.devnet1.vega.rocks/tcp/4001/ipfs/12D3KooWSjnLDRMwrNxWqyyzkWCkiP7JaHpKkgbNGpo8fWWfkXoy"},
		},
	}

	MainnetMirror = Network{
		ArtifactsRepository: "vegaprotocol/vega",
		GenesisURL:          "https://raw.githubusercontent.com/vegaprotocol/networks-internal/main/mainnet-mirror/genesis.json",
		DataNodesREST: []string{
			"https://api.n00.mainnet-mirror.vega.rocks",
			"https://api.n06.mainnet-mirror.vega.rocks",
		},
		RPCPeers: []EndpointWithREST{
			{CoreREST: "https://n00.mainnet-mirror.vega.rocks", Endpoint: "n00.mainnet-mirror.vega.rocks:26657"},
			{CoreREST: "https://n01.mainnet-mirror.vega.rocks", Endpoint: "n01.mainnet-mirror.vega.rocks:26657"},
			{CoreREST: "https://n02.mainnet-mirror.vega.rocks", Endpoint: "n02.mainnet-mirror.vega.rocks:26657"},
			{CoreREST: "https://n03.mainnet-mirror.vega.rocks", Endpoint: "n03.mainnet-mirror.vega.rocks:26657"},
			{CoreREST: "https://n04.mainnet-mirror.vega.rocks", Endpoint: "n04.mainnet-mirror.vega.rocks:26657"},
			{CoreREST: "https://n05.mainnet-mirror.vega.rocks", Endpoint: "n05.mainnet-mirror.vega.rocks:26657"},
			{CoreREST: "https://n06.mainnet-mirror.vega.rocks", Endpoint: "n06.mainnet-mirror.vega.rocks:26657"},
		},
		Seeds: []string{
			"6b4d261bfbf198e8d7c09fd514f3a0dcd4257e99@n01.mainnet-mirror.vega.rocks:26656",
			"bed5110d707cf760bdb6ab0ef0ddecddef8a1c34@n02.mainnet-mirror.vega.rocks:26656",
			"0fef54f45d60ec194117346910c3f65e88989733@n03.mainnet-mirror.vega.rocks:26656",
			"49e3520fa334106893294d0cfe685a01b7e6f8a9@n04.mainnet-mirror.vega.rocks:26656",
			"23313341785e9a43de0fabba9fc16fe21746350d@n05.mainnet-mirror.vega.rocks:26656",
		},
		BootstrapPeers: []EndpointWithREST{
			{CoreREST: "https://n00.mainnet-mirror.vega.rocks", Endpoint: "/dns/n00.mainnet-mirror.vega.rocks/tcp/4001/ipfs/12D3KooWLTtXvPevvqe2588ZEMDzy7tUmP8JwRZwDs2TNQYQwdyt"},
			{CoreREST: "https://n06.mainnet-mirror.vega.rocks", Endpoint: "/dns/n06.mainnet-mirror.vega.rocks/tcp/4001/ipfs/12D3KooWPgpoSABsN5zH9JzLAmnWMJyA7tycHTp7K1q3nRnm5c2A"},
		},
	}

	ValidatorsTestnet = Network{
		ArtifactsRepository: "vegaprotocol/vega",
		GenesisURL:          "https://raw.githubusercontent.com/vegaprotocol/networks/master/testnet2/genesis.json",
		DataNodesREST: []string{
			"https://api.n00.validators-testnet.vega.rocks",
			"https://api.n02.validators-testnet.vega.rocks",
			"https://api.metabase00.validators-testnet.vega.rocks",
		},
		RPCPeers: []EndpointWithREST{
			{CoreREST: "https://n00.validators-testnet.vega.rocks", Endpoint: "n00.validators-testnet.vega.rocks:26657"},
			{CoreREST: "https://n01.validators-testnet.vega.rocks", Endpoint: "n01.validators-testnet.vega.rocks:26657"},
			{CoreREST: "https://n02.validators-testnet.vega.rocks", Endpoint: "n02.validators-testnet.vega.rocks:26657"},
			{CoreREST: "https://metabase00.validators-testnet.vega.rocks", Endpoint: "metabase00.validators-testnet.vega.rocks:26657"},
		},
		Seeds: []string{
			"fdd97c9dba30ad45d24bee19503ba164378e7676@65.108.77.179:26656",
			"b167d445d103864cc43b8685a5b559c43d7874c2@sn011.validators-testnet.vega.rocks:40116",
			"8d56e01212501d839ab385487347f4b3110f0b29@sn010.validators-testnet.vega.rocks:40106",
			"19ff93fb93e9d1b275cc395b228afba5161abb69@34.88.143.93:26656",
			"bdd14fe2b171deae3850a8022ea672e4b031e61b@146.59.55.53:36656",
			"71b74583f666d14b9422bf76bcc0967da2b8ea1e@5.9.95.147:26656",
			"03d7b7153e33f3109b61854ed4f07da3048479a8@34.88.143.93:26656",
			"bcde8a5e531d2bddf6562b00c868edec6131cbc6@n02.validators-testnet.vega.rocks:26656",
			"5729b2e5f4612718e7bb8fb13293cbcf9e29e745@5.181.190.159:26656",
			"4a848271a1f689f5bea1a6d0634b3ee2ab8879df@metabase00.validators-testnet.vega.xyz:26656",
			"dcd7690daeb1d07c606c3f373db8202f4a96e866@34.88.143.93:26656",
			"16f5f15024530f4a2d966e13bc81d3aaa536f726@54.234.87.229:26656",
			"a07c6df1f34a15db414d8c7de88ac2b4045b53f1@be.validators-testnet.vega.xyz:26656",
			"96cd04a559a06503812388856b0fda3130f2cd83@141.95.97.28:26656",
			"51ffb62faac4256dcae01a4c46c2623a1c19ad1d@51.75.145.104:16456",
			"66a2375c146cf85dd6f0b3c54c337d799225e5db@65.108.57.71:26656",
			"53df354f81d9330500c3a36163434813f7bcbd05@85.207.33.71:26656",
			"b180c59bc8299fa513ea101d257724c87a2e160b@65.21.151.106:26656",
		},
		BootstrapPeers: []EndpointWithREST{
			{CoreREST: "https://n00.validators-testnet.vega.rocks", Endpoint: "/dns/n00.validators-testnet.vega.rocks/tcp/4001/ipfs/12D3KooWQbCMy5echT1sMKwRQh8GJJk5zmHmg6VNg1qEbpysNACN"},
			{CoreREST: "https://n02.validators-testnet.vega.rocks", Endpoint: "/dns/n02.validators-testnet.vega.rocks/tcp/4001/ipfs/12D3KooWHffX2tdw2phH7ai8GCo2K3ehJfnLRATve5otVr4D3ggK"},
			{CoreREST: "https://metabase00.validators-testnet.vega.rocks", Endpoint: "/dns/metabase00.validators-testnet.vega.rocks/tcp/4001/ipfs/12D3KooWKPDZ1s5FM8YewZVeRb9XwaQ7PdaoyD84hFnKmVbn94gN"},
		},
	}
)
