package keycloak

import (
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

type VerySimpleRole struct {
	UserId   string `json:"userId"`
	ClientId string `json:"clientId"`
	RoleName string `json:"roleName"`
}

type PairWise struct {
	InternalSubject string
	ExternalSubject string
}

type RoleRepresentation struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	ScopeParamRequired bool   `json:"scopeParamRequired"`
	Composite          bool   `json:"composite"`
	ClientRole         bool   `json:"clientRole"`
	ContainerId        string `json:"containerId"`
	Description        string `json:"description"`
}

type RealmAccessRoles struct {
	Roles []string `json:"roles"`
}

type ResourceAccessRoles map[string]RealmAccessRoles

type AccessTokenClaimsWithRoles struct {
	RealmAccess    RealmAccessRoles    `json:"realm_access"`
	ResourceAccess ResourceAccessRoles `json:"resource_access"`
	jwt.StandardClaims
}

type KcClient struct {
	oauthConfig    *oauth2.Config
	server         string
	token          *oauth2.Token
	UserWithAccess string
}

type RoleMappings struct {
	RealmMappings  []RoleRepresentation `json:"realmMappings"`
	ClientMappings RoleMappingClient    `json:"clientMappings"`
}

type RoleMappingClient map[string]RoleMap

type RoleMap struct {
	Id       string               `json:"id"`
	Client   string               `json:"client"`
	Mappings []RoleRepresentation `json:"mappings"`
}

type ClientSecret struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
