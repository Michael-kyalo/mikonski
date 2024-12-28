package main

import (
	"github.com/Michael-kyalo/mikonski/cmd"
	"github.com/Michael-kyalo/mikonski/pkg/logging"
)

func main() {
	//initalize logging
	logging.InitLogger()
	defer logging.Sync() // flush log buffer on exit

	// Execute the root command to start the CLI
	cmd.Execute()
}
