package commands

import (
	"fmt"

	"github.com/TeliaSoneraNorge/divx-taas-keycloak-golang/keycloak"
	"github.com/spf13/cobra"
)

func init() {
	realmRoleCmd.AddCommand(realmRoleDeleteCmd)
}

var realmRoleDeleteCmd = &cobra.Command{
	Use:   "delete realm name",
	Short: "Work with the realm.",
	Long:  `Work with the realm.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		realm := args[0]
		roleName := args[1]
		client := Initialise()

		_, errorMessage := client.KeycloakClient.DeleteRealmRole(realm, roleName)
		if (keycloak.ErrorMessageResponseFromKeycloak{}) != errorMessage {
			fmt.Println(errorMessage.ErrorMessage)
			return
		}

		fmt.Printf("Role %s deleted from realm %s\n",
			roleName,
			realm)
	},
}
