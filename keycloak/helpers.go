package keycloak

import "encoding/json"

// Filter client roles by role name
func FilterClientIDByRoleName(roleMappings RoleMappings, filterOn string) []string {
	response := []string{}

	for _, roleMappingClient := range roleMappings.ClientMappings {
		for _, role := range roleMappingClient.Mappings {
			if role.Name == filterOn {
				response = append(response, roleMappingClient.Id)
			}
		}
	}
	return response
}

func NewTestTokenErrorResponse(response string) TokenErrorResponse {
	var errorMessage TokenErrorResponse
	json.Unmarshal([]byte(response), &errorMessage)
	return errorMessage
}
