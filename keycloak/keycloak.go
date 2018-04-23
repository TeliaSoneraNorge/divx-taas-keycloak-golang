package keycloak

import jwt "github.com/dgrijalva/jwt-go"

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
