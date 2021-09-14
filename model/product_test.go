package model_test

import (
	"math/rand"
	"price-comparator/model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/xyproto/randomstring"
)

func TestNewProductFields(t *testing.T) {
	t.Parallel()

	result := map[string]interface{}{}
	result["name"] = randomstring.HumanFriendlyString(10)
	result["bio"] = rand.Int()%2 == 1
	result["vrac"] = rand.Int()%2 == 1

	p := model.NewProduct(result["name"].(string), result["bio"].(bool), result["vrac"].(bool))

	assert.NotEqual(t, result["id"], "")
	assert.Equal(t, result["name"], p.Name)
	assert.Equal(t, result["bio"], p.Bio)
	assert.Equal(t, result["vrac"], p.Vrac)
}

func TestNewProductUUID(t *testing.T) {
	t.Parallel()

	p1 := model.NewProduct("", false, false)
	p2 := model.NewProduct("", false, false)

	assert.NotEqual(t, p1.ID, p2.ID)
}

func TestProductGenerateID(t *testing.T) {
	t.Parallel()

	product := model.NewProduct("", false, false)
	id1 := product.ID
	product.GenerateID()
	id2 := product.ID

	assert.NotEqual(t, id1, id2)
}

func TestProductEquals(t *testing.T) {
	t.Parallel()

	product1 := &model.Product{
		ID:   uuid.New().String(),
		Name: randomstring.HumanFriendlyString(10),
		Bio:  rand.Int()%2 == 1,
		Vrac: rand.Int()%2 == 1,
	}
	product2 := &model.Product{
		ID:   uuid.New().String(),
		Name: product1.Name,
		Bio:  product1.Bio,
		Vrac: product1.Vrac,
	}
	product3 := &model.Product{
		ID:   product1.ID,
		Name: randomstring.HumanFriendlyString(10),
		Bio:  !product1.Bio,
		Vrac: !product1.Vrac,
	}

	assert.Equal(t, false, product1.Equals(product2))
	assert.Equal(t, true, product1.Equals(product3))

	assert.Equal(t, true, product1.ValueEquals(product2))
	assert.Equal(t, false, product1.ValueEquals(product3))
}
