/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type User struct {

	Authorities []GrantedAuthority `json:"authorities,omitempty"`

	Roles []string `json:"roles,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Email string `json:"email,omitempty"`

	Username string `json:"username,omitempty"`

	CredentialsNonExpired bool `json:"credentialsNonExpired,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	AllowedAccounts []string `json:"allowedAccounts,omitempty"`

	LastName string `json:"lastName,omitempty"`

	AccountNonLocked bool `json:"accountNonLocked,omitempty"`

	AccountNonExpired bool `json:"accountNonExpired,omitempty"`
}
