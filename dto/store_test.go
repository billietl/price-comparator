package dto

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/xyproto/randomstring"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"testing"
)

func init() {
	randomstring.Seed()
}

func generateStoreTestData(t *testing.T) (id string, result map[string]string) {
	result = map[string]string{}
	result["name"] = randomstring.HumanFriendlyString(10)
	result["city"] = randomstring.HumanFriendlyString(10)
	result["zipcode"] = randomstring.HumanFriendlyString(5)

	ctx := context.Background()
	doc, _, err := firestoreClient.Collection(firestoreStoreCollection).Add(ctx, result)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	id = doc.ID

	return
}

func cleanupStoreTestData(t *testing.T, id string) {
	ctx := context.Background()
	firestoreClient.Collection(firestoreStoreCollection).Doc(id).Delete(ctx)
}

func TestStoreCreate(t *testing.T) {
	ctx := context.Background()
	// Upsert new store
	createdStore := Store{
		Name:    randomstring.HumanFriendlyString(10),
		City:    randomstring.HumanFriendlyString(10),
		Zipcode: randomstring.HumanFriendlyString(5),
	}
	err := createdStore.Upsert()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.NotEqual(t, "", createdStore.ID, "Store ID should have been genetared at upsert time")
	// Reload store
	doc, err := firestoreClient.Collection(firestoreStoreCollection).Doc(createdStore.ID).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	assert.Equal(t, createdStore.Name, docData["name"])
	assert.Equal(t, createdStore.City, docData["city"])
	assert.Equal(t, createdStore.Zipcode, docData["zipcode"])
	// Cleanup test data
	cleanupStoreTestData(t, createdStore.ID)
}

func TestStoreRead(t *testing.T) {
	// Setup test data
	id, testData := generateStoreTestData(t)

	loadedStore := Store{
		ID: id,
	}
	err := loadedStore.Load()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, testData["name"], loadedStore.Name, "Didn't find the right store name")
	assert.Equal(t, testData["city"], loadedStore.City, "Didn't find the right store city")
	assert.Equal(t, testData["zipcode"], loadedStore.Zipcode, "Didn't find the right store zipcode")

	// Cleanup test data
	cleanupStoreTestData(t, id)
}

func TestStoreUpdate(t *testing.T) {
	ctx := context.Background()
	// Setup test data
	id, testData := generateStoreTestData(t)

	// Update store data
	store := Store{
		ID:      id,
		Name:    randomstring.HumanFriendlyString(10),
		City:    randomstring.HumanFriendlyString(10),
		Zipcode: randomstring.HumanFriendlyString(5),
	}
	store.Upsert()

	// Reload data
	doc, err := firestoreClient.Collection(firestoreStoreCollection).Doc(id).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	assert.NotEqual(t, docData["name"], testData["name"])
	assert.NotEqual(t, docData["city"], testData["city"])
	assert.NotEqual(t, docData["zipcode"], testData["zipcode"])

	// Cleanup test data
	cleanupStoreTestData(t, id)
}

func TestStoreDelete(t *testing.T) {
	ctx := context.Background()
	// Setup test data
	id, _ := generateStoreTestData(t)

	// Delete data
	toDeleteStore := Store{
		ID: id,
	}
	toDeleteStore.Delete()
	_, err := firestoreClient.Collection(firestoreStoreCollection).Doc(id).Get(ctx)
	if grpc.Code(err) != codes.NotFound {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestStoreSearch(t *testing.T) {
	// Setup test data
	id, testData := generateStoreTestData(t)

	searchedByNameStore := Store{
		Name: testData["name"],
	}
	storeList, err := searchedByNameStore.Search()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, 1, len(storeList), "Didn't find the right amount of stores")
	assert.Equal(t, testData["name"], storeList[0].Name, "Didn't find the right store name")
	assert.Equal(t, testData["city"], storeList[0].City, "Didn't find the right store city")
	assert.Equal(t, testData["zipcode"], storeList[0].Zipcode, "Didn't find the right store zipcode")

	// Cleanup test data
	cleanupStoreTestData(t, id)
}
