package main

import (
	"fmt"

	"github.com/vegaprotocol/snapshot-testing/config"
	"github.com/vegaprotocol/snapshot-testing/logging"
	"github.com/vegaprotocol/snapshot-testing/networkutils"
	"go.uber.org/zap"
)

func main() {
	mainLogger := logging.CreateLogger(zap.InfoLevel, "./logs/main.log")

	network, err := networkutils.NewNetwork(mainLogger, config.Mainnet, "./")
	if err != nil {
		panic(err)
	}

	details, err := network.SetupLocalNode()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", details)
}

// cli, err := docker.NewClient()
// if err != nil {
// 	panic(err)
// }

// containerExist, err := cli.ContainerExist(context.TODO(), config.PostgresqlConfig.Name)
// if err != nil {
// 	panic(err)
// }

// if containerExist {
// 	err := cli.ContainerRemoveForce(context.TODO(), config.PostgresqlConfig.Name)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// err = cli.RunContainer(context.TODO(), config.PostgresqlConfig)
// if err != nil {
// 	panic(err)
// }

// go func() {
// 	stream, err := cli.Stdout(context.TODO(), config.PostgresqlConfig.Name, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stream.Close()

// 	scanner := bufio.NewScanner(stream)
// 	for scanner.Scan() {
// 		// loglineBytes := scanner.By

// 		fmt.Printf("POSTGRESQL: %s\n ", scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		panic(err)
// 	}

// }()

// for {
// 	running, err := cli.ContainerRunning(context.TODO(), config.PostgresqlConfig.Name)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if !running {
// 		break
// 	}
// }

// fmt.Println("FINISHED")
