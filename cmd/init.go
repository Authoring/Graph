package cmd

import (
	"github.com/Authoring/Graph/engine"
	"github.com/spf13/cobra"
)

func initInit() {
	var verbose bool

	var initCmd = &cobra.Command{
		Use:   "init [name of instance]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Initializes a new instance of Graph",
		Long:  "Initializes a new instance of Graph. Previous instances will not be erased",
		Run: func(cmd *cobra.Command, args []string) {
			var e = &engine.Engine{}
			e.LoadOrCreate(args[0])
		},
	}

	initCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print the steps to initialize a new Graph instance")

	rootCmd.AddCommand(initCmd)
}
