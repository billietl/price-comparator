package dao

import "errors"

var (
	// ErrorDAONotFound is used for unknown DAO type
	ErrorDAONotFound = errors.New("unknown DAO type")
)

type DAOBundle struct {
	ProductDAO ProductDAO
	StoreDAO   StoreDAO
}

func GetDAOBundle(daoType string) (bundle *DAOBundle, err error) {
	bundle = &DAOBundle{}
	switch daoType {
	case "firestore":
		initFirestore()
		bundle.ProductDAO = ProductDAOFirestore{}
		bundle.StoreDAO = StoreDAOFirestore{}
		return
	}
	return nil, ErrorDAONotFound
}
