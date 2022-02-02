package main

import (
	"fmt"
	"os"

	"github.com/TechPay-io/go-photon/cmd/photon/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
