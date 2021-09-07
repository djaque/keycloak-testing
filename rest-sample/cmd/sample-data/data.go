package sample

import (
	gopher "github.com/djaque/compose-keycloak/rest-sample/pkg"
)

var Gophers = map[string]*gopher.Gopher{
	"Jenny": &gopher.Gopher{
		ID:        "01D3XZ3ZHCP3KG9VT4FGAD8KDR",
		Name:      "Jenny",
		FirstName: "Jenny",
		LastName:  "Jenny",
		Email:     "Jenny@mail.com",

		Enabled:       "true",
		EmailVerified: "true",

		Roles:           []string{"admin"},
		Groups:          []string{"migrated_users"},
		RequiredActions: []string{"CONFIGURE_TOTP", "UPDATE_PASSWORD", "UPDATE_PROFILE", "update_user_locale"},
		Password:        "123456789",
	},
	"Billy": &gopher.Gopher{
		ID:              "01D3XZ7CN92AKS9HAPSZ4D5DP9",
		Name:            "billy",
		FirstName:       "Billy",
		LastName:        "Billy",
		Email:           "Billy@mail.com",
		Enabled:         "true",
		EmailVerified:   "true",
		Roles:           []string{"admin"},
		Groups:          []string{"migrated_users"},
		RequiredActions: []string{},
		Password:        "123456789",
	},
	"Rainbow": &gopher.Gopher{
		ID:              "01D3XZ89NFJZ9QT2DHVD462AC2",
		Name:            "rainbow",
		FirstName:       "Rainbow",
		LastName:        "Rainbow",
		Email:           "Rainbow@mail.com",
		Enabled:         "true",
		EmailVerified:   "true",
		Roles:           []string{"admin"},
		Groups:          []string{"migrated_users"},
		RequiredActions: []string{},
		Password:        "123456789",
	},
	"Bjorn": &gopher.Gopher{
		ID:              "01D3XZ8JXHTDA6XY05EVJVE9Z2",
		Name:            "bjorn",
		FirstName:       "Bjorn",
		LastName:        "Bjorn",
		Email:           "Bjorn@mail.com",
		Enabled:         "true",
		EmailVerified:   "true",
		Roles:           []string{"admin"},
		Groups:          []string{"migrated_users"},
		RequiredActions: []string{},
		Password:        "123123123",
	},
}
