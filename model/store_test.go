package model

import (
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/xyproto/randomstring"
)

func TestNewStoreFields(t *testing.T) {
	t.Parallel()

	result := map[string]string{}
	result["name"] = randomstring.HumanFriendlyString(10)
	result["city"] = randomstring.HumanFriendlyString(10)
	result["zipcode"] = randomstring.HumanFriendlyString(5)

	s := NewStore(result["name"], result["city"], result["zipcode"])

	assert.Equal(t, result["name"], s.Name)
	assert.Equal(t, result["city"], s.City)
	assert.Equal(t, result["zipcode"], s.Zipcode)
}

func TestNewStoreUUID(t *testing.T) {
	t.Parallel()

	s1 := NewStore("", "", "")
	s2 := NewStore("", "", "")

	assert.NotEqual(t, s1.ID, s2.ID)
}

func TestStoreGenerateID(t *testing.T) {
	t.Parallel()

	Store := NewStore("", "", "")
	id1 := Store.ID
	Store.GenerateID()
	id2 := Store.ID

	assert.NotEqual(t, id1, id2)
}

func TestStoreEquals(t *testing.T) {
	t.Parallel()

	store1 := &Store{
		ID:      uuid.New().String(),
		Name:    randomstring.HumanFriendlyString(10),
		City:    randomstring.HumanFriendlyString(10),
		Zipcode: randomstring.HumanFriendlyString(5),
	}
	store2 := &Store{
		ID:      uuid.New().String(),
		Name:    store1.Name,
		City:    store1.City,
		Zipcode: store1.Zipcode,
	}
	store3 := &Store{
		ID:      store1.ID,
		Name:    randomstring.HumanFriendlyString(10),
		City:    randomstring.HumanFriendlyString(10),
		Zipcode: randomstring.HumanFriendlyString(5),
	}

	assert.Equal(t, false, store1.Equals(store2))
	assert.Equal(t, true, store1.Equals(store3))
}

func TestStoreValueEquals(t *testing.T) {
	t.Parallel()

	store1 := &Store{
		ID:      uuid.New().String(),
		Name:    randomstring.HumanFriendlyString(10),
		City:    randomstring.HumanFriendlyString(10),
		Zipcode: randomstring.HumanFriendlyString(5),
	}
	store2 := &Store{
		ID:      store1.ID,
		Name:    store1.Name,
		City:    store1.City,
		Zipcode: store1.Zipcode,
	}
	store3 := &Store{
		ID:      store1.ID,
		Name:    randomstring.HumanFriendlyString(10),
		City:    randomstring.HumanFriendlyString(10),
		Zipcode: randomstring.HumanFriendlyString(5),
	}

	assert.Equal(t, true, store1.ValueEquals(store2))
	assert.Equal(t, false, store1.ValueEquals(store3))
}
