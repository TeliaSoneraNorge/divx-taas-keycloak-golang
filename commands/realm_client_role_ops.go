package commands

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	realmClientRoleCmd.AddCommand(realmClientRoleGetCmd)
}

var realmClientRoleGetCmd = &cobra.Command{
	Use:   "get [realm clientId role]",
	Short: "Get a client role",
	Long:  "Get a client role",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		realm := args[0]
		clientId := args[1]
		rolename := args[2]
		client := Initialise()

		role, errorMessage := client.KeycloakClient.GetClientByRoleName(realm, clientId, rolename)
		if errorMessage != nil {
			fmt.Println(errorMessage)
			return
		}
		fmt.Printf("Role name %s has id : %s. Description: %s\n", role.Name, role.Id, role.Description);
	},
}

