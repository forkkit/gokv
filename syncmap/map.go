package syncmap

import (
	"sync"

	"github.com/philippgille/gokv/util"
)

// GoMap is a gokv.Store implementation for a simple Go sync.Map.
type Store struct {
	m *sync.Map
}

// Set stores the given object for the given key.
func (m Store) Set(k string, v interface{}) error {
	data, err := util.ToJSON(v)
	if err != nil {
		return err
	}
	m.m.Store(k, data)
	return nil
}

// Get retrieves the stored object for the given key and populates the fields of the object that v points to
// with the values of the retrieved object's values.
func (m Store) Get(k string, v interface{}) (bool, error) {
	data, found := m.m.Load(k)
	if !found {
		return false, nil
	}

	return true, util.FromJSON(data.([]byte), v)
}

// NewGoMap creates a new GoMap.
func NewStore() Store {
	return Store{
		m: &sync.Map{},
	}
}
