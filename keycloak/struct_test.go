package keycloak_test

import (
	"encoding/json"
	"testing"

	"github.com/TeliaSoneraNorge/divx-taas-keycloak-golang/keycloak"
	"github.com/TeliaSoneraNorge/divx-taas-keycloak-golang/utils"
)

func TestRoleRepresentation(t *testing.T) {
	str := `{
  "id": "8dcc49d2-72d6-4ed2-af95-52d9874d5e34",
  "name": "client_admin",
  "description": "a",
  "scopeParamRequired": true,
  "composite": false,
  "clientRole": true,
  "containerId": "efe5f41680da475f"
}`
	a := keycloak.RoleRepresentation{}
	json.Unmarshal([]byte(str), &a)

	b := keycloak.RoleRepresentation{
		Id:                 "8dcc49d2-72d6-4ed2-af95-52d9874d5e34",
		Name:               "client_admin",
		Description:        "a",
		ScopeParamRequired: true,
		Composite:          false,
		ClientRole:         true,
		ContainerId:        "efe5f41680da475f",
	}
	if a != b {
		t.Errorf("Object does not match")
	}
	utils.PrettyPrintJSON(a, " ")
}

func TestRoleMappingClient(t *testing.T) {
	str := `{
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
    }
  }`
	a := keycloak.RoleMappingClient{}
	json.Unmarshal([]byte(str), &a)

	for key, _ := range a {
		if key != "efe5f41680da475f" {
			t.Errorf("Key not matched")
		}
	}

	utils.PrettyPrintJSON(a, " ")
}

func TestRoleMap(t *testing.T) {
	str := `{
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
    }`
	res := keycloak.RoleMap{}

	json.Unmarshal([]byte(str), &res)

	if res.Id != "efe5f41680da475f" {
		t.Errorf("Key not matched")
	}
	if res.Client != "efe5f41680da475f" {
		t.Errorf("Key not matched")
	}

	utils.PrettyPrintJSON(res, " ")
}

func TestRoleMappings(t *testing.T) {
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
	res := keycloak.RoleMappings{}
	json.Unmarshal([]byte(str), &res)
	if res.ClientMappings["account"].Client != "account" {
		t.Errorf("Failed to get the RoleMap from ClientMappings")
	}

	if res.ClientMappings["account"].Mappings[0].Name != "view-profile" {
		t.Errorf("Failed to get the RoleMap from ClientMappings")
	}
	utils.PrettyPrintJSON(res, " ")
}

func TestTokenErrorResponse(t *testing.T) {
	str := `{
	"error":"invalid_grant",
	"error_description":"Refresh token expired"
}`
	errorMessage := keycloak.NewTestTokenErrorResponse(str)
	if errorMessage.Error != "invalid_grant" {
		t.Fail()
	}
	if errorMessage.ErrorDescription != "Refresh token expired" {
		t.Fail()
	}
}
