package commands

import (
	"context"
	"log"
	"strings"

	"github.com/TeliaSoneraNorge/divx-taas-keycloak-golang/keycloak"
	oidc "github.com/coreos/go-oidc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var (
	RootCmd = &cobra.Command{
		Use:           "api",
		Short:         "Trust as a service, api.",
		Long:          `Trust as a service, api.`,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {
	viper.AutomaticEnv()

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		log.Printf("No config found to read: %s \n", err)
	}

	RootCmd.PersistentFlags().String("keycloak-user", "change", "user who will access keycloak admin-cli")
	RootCmd.PersistentFlags().String("keycloak-password", "change", "password for the admin-user.")
	RootCmd.PersistentFlags().String("keycloak-server", "https://staging.login.telia.io", "Server to talk to.")
	RootCmd.PersistentFlags().String("keycloak-realm", "telia", "Realm to use")

	viper.BindPFlag("keycloak-user", RootCmd.PersistentFlags().Lookup("keycloak-user"))
	viper.BindPFlag("keycloak-password", RootCmd.PersistentFlags().Lookup("keycloak-password"))
	viper.BindPFlag("keycloak-server", RootCmd.PersistentFlags().Lookup("keycloak-server"))
	viper.BindPFlag("keycloak-realm", RootCmd.PersistentFlags().Lookup("keycloak-realm"))
}

func Initialise() *keycloak.KeycloakOidcClient {
	user := viper.GetString("keycloak-user")
	password := viper.GetString("keycloak-password")
	server := viper.GetString("keycloak-server")
	realm := viper.GetString("keycloak-realm")
	ctx := context.Background()
	oidcProvider := server + "/realms/" + realm

	provider, err := oidc.NewProvider(ctx, oidcProvider)
	if err != nil {
		log.Fatalln("Failed to get an oidc Provider")
		log.Fatalln("Error %s", err)
	}
	oidcConfig := &oidc.Config{
		SkipClientIDCheck: true,
	}
	verifier := provider.Verifier(oidcConfig)

	oauthConfig := &oauth2.Config{
		ClientID: "admin-cli",
		Endpoint: oauth2.Endpoint{
			AuthURL:  oidcProvider + "/protocol/openid-connect/auth",
			TokenURL: oidcProvider + "/protocol/openid-connect/token",
		},
	}

	keycloakClient := keycloak.NewKcClient(oauthConfig, server, user, password)

	keycloakOidcClient := keycloak.KeycloakOidcClient{
		KeycloakClient: keycloakClient,
		Provider:       provider,
		Verifier:       verifier,
		Server:         server,
	}
	return &keycloakOidcClient
}
