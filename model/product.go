package model

import (
	"github.com/google/uuid"
)

type Product struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Bio  bool   `json:"bio" bson:"bio"`
	Vrac bool   `json:"vrac" bson:"vrac"`
}

func NewProduct(name string, bio bool, vrac bool) *Product {
	return &Product{
		ID:   uuid.New().String(),
		Name: name,
		Bio:  bio,
		Vrac: vrac,
	}
}

func (p *Product) GenerateID() {
	p.ID = uuid.New().String()
}

func (p Product) Equals(product *Product) bool {
	return p.ID == product.ID
}

func (p Product) ValueEquals(product *Product) bool {
	return p.Name == product.Name &&
		p.Bio == product.Bio &&
		p.Vrac == product.Vrac
}
