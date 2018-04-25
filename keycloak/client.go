package keycloak

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

func NewKcClient(oauthConfig *oauth2.Config, server string, user string, password string) *KcClient {
	client := &KcClient{
		oauthConfig: oauthConfig,
		server:      server,
	}

	token, err := oauthConfig.PasswordCredentialsToken(context.Background(), user, password)
	if err != nil {
		log.Println("Something went wrong creating a new client.")
		log.Fatalln(err.Error())
	}

	sourceToken := oauthConfig.TokenSource(context.Background(), token)
	client.sourceToken = oauth2.ReuseTokenSource(nil, sourceToken)
	return client
}

func (kc *KcClient) GetUserRolesForClient(realm string, user PairWise, clientID string) ([]RoleRepresentation, error) {
	userID := user.InternalSubject
	url := fmt.Sprintf("%s/admin/realms/%s/users/%s/role-mappings/clients/%s/composite",
		kc.server,
		realm,
		userID,
		clientID,
	)

	httpClient := kc.GetHttpClient()
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := httpClient.Do(req)

	if resp.StatusCode == 200 {
		defer resp.Body.Close()

		roles := []RoleRepresentation{}

		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&roles); err != nil {
			log.Println(err)
		}
		return roles, nil
	}

	err := errors.New("Failed to get roles for this userID.")
	return nil, err
}

func (kc *KcClient) GetMasterRealmUserRoles(userId string) ([]RoleRepresentation, error) {

	url := fmt.Sprintf("%s/admin/realms/master/users/%s/role-mappings/realm/composite",
		kc.server,
		userId,
	)

	httpClient := kc.GetHttpClient()
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()

		roles := []RoleRepresentation{}

		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&roles); err != nil {
			log.Println(err)
		}
		return roles, nil
	}

	err = errors.New("Role name not found for this client.")
	return nil, err
}

func (kc *KcClient) HasRoleInMasterRealm(name string, userId string) bool {
	var roles []RoleRepresentation
	roles, _ = kc.GetMasterRealmUserRoles(userId)

	for _, role := range roles {
		if role.Name == name {
			return true
		}
	}

	return false
}

func (kc *KcClient) GetClientByRoleName(taasRealm string, clientId string, roleName string) (RoleRepresentation, error) {
	var role RoleRepresentation
	url := fmt.Sprintf("%s/admin/realms/%s/clients/%s/roles/%s",
		kc.server,
		taasRealm,
		clientId,
		roleName,
	)

	httpClient := kc.GetHttpClient()
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := httpClient.Do(req)

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()

		err := json.NewDecoder(resp.Body).Decode(&role)
		if err != nil {
			err := errors.New("Failed to convert json object to role.")
			return role, err
		}
		return role, nil
	}

	err := errors.New("Role name not found for this client.")
	return role, err
}

func (kc *KcClient) LinkUserToClientRole(realm string, user PairWise, clientID string, role *RoleRepresentation) bool {
	userID := user.InternalSubject
	url := fmt.Sprintf("%s/admin/realms/%s/users/%s/role-mappings/clients/%s",
		kc.server,
		realm,
		userID,
		clientID,
	)

	httpClient := kc.GetHttpClient()
	roles := []*RoleRepresentation{role}

	b, err := json.Marshal(roles)
	if err != nil {
		log.Println("Failed to marshall roles :(")
		return false
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)

	if err != nil {
		log.Printf("Failed to link User %s, linked user %s with client %s, granting role %s",
			kc.UserWithAccess,
			userID,
			clientID,
			role.Name,
		)
		log.Println(err)
		return false
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNoContent {
		log.Printf("User %s, linked user %s with client %s, granting role %s",
			kc.UserWithAccess,
			userID,
			clientID,
			role.Name,
		)
		return true
	}

	log.Printf("Failed to link User %s, linked user %s with client %s, granting role %s",
		kc.UserWithAccess,
		userID,
		clientID,
		role.Name,
	)
	return false
}

func (kc *KcClient) GetClientRoles(realm string, clientID string) ([]RoleRepresentation, error) {
	url := fmt.Sprintf("%s/admin/realms/%s/clients/%s/roles",
		kc.server,
		realm,
		clientID,
	)

	httpClient := kc.GetHttpClient()
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := httpClient.Do(req)

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()

		roles := []RoleRepresentation{}

		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&roles); err != nil {
			log.Println(err)
		}
		return roles, nil
	}

	err := errors.New("Failed to get roles for this clientID.")
	return nil, err
}

func (kc *KcClient) DeleteClientRoles(realm string, clientID string, roleName string) (bool, ErrorMessageResponseFromKeycloak) {
	var errorMessage ErrorMessageResponseFromKeycloak
	var response bool

	url := fmt.Sprintf("%s/admin/realms/%s/clients/%s/roles/%s",
		kc.server,
		realm,
		clientID,
		roleName,
	)

	httpClient := kc.GetHttpClient()
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)

	if err != nil {
		errorMessage.ErrorMessage = fmt.Sprintf("Failed to DELETE role with name %s to client %s",
			roleName,
			clientID,
		)
		log.Println(errorMessage.ErrorMessage)
		log.Println(err)
		return response, errorMessage
	}

	defer resp.Body.Close()
	response = true
	return response, errorMessage
}

func (kc *KcClient) PostClientRoles(realm string, clientID string, role RoleRepresentation) (RoleRepresentation, ErrorMessageResponseFromKeycloak) {
	var errorMessage ErrorMessageResponseFromKeycloak
	var response RoleRepresentation

	url := fmt.Sprintf("%s/admin/realms/%s/clients/%s/roles",
		kc.server,
		realm,
		clientID,
	)

	b, err := json.Marshal(role)
	if err != nil {
		log.Println("Failed to marshall roles :(")
		errorMessage.ErrorMessage = "Failed to convert the role into json."
		return response, errorMessage
	}

	httpClient := kc.GetHttpClient()
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)

	if err != nil {
		errorMessage.ErrorMessage = fmt.Sprintf("Failed to POST role %s to client %s",
			role.Name,
			clientID,
		)
		log.Println(errorMessage.ErrorMessage)
		log.Println(err)
		return response, errorMessage
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		// Get the role
		log.Printf("Role %s added to client %s",
			role.Name,
			clientID,
		)

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			log.Println(err)
		}

		return response, errorMessage
	}

	if resp.StatusCode == http.StatusConflict {
		if err := json.NewDecoder(resp.Body).Decode(&errorMessage); err != nil {
			log.Println(err)
		}

		return response, errorMessage
	}

	log.Printf("Role %s not added to client %s",
		role.Name,
		clientID,
	)
	errorMessage.ErrorMessage = fmt.Sprintf("Not a 204, but no error. StatusCode is %d", resp.StatusCode)
	return response, errorMessage
}

func (kc *KcClient) GetUserRoleMappings(realm string, userID string) (RoleMappings, error) {
	var roleMappings RoleMappings
	url := fmt.Sprintf("%s/admin/realms/%s/users/%s/role-mappings",
		kc.server,
		realm,
		userID,
	)

	httpClient := kc.GetHttpClient()
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := httpClient.Do(req)

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()

		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&roleMappings); err != nil {
			log.Println(err)
		}
		return roleMappings, nil
	}

	err := errors.New("Failed to get roleMappings for this clientID.")
	return roleMappings, err
}

func (kc *KcClient) GetClientSecret(realm string, userID string, clientID string) (*ClientSecret, error) {
	// @todo we should log this request
	url := fmt.Sprintf("%s/admin/realms/%s/clients/%s/client-secret",
		kc.server,
		realm,
		clientID,
	)

	httpClient := kc.GetHttpClient()
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := httpClient.Do(req)

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()

		clientSecret := ClientSecret{}

		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&clientSecret); err != nil {
			log.Println(err)
		}
		return &clientSecret, nil
	}

	err := errors.New("Failed to get clientSecret for this clientID.")
	return nil, err
}

func (kc *KcClient) GetHttpClient() *http.Client {
	httpClient := kc.oauthConfig.Client(context.Background(), kc.GetToken())
	return httpClient
}

func (kc *KcClient) GetToken() *oauth2.Token {
	token, err := kc.sourceToken.Token()
	if err != nil {
		log.Println("Failed to get the token. Maybe refresh failed.")
		log.Fatalln("Sad times " + err.Error())
	}
	return token
}

func (kc *KcClient) DeleteRealmRole(realm string, roleName string) (bool, ErrorMessageResponseFromKeycloak) {
	var errorMessage ErrorMessageResponseFromKeycloak
	var response bool
	url := fmt.Sprintf("%s/admin/realms/%s/roles/%s",
		kc.server,
		realm,
		roleName,
	)

	httpClient := kc.GetHttpClient()
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)

	if err != nil {
		errorMessage.ErrorMessage = fmt.Sprintf("Failed to DELETE role with name %s from realm %s",
			roleName,
			realm,
		)
		log.Println(errorMessage.ErrorMessage)
		log.Println(err)
		return response, errorMessage
	}

	defer resp.Body.Close()
	response = true
	return response, errorMessage
}

func (kc *KcClient) PostRealmRole(realm string, role RoleRepresentation) (bool, ErrorMessageResponseFromKeycloak) {
	var errorMessage ErrorMessageResponseFromKeycloak
	response := false

	url := fmt.Sprintf("%s/admin/realms/%s/roles",
		kc.server,
		realm,
	)

	b, err := json.Marshal(role)
	if err != nil {
		log.Println("Failed to marshall roles :(")
		errorMessage.ErrorMessage = "Failed to convert the role into json."
		return response, errorMessage
	}

	httpClient := kc.GetHttpClient()
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)

	if err != nil {
		errorMessage.ErrorMessage = fmt.Sprintf("Failed to POST role %s to realm %s",
			role.Name,
			realm,
		)
		log.Println(errorMessage.ErrorMessage)
		log.Println(err)
		return response, errorMessage
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		// Get the role
		log.Printf("Role %s added to realm %s",
			role.Name,
			realm,
		)

		response = true
		return response, errorMessage
	}

	if resp.StatusCode == http.StatusConflict {
		if err := json.NewDecoder(resp.Body).Decode(&errorMessage); err != nil {
			log.Println(err)
		}
		response = false
		return response, errorMessage
	}

	log.Printf("Role %s not added to realm %s",
		role.Name,
		realm,
	)
	errorMessage.ErrorMessage = fmt.Sprintf("Not a 204, but no error. StatusCode is %d", resp.StatusCode)
	return response, errorMessage
}
