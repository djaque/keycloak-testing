package gopher

// Gopher defines the properties of a gopher to be listed
type Gopher struct {
	ID        string `json:"id"`
	Name      string `json:"username,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`

	Enabled       string              `json:"enabled,omitempty"`
	EmailVerified string              `json:"emailVerified,omitempty"`
	Attributes    map[string][]string `json:"attributes,omitempty"`

	Roles           []string `json:"roles,omitempty"`
	Groups          []string `json:"groups,omitempty"`
	RequiredActions []string `json:"requiredActions,omitempty"`

	Password string `json:"-"`
}

// Repository provides access to the gopher storage
type GopherRepository interface {
	// CreateGopher saves a given gopher
	CreateGopher(g *Gopher) error
	// FetchGophers return all gophers saved in storage
	FetchGophers() ([]*Gopher, error)
	// DeleteGopher remove gopher with given ID
	DeleteGopher(ID string) error
	// UpdateGopher modify gopher with given ID and given new data
	UpdateGopher(ID string, g *Gopher) error
	// FetchGopherByID returns the gopher with given ID
	FetchGopherByID(ID string) (*Gopher, error)
	// FetchGopherByName returns the gopher with given name
	FetchGopherByName(name string) (*Gopher, error)
}
