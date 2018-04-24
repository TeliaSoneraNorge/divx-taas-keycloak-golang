package keycloak

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
