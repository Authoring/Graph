package cmd

import (
	"github.com/Authoring/Graph/engine"
	"github.com/Authoring/Graph/engine/node"
	"github.com/spf13/cobra"
)

func initCreate() {
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "The Create namespace",
	}

	var createNodeCmd = &cobra.Command{
		Use:   "node [name]",
		Short: "Create a new node",
		Long:  "Creates a new node with the given name",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var e = &engine.Engine{}
			e.LoadDefault()
			var n = node.NewNode(e)
			n.Name = args[0]
			n.Save()
		},
	}

	createCmd.AddCommand(createNodeCmd)

	rootCmd.AddCommand(createCmd)
}
