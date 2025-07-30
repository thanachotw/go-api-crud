package cmds

import (
	"go-wallet-api/protocol"

	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "serve-rest",
	Short: "Start the REST API server",
	Long:  "This command starts the REST API server to handle HTTP requests.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ServeREST(); err != nil {
			panic(err)
		}
	},
}

func ServeREST() error {
	return protocol.ServeREST()
}
