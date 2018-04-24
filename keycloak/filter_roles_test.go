package keycloak_test

import (
	"encoding/json"
	"testing"

	"github.com/TeliaSoneraNorge/divx-taas-keycloak-golang/keycloak"
)

func TestFilterRoleByName(t *testing.T) {
	str := `{
  "realmMappings": [
    {
      "id": "e75ae0e2-e11e-4be3-b825-3002482b5682",
      "name": "telia_view_admin_clients",
      "scopeParamRequired": true,
      "composite": false,
      "clientRole": false,
      "containerId": "telia"
    }
  ],
  "clientMappings": {
    "efe5f41680da475f": {
      "id": "efe5f41680da475f",
      "client": "efe5f41680da475f",
      "mappings": [
        {
          "id": "8dcc49d2-72d6-4ed2-af95-52d9874d5e34",
          "name": "client_admin",
          "description": "a",
          "scopeParamRequired": true,
          "composite": false,
          "clientRole": true,
          "containerId": "efe5f41680da475f"
        }
      ]
    },
    "9b17ee9d920b43f4": {
      "id": "9b17ee9d920b43f4",
      "client": "9b17ee9d920b43f4",
      "mappings": [
        {
          "id": "6964938b-0869-4f62-b255-9a67d061a310",
          "name": "client_admin",
          "scopeParamRequired": true,
          "composite": false,
          "clientRole": true,
          "containerId": "9b17ee9d920b43f4"
        }
      ]
    },
    "account": {
      "id": "6df0bb6e-aa68-48fd-89d2-a7c10460185f",
      "client": "account",
      "mappings": [
        {
          "id": "60926248-783c-4851-93f3-3e6631dbea65",
          "name": "view-profile",
          "description": "${role_view-profile}",
          "scopeParamRequired": false,
          "composite": false,
          "clientRole": true,
          "containerId": "6df0bb6e-aa68-48fd-89d2-a7c10460185f"
        }
      ]
    }
  }
}`
	roleMappings := keycloak.RoleMappings{}
	json.Unmarshal([]byte(str), &roleMappings)

	a := keycloak.FilterClientIDByRoleName(roleMappings, "client_admin")

	if len(a) != 2 {
		t.Errorf("fail length")
	}

	if a[0] != "efe5f41680da475f" {
		t.Errorf("fail index 0")
	}
	if a[1] != "9b17ee9d920b43f4" {
		t.Errorf("fail index 1")
	}
}
