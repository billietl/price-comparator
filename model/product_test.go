package model

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyproto/randomstring"
)

func TestNewProductFields(t *testing.T) {
	t.Parallel()

	result := map[string]interface{}{}
	result["name"] = randomstring.HumanFriendlyString(10)
	result["bio"] = rand.Int()%2 == 1

	p := NewProduct(result["name"].(string), result["bio"].(bool))

	assert.NotEqual(t, result["id"], "")
	assert.Equal(t, result["name"], p.Name)
	assert.Equal(t, result["bio"], p.Bio)
}

func TestNewProductUUID(t *testing.T) {
	t.Parallel()

	p1 := NewProduct("", false)
	p2 := NewProduct("", false)

	assert.NotEqual(t, p1.ID, p2.ID)
}
