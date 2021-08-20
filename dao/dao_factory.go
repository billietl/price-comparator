package dao

import (
	"context"
	"errors"
)

var (
	// ErrorDAONotFound is used for unknown DAO type
	ErrorDAONotFound = errors.New("unknown DAO type")
)

type Bundle struct {
	ProductDAO ProductDAO
	StoreDAO   StoreDAO
	Shutdown   func()
}

func GetBundle(ctx context.Context, daoType string) (bundle *Bundle, err error) {
	bundle = &Bundle{}
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
