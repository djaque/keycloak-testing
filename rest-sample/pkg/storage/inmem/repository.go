package inmem

import (
	"fmt"
	"strings"
	"sync"

	gopher "github.com/djaque/compose-keycloak/rest-sample/pkg"
)

type gopherRepository struct {
	mtx     sync.RWMutex
	gophers map[string]*gopher.Gopher
}

func NewGopherRepository(gophers map[string]*gopher.Gopher) gopher.GopherRepository {
	if gophers == nil {
		gophers = make(map[string]*gopher.Gopher)
	}

	return &gopherRepository{
		gophers: gophers,
	}
}

func (r *gopherRepository) CreateGopher(g *gopher.Gopher) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if err := r.checkIfExists(g.ID); err != nil {
		return err
	}
	r.gophers[g.ID] = g
	return nil
}

func (r *gopherRepository) FetchGophers() ([]*gopher.Gopher, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	values := make([]*gopher.Gopher, 0, len(r.gophers))
	for _, value := range r.gophers {
		values = append(values, value)
	}
	return values, nil
}

func (r *gopherRepository) DeleteGopher(ID string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.gophers, ID)

	return nil
}

func (r *gopherRepository) UpdateGopher(ID string, g *gopher.Gopher) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.gophers[ID] = g
	return nil
}

func (r *gopherRepository) FetchGopherByID(ID string) (*gopher.Gopher, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, v := range r.gophers {
		if v.ID == ID {
			return v, nil
		}
	}

	return nil, fmt.Errorf("the ID %s doesn't exist", ID)
}

func (r *gopherRepository) FetchGopherByName(name string) (*gopher.Gopher, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, v := range r.gophers {
		if strings.EqualFold(v.Name, name) {
			return v, nil
		}
	}

	return nil, fmt.Errorf("the name %s doesn't exist", name)
}

func (r *gopherRepository) checkIfExists(ID string) error {
	for _, v := range r.gophers {
		if v.ID == ID {
			return fmt.Errorf("the gopher %s is already exist", ID)
		}
	}

	return nil
}
