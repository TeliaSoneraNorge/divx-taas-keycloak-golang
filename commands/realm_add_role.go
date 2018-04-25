package commands

import (
	"fmt"

	"github.com/TeliaSoneraNorge/divx-taas-keycloak-golang/keycloak"
	"github.com/spf13/cobra"
)

func init() {
	realmRoleCmd.AddCommand(realmRoleAddCmd)
}

var realmRoleAddCmd = &cobra.Command{
	Use:   "add [realm name]",
	Short: "Add a role to a realm.",
	Long:  `Add a role to a realm.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		realm := args[0]
		roleName := args[1]
		client := Initialise()

		role := keycloak.RoleRepresentation{
			Name:               roleName,
			ScopeParamRequired: true,
		}
		_, errorMessage := client.KeycloakClient.PostRealmRole(realm, role)
		if (keycloak.ErrorMessageResponseFromKeycloak{}) != errorMessage {
			fmt.Println(errorMessage.ErrorMessage)
			return
		}

		fmt.Printf("Role %s add to realm %s\n",
			roleName,
			realm)
	},
}
