package commands

import (
	"github.com/spf13/cobra"
	)

func init() {
	realmClientRoleCmd.AddCommand(realmClientRoleUsersCmd)
}

var realmClientRoleUsersCmd = &cobra.Command{
	Use:   "users [realm clientId username]",
	Short: "Handle users for a realm client role",
	Long:  "Handle operations for get, add and delete users for a realm client role",
}

