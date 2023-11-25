package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kshwedha/core/common/config"
	"github.com/kshwedha/core/src/cmd/server"
	// "github.com/kshwedha/core/src/cmd/server"
)

// go run src/cmd/main/main.go
func main() {
	environment := flag.String("e", "dev", "server environment")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)

	server.Serve()
}
