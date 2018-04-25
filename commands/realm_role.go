package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	realmCmd.AddCommand(realmRoleCmd)
}

var realmRoleCmd = &cobra.Command{
	Use:   "role",
	Short: "Work with the realm role.",
	Long:  `Work with the realm role.`,
}
