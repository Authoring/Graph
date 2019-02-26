package cmd

import (
	"github.com/Authoring/Graph/api/server"
	"github.com/Authoring/Graph/api/server/router"
	"github.com/Authoring/Graph/daemon"
	"github.com/spf13/cobra"
)

func initStart() {
	var verbose bool

	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the graph daemon",
		Long:  "Start the graph daemon",
		Run: func(cmd *cobra.Command, args []string) {
			var d = daemon.InitDaemon()
			var opts = &router.Options{
				Backend: d,
				Verbose: verbose,
			}
			server.InitServer(opts)
		},
	}

	startCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print the server logs")

	rootCmd.AddCommand(startCmd)
}
