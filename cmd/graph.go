package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "graph",
	Short: "Graph is a new database engine",
}

// Execute launches the command lien
func Execute() {
	initStart()
	initInit()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
