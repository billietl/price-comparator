package dto

type Savable interface {
	Load() error
	Upsert() error
	Delete() error
}
