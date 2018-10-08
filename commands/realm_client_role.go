package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	realmClientCmd.AddCommand(realmClientRoleCmd)
}

var realmClientRoleCmd = &cobra.Command{
	Use:   "role [realm clientId role]",
	Short: "Handle client roles",
	Long:  "Handle client roles: list, add, delete",
}
