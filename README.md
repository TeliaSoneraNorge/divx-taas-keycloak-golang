# divx-taas-keycloak-golang
Used to interact with Telia TAAS Keycloak.

# The api command line tool

```
$ go run cmd/cli/main.go -h
Trust as a service, api.

Usage:
  realm [command]

Available Commands:
  client         Work with a client
    get          Get client info
    role         Work with a client role
      get        Get client role info
      users
        get      Get list of users who has this role
        add      Add role to user
        delete   Delete role from user
  role           Work with realm level roles
    add          Add a role
    delete       Delete a role

Examples:
  The examples assumes that the file config.json exists with appropriate flags
    -- Show help for realm client
    go run cmd/cli/main.go realm client role -h

    -- add client role pairwise1/client_owner to a user
    go run cmd/cli/main.go realm client role users add telia pairwise1 client_owner 55caff40-9d74-4fab-ade2-545aadeca8ed


Flags:
  -h, --help                       help for api
      --keycloak-password string   password for the admin-user. (default "change")
      --keycloak-realm string      Realm to use (default "telia")
      --keycloak-server string     Server to talk to. (default "https://staging.login.telia.io")
      --keycloak-user string       user who will access keycloak admin-cli (default "change")

  The flags may be enterd in a config.json file in order not to clutter the interface:
  {
      "keycloak-user" : "jupiter",
      "keycloak-password" : "god",
      "keycloak-server" : "http://localhost:8000",
      "keycloak-realm" : "master"
  }

Use "api [command] --h" for more information about a command.
```
