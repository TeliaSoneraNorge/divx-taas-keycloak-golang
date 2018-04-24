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
	client.token = token
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

	httpClient := kc.oauthConfig.Client(context.Background(), kc.token)
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

	httpClient := kc.oauthConfig.Client(context.Background(), kc.token)
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

	err := errors.New("Role name not found for this client.")
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

func (kc *KcClient) GetClientByRoleName(taasRealm string, clientId string, roleName string) (*RoleRepresentation, error) {

	url := fmt.Sprintf("%s/admin/realms/%s/clients/%s/roles/%s",
		kc.server,
		taasRealm,
		clientId,
		roleName,
	)

	httpClient := kc.oauthConfig.Client(context.Background(), kc.token)
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := httpClient.Do(req)

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()

		var role RoleRepresentation
		err := json.NewDecoder(resp.Body).Decode(&role)
		if err != nil {
			err := errors.New("Failed to convert json object to role.")
			return nil, err
		}
		return &role, nil
	}

	err := errors.New("Role name not found for this client.")
	return nil, err
}

func (kc *KcClient) LinkUserToClientRole(realm string, user PairWise, clientID string, role *RoleRepresentation) bool {
	userID := user.InternalSubject
	url := fmt.Sprintf("%s/admin/realms/%s/users/%s/role-mappings/clients/%s",
		kc.server,
		realm,
		userID,
		clientID,
	)

	httpClient := kc.oauthConfig.Client(context.Background(), kc.token)
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

	httpClient := kc.oauthConfig.Client(context.Background(), kc.token)
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

func (kc *KcClient) GetUserRoleMappings(realm string, userID string) (*RoleMappings, error) {
	url := fmt.Sprintf("%s/admin/realms/%s/users/%s/role-mappings",
		kc.server,
		realm,
		userID,
	)

	httpClient := kc.oauthConfig.Client(context.Background(), kc.token)
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := httpClient.Do(req)

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()

		roleMappings := RoleMappings{}

		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&roleMappings); err != nil {
			log.Println(err)
		}
		return &roleMappings, nil
	}

	err := errors.New("Failed to get roleMappings for this clientID.")
	return nil, err
}

func (kc *KcClient) GetClientSecret(realm string, userID string, clientID string) (*ClientSecret, error) {
	// @todo we should log this request
	url := fmt.Sprintf("%s/admin/realms/%s/clients/%s/client-secret",
		kc.server,
		realm,
		clientID,
	)

	httpClient := kc.oauthConfig.Client(context.Background(), kc.token)
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
