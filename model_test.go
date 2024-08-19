package gocloak

import (
	"reflect"
	"testing"
)

func TestRequestingPartyTokenOptionsFormData(t *testing.T) {
	tests := []struct {
		name     string
		input    *RequestingPartyTokenOptions
		expected map[string]string
	}{
		{
			name:  "Empty input",
			input: &RequestingPartyTokenOptions{},
			expected: map[string]string{
				"grant_type":                     "urn:ietf:params:oauth:grant-type:uma-ticket",
				"response_include_resource_name": "true",
			},
		},
		{
			name: "With grant type and response include resource name",
			input: &RequestingPartyTokenOptions{
				GrantType:                   ptr("custom_grant_type"),
				ResponseIncludeResourceName: ptr(false),
			},
			expected: map[string]string{
				"grant_type":                     "custom_grant_type",
				"response_include_resource_name": "false",
			},
		},
		{
			name: "With various field types",
			input: &RequestingPartyTokenOptions{
				Ticket:                        ptr("ticket123"),
				PermissionResourceMatchingURI: ptr(true),
				ResponsePermissionsLimit:      ptr(uint32(10)),
			},
			expected: map[string]string{
				"grant_type":                       "urn:ietf:params:oauth:grant-type:uma-ticket",
				"response_include_resource_name":   "true",
				"ticket":                           "ticket123",
				"permission_resource_matching_uri": "true",
				"response_permissions_limit":       "10",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.FormData()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FormData() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStringer(t *testing.T) {
	// nested structs
	actions := []string{"someAction", "anotherAction"}
	access := gocloak.AccessRepresentation{
		Manage:      gocloak.BoolP(true),
		Impersonate: gocloak.BoolP(false),
	}
	v := gocloak.PermissionTicketDescriptionRepresentation{
		ID:               gocloak.StringP("someID"),
		CreatedTimeStamp: gocloak.Int64P(1607702613),
		Enabled:          gocloak.BoolP(true),
		RequiredActions:  &actions,
		Access:           &access,
	}

	str := v.String()

	expectedStr := `{
	"id": "someID",
	"createdTimestamp": 1607702613,
	"enabled": true,
	"requiredActions": [
		"someAction",
		"anotherAction"
	],
	"access": {
		"impersonate": false,
		"manage": true
	}
}`

	assert.Equal(t, expectedStr, str)

	// nested arrays
	config := make(map[string]string)
	config["bar"] = "foo"
	config["ping"] = "pong"

	pmappers := []gocloak.ProtocolMapperRepresentation{
		{
			Name:   gocloak.StringP("someMapper"),
			Config: &config,
		},
	}
	clients := []gocloak.Client{
		{
			Name:            gocloak.StringP("someClient"),
			ProtocolMappers: &pmappers,
		},
		{
			Name: gocloak.StringP("AnotherClient"),
		},
	}

	realmRep := gocloak.RealmRepresentation{
		DisplayName: gocloak.StringP("someRealm"),
		Clients:     &clients,
	}

	str = realmRep.String()
	expectedStr = `{
	"clients": [
		{
			"name": "someClient",
			"protocolMappers": [
				{
					"config": {
						"bar": "foo",
						"ping": "pong"
					},
					"name": "someMapper"
				}
			]
		},
		{
			"name": "AnotherClient"
		}
	],
	"displayName": "someRealm"
}`
	assert.Equal(t, expectedStr, str)
}

type Stringable interface {
	String() string
}

func TestStringerOmitEmpty(t *testing.T) {
	customs := []Stringable{
		&gocloak.CertResponseKey{},
		&gocloak.CertResponse{},
		&gocloak.IssuerResponse{},
		&gocloak.ResourcePermission{},
		&gocloak.PermissionResource{},
		&gocloak.PermissionScope{},
		&gocloak.IntroSpectTokenResult{},
		&gocloak.User{},
		&gocloak.SetPasswordRequest{},
		&gocloak.Component{},
		&gocloak.KeyStoreConfig{},
		&gocloak.ActiveKeys{},
		&gocloak.Key{},
		&gocloak.Attributes{},
		&gocloak.Access{},
		&gocloak.UserGroup{},
		&gocloak.ExecuteActionsEmail{},
		&gocloak.Group{},
		&gocloak.GroupsCount{},
		&gocloak.GetGroupsParams{},
		&gocloak.CompositesRepresentation{},
		&gocloak.Role{},
		&gocloak.GetRoleParams{},
		&gocloak.ClientMappingsRepresentation{},
		&gocloak.MappingsRepresentation{},
		&gocloak.ClientScope{},
		&gocloak.ClientScopeAttributes{},
		&gocloak.ProtocolMappers{},
		&gocloak.ProtocolMappersConfig{},
		&gocloak.Client{},
		&gocloak.ResourceServerRepresentation{},
		&gocloak.RoleDefinition{},
		&gocloak.PolicyRepresentation{},
		&gocloak.RolePolicyRepresentation{},
		&gocloak.JSPolicyRepresentation{},
		&gocloak.ClientPolicyRepresentation{},
		&gocloak.TimePolicyRepresentation{},
		&gocloak.UserPolicyRepresentation{},
		&gocloak.AggregatedPolicyRepresentation{},
		&gocloak.GroupPolicyRepresentation{},
		&gocloak.GroupDefinition{},
		&gocloak.ResourceRepresentation{},
		&gocloak.ResourceOwnerRepresentation{},
		&gocloak.ScopeRepresentation{},
		&gocloak.ProtocolMapperRepresentation{},
		&gocloak.UserInfoAddress{},
		&gocloak.UserInfo{},
		&gocloak.RolesRepresentation{},
		&gocloak.RealmRepresentation{},
		&gocloak.MultiValuedHashMap{},
		&gocloak.TokenOptions{},
		&gocloak.UserSessionRepresentation{},
		&gocloak.SystemInfoRepresentation{},
		&gocloak.MemoryInfoRepresentation{},
		&gocloak.ServerInfoRepresentation{},
		&gocloak.FederatedIdentityRepresentation{},
		&gocloak.IdentityProviderRepresentation{},
		&gocloak.GetResourceParams{},
		&gocloak.GetScopeParams{},
		&gocloak.GetPolicyParams{},
		&gocloak.GetPermissionParams{},
		&gocloak.GetUsersByRoleParams{},
		&gocloak.PermissionRepresentation{},
		&gocloak.CreatePermissionTicketParams{},
		&gocloak.PermissionTicketDescriptionRepresentation{},
		&gocloak.AccessRepresentation{},
		&gocloak.PermissionTicketResponseRepresentation{},
		&gocloak.PermissionTicketRepresentation{},
		&gocloak.PermissionTicketPermissionRepresentation{},
		&gocloak.PermissionGrantParams{},
		&gocloak.PermissionGrantResponseRepresentation{},
		&gocloak.GetUserPermissionParams{},
		&gocloak.ResourcePolicyRepresentation{},
		&gocloak.GetResourcePoliciesParams{},
		&gocloak.CredentialRepresentation{},
		&gocloak.GetUsersParams{},
		&gocloak.GetComponentsParams{},
		&gocloak.GetClientsParams{},
		&gocloak.RequestingPartyTokenOptions{},
		&gocloak.RequestingPartyPermission{},
		&gocloak.GetClientUserSessionsParams{},
		&gocloak.GetOrganizationsParams{},
		&gocloak.OrganizationDomainRepresentation{},
		&gocloak.OrganizationRepresentation{},
	}

	for _, custom := range customs {
		assert.Equal(t, "{}", custom.String())
	}
// Helper function for creating pointers to values
func ptr[T any](v T) *T {
	return &v
}
