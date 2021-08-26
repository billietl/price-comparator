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
	PriceDAO   PriceDAO
	Shutdown   func()
}

func GetBundle(ctx context.Context, daoType string) (bundle *Bundle, err error) {
	bundle = &Bundle{}
	switch daoType {
	case "firestore":
		initFirestore(ctx)
		bundle.ProductDAO = NewProductDAOFirestore()
		bundle.StoreDAO = NewStoreDAOFirestore()
		// bundle.PriceDAO = NewPriceDAOFirestore()
		bundle.Shutdown = shutDownFirestoreClient
		return
	}
	return nil, ErrorDAONotFound
}
