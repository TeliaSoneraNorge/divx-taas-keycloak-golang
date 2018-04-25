package main

import (
	"fmt"

	"github.com/TeliaSoneraNorge/divx-taas-keycloak-golang/commands"
)

func main() {
	err := commands.RootCmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
