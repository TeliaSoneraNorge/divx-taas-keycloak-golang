# divx-taas-keycloak-golang
Used to interact with Telia TAAS Keycloak.

# The api command line tool

```
$ go run cmd/cli/main.go -h
Trust as a service, api.

Usage:
  api [command]

Available Commands:
  help        Help about any command
  realm       Work with the realm.

Flags:
  -h, --help                       help for api
      --keycloak-password string   password for the admin-user. (default "change")
      --keycloak-realm string      Realm to use (default "telia")
      --keycloak-server string     Server to talk to. (default "https://staging.login.telia.io")
      --keycloak-user string       user who will access keycloak admin-cli (default "change")

Use "api [command] --help" for more information about a command.
```
