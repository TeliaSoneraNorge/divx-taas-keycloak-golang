package utils_test

import (
	"testing"

	"github.com/TeliaSoneraNorge/divx-taas-keycloak-golang/utils"
)

func TestRefreshResponse(t *testing.T) {
	errorMessage := `oauth2: cannot fetch token: 400 Bad Request {"error":"invalid_grant","error_description":"Refresh token expired"}`
	invalidGrant := utils.HasRefreshTokenExpired(errorMessage)
	if !invalidGrant {
		t.Fail()
	}
}
