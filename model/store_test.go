package model

import (
	"testing"

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
