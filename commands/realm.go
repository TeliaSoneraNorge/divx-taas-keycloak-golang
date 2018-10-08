package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(realmCmd)
}

var realmCmd = &cobra.Command{
	Use:   "realm",
	Short: "Work with the realm.",
	Long:  `Work with the realm.`,
	
}
