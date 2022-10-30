package main

import (
	"context"
	"fmt"

	"github.com/k3d-io/k3d/v5/pkg/client"
	"github.com/k3d-io/k3d/v5/pkg/runtimes"
	"github.com/pthomison/errcheck"
)

func main() {

	fmt.Println("hi")

	// cmd := cluster.NewCmdClusterList()
	// cmd.Execute()

	dck_rt, err := runtimes.GetRuntime("docker")
	errcheck.Check(err)

	clusters, err := client.ClusterList(context.TODO(), dck_rt)
	errcheck.Check(err)

	fmt.Printf("%+v\n", clusters[0])
}
