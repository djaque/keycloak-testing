package sample

import (
	gopher "github.mpi-internal.com/Yapo/keycloak-testing/rest-sample/pkg"
)

var Gophers = map[string]*gopher.Gopher{
	"jenny": &gopher.Gopher{
		ID:        "01D3XZ3ZHCP3KG9VT4FGAD8KDR",
		Name:      "jenny",
		FirstName: "Jenny",
		LastName:  "Jenny",
		Email:     "jenny@mail.com",

		Enabled:       "true",
		EmailVerified: "true",

		Roles:           []string{"admin"},
		Groups:          []string{"migrated_users"},
		RequiredActions: []string{},
		//RequiredActions: []string{"CONFIGURE_TOTP", "UPDATE_PASSWORD", "UPDATE_PROFILE", "update_user_locale"},
		Password: "123456789",
	},
	"billy": &gopher.Gopher{
		ID:              "01D3XZ7CN92AKS9HAPSZ4D5DP9",
		Name:            "billy",
		FirstName:       "billy",
		LastName:        "billy",
		Email:           "billy@mail.com",
		Enabled:         "true",
		EmailVerified:   "true",
		Roles:           []string{"admin"},
		Groups:          []string{"migrated_users"},
		RequiredActions: []string{},
		Password:        "123456789",
	},
	"rainbow": &gopher.Gopher{
		ID:              "01D3XZ89NFJZ9QT2DHVD462AC2",
		Name:            "rainbow",
		FirstName:       "rainbow",
		LastName:        "rainbow",
		Email:           "rainbow@mail.com",
		Enabled:         "true",
		EmailVerified:   "true",
		Roles:           []string{"admin"},
		Groups:          []string{"migrated_users"},
		RequiredActions: []string{},
		Password:        "123456789",
	},
	"diego": &gopher.Gopher{
		ID:              "283920192",
		Name:            "reiby-viper@hotmail.com",
		FirstName:       "Diego",
		LastName:        "Vergara",
		Email:           "reiby-viper@hotmail.com",
		Enabled:         "true",
		EmailVerified:   "true",
		Roles:           []string{"admin"},
		Groups:          []string{"migrated_users"},
		RequiredActions: []string{},
		Password:        "123456789",
	},
	"bjorn": &gopher.Gopher{
		ID:              "01D3XZ8JXHTDA6XY05EVJVE9Z2",
		Name:            "bjorn",
		FirstName:       "bjorn",
		LastName:        "bjorn",
		Email:           "bjorn@mail.com",
		Enabled:         "true",
		EmailVerified:   "true",
		Roles:           []string{"admin"},
		Groups:          []string{"migrated_users"},
		RequiredActions: []string{},
		Password:        "123123123",
		Attributes: map[string][]string{
			"age":   []string{"18"},
			"image": []string{"https://golang.org/doc/gopher/ref.png", "https://golang.org/lib/godoc/images/footer-gopher.jpg"},
		},
	},
}
