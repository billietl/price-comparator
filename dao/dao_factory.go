package dao

import (
	"context"
	"errors"
)

var (
	// ErrorDAONotFound is used for unknown DAO type
	ErrorDAONotFound = errors.New("unknown DAO type")
)

type DAOBundle struct {
	ProductDAO ProductDAO
	StoreDAO   StoreDAO
	Shutdown   func()
}

func GetDAOBundle(ctx context.Context, daoType string) (bundle *DAOBundle, err error) {
	bundle = &DAOBundle{}
	switch daoType {
	case "firestore":
		initFirestore(ctx)
		bundle.ProductDAO = NewProductDAOFirestore()
		bundle.StoreDAO = NewStoreDAOFirestore()
		bundle.Shutdown = shutDownFirestoreClient
		return
	}
	return nil, ErrorDAONotFound
}
