package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	realmCmd.AddCommand(realmUserCmd)
}

var realmUserCmd = &cobra.Command{
	Use:   "client [realm clientId]",
	Short: "Operations for a client",
	Long:  "Operations for a client",
}
