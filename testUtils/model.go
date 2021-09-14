package testUtils

import (
	"math/rand"
	"price-comparator/model"

	"github.com/xyproto/randomstring"
)

func GenerateRandomPrice() *model.Price {
	date := Randate()
	return model.NewPrice(
		GenerateRandomProduct(),
		GenerateRandomStore(),
		rand.Float64(),
		&date,
	)
}

func GenerateRandomProduct() *model.Product {
	return model.NewProduct(
		randomstring.HumanFriendlyString(10),
		rand.Int()%2 == 1,
		rand.Int()%2 == 1,
	)
}

func GenerateRandomStore() *model.Store {
	return model.NewStore(
		randomstring.HumanFriendlyString(10),
		randomstring.HumanFriendlyString(10),
		randomstring.HumanFriendlyString(5),
	)
}
