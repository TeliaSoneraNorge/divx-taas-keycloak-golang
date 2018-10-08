package commands

import (
	"fmt"

	"github.com/TeliaSoneraNorge/divx-taas-keycloak-golang/keycloak"
	"github.com/spf13/cobra"
	)

func init() {
	realmClientRoleUsersCmd.AddCommand(realmClientRoleUsersGetOpsCmd)
	realmClientRoleUsersCmd.AddCommand(realmClientRoleUsersAddOpsCmd)
	realmClientRoleUsersCmd.AddCommand(realmClientRoleUsersDeleteOpsCmd)
}

var realmClientRoleUsersGetOpsCmd = &cobra.Command{
	Use:   "get [realm clientId rolename]",
	Short: "Get users for role",
	Long:  "Get list of users who has a given role",
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
			fmt.Printf("UserId %s with name %s and email %s\n", user.Id, user.Name, user.Email);
		}
	},
}

var realmClientRoleUsersAddOpsCmd = &cobra.Command{
	Use:   "add [realm clientId rolename userId]",
	Short: "Add user for role",
	Long:  "Add a user with role to an existing client. Parameters <realm> <clientId> <rolename> <userid>",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		realm := args[0]
		clientId := args[1]
		rolename := args[2]
		userId := args[3]
		client := Initialise()

		role, errorMessage := client.KeycloakClient.GetClientByRoleName(realm, clientId, rolename)
		if errorMessage != nil {
			fmt.Println(errorMessage)
			return
		}
		pairwise := keycloak.PairWise { InternalSubject : userId }

		ok := client.KeycloakClient.LinkUserToClientRole("<Interactive>", realm, pairwise, clientId, &role)
		if ok {
			fmt.Printf("Success\n")
		} else {
			fmt.Printf("Failure\n")
		}
	},
}

var realmClientRoleUsersDeleteOpsCmd = &cobra.Command{
	Use:   "delete [realm clientId rolename userId]",
	Short: "Remove client role from user",
	Long:  "Remove a client role from an existing user. Parameters <realm> <clientId> <rolename> <userId>",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		realm := args[0]
		clientId := args[1]
		rolename := args[2]
		userId := args[3]
		client := Initialise()

		role, errorMessage := client.KeycloakClient.GetClientByRoleName(realm, clientId, rolename)
		if errorMessage != nil {
			fmt.Println(errorMessage)
			return
		}

		pairwise := keycloak.PairWise { InternalSubject : userId }
		ok := client.KeycloakClient.UnlinkUserFromClientRole("<Interactive>", realm, pairwise, clientId, &role)
		if ok {
			fmt.Printf("Success\n")
		} else {
			fmt.Printf("Failure\n")
		}
	},
}
