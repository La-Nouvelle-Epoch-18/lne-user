package store

import (
	"github.com/go-xorm/xorm"
	"github.com/hashicorp/go-multierror"

	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/types"
)

// compile time static check
var _ Store = &store{}

var _ storeUtils = &store{}

type storeUtils interface {
	// return associated db engine
	// used for debuging
	GetEngine() *xorm.Engine

	// sync database: create table if dont exist or update them
	Sync() error
}

// Store .
type Store interface {
	UserStore
	storeUtils
}

// store type is the abstraction behind data interactions
type store struct {
	Engine *xorm.Engine
}

// New return a store
func New(engine *xorm.Engine) (*store, error) {
	return &store{Engine: engine}, nil
}

// GetEngine return xorm engine
func (s *store) GetEngine() *xorm.Engine {
	return s.Engine
}

// Sync store models - create tables in database
func (s *store) Sync() error {
	// create models list
	models := []interface{}{
		&types.User{},
	}

	var errs error

	// loop over models and create or update their tables
	for _, m := range models {
		if err := s.Engine.Sync2(m); err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	return errs
}
