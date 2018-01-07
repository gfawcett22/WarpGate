package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version        string
	configFilePath string
	versionFlag    bool
	//root command
	rootCmd = &cobra.Command{
		Use:   "warpgate",
		Short: "WarpGate is an API Gateway",
		Long:  "WarpGate is a lightweight API Gateway.",
		Run:   RunServer,
	}

	checkCmd = &cobra.Command{
		Use:   "check [config-file]",
		Short: "Check the validity of the configuration file",
		Args:  cobra.MinimumNArgs(1),
		Run:   RunCheck,
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the WarpGate version",
		Run:   RunVersion,
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
