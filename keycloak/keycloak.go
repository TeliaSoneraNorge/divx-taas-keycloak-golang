package keycloak

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
