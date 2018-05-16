package commands

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	realmCmd.AddCommand(realmClientCmd)
}

var realmClientCmd = &cobra.Command{
	Use:   "client [realm clientId]",
	Short: "Get information about a client",
	Long:  "Get information about a client",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		realm := args[0]
		clientId := args[1]
		client := Initialise()

		clientInformation, errorMessage := client.KeycloakClient.GetClientInformation(realm, clientId)
		if errorMessage != nil {
			fmt.Println(errorMessage)
			return
		}

		fmt.Printf("ClientId %s has name %s and description %s\n",
			clientId,
			clientInformation.Name,
			clientInformation.Description)
	},
}
