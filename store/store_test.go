package store

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"cloud.google.com/go/firestore"
	"github.com/dchest/uniuri"
)

func Test_StoreNew(t *testing.T) {
	s := New()
	assert.NotEqual(t, "", s.Id, "New stores should have generated ID")
	assert.Equal(t, "", s.Name, "New stores should not have name set")
	assert.Equal(t, "", s.City, "New stores should not have city set")
	assert.Equal(t, "", s.Zipcode, "New stores should not have zipcode set")
}

func Test_StoreLoad(t *testing.T) {
	// new client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "price-comparator-dev")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	defer client.Close()
	// new store
	newStoreId := uniuri.New()
	newStoreName := uniuri.New()
	newStoreCity := uniuri.New()
	newStoreCode := uniuri.New()
	store := client.Collection("store").Doc(newStoreId)
	_, err = store.Create(ctx, map[string]interface{}{
			"name":  newStoreName,
			"city": newStoreCity,
			"zipcode": newStoreCode,
		})
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	// load
	s, err := Load(store.ID)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	// checks
	assert.Equal(t, newStoreId, s.Id)
	assert.Equal(t, newStoreName, s.Name)
	assert.Equal(t, newStoreCity, s.City)
	assert.Equal(t, newStoreCode, s.Zipcode)
	// cleanup
	_, err = store.Delete(ctx)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
}
