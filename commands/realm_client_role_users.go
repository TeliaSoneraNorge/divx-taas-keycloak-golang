package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	)

func init() {
	realmCmd.AddCommand(realmClientRoleUsersCmd)
}

var realmClientRoleUsersCmd = &cobra.Command{
	Use:   "roleusers [realm clientId name]",
	Short: "Get users for role",
	Long:  "Get list of users who has given role",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		realm := args[0]
		clientId := args[1]
		rolename := args[2]
		client := Initialise()

		userRepresentation, errorMessage := client.KeycloakClient.GetUsersForRole(realm, clientId, rolename)
		if errorMessage != nil {
			fmt.Println(errorMessage)
			return
		}
		for _, user := range userRepresentation {
			fmt.Printf("UserId %s with name %s:\n", user.Id, user.Name);
		}
	},
}
