package dto

import (
	"context"
	"math/rand"
	"github.com/xyproto/randomstring"
	"testing"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func init() {
	randomstring.Seed()
}

func generateProductTestData(t *testing.T) (id string, result map[string]interface{}) {
	result = map[string]interface{}{}
	result["name"] = randomstring.HumanFriendlyString(10)
	result["bio"] = rand.Int()%2==1

	ctx := context.Background()
	doc, _, err := firestoreClient.Collection(firestoreProductCollection).Add(ctx, result)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	id = doc.ID

	return
}

func cleanupProductTestData(t *testing.T, id string) {
	ctx := context.Background()
	firestoreClient.Collection(firestoreProductCollection).Doc(id).Delete(ctx)
}

func TestProductCreate(t *testing.T) {
	ctx := context.Background()
	// Upsert new product
	createdProduct := Product{
		Name:    randomstring.HumanFriendlyString(10),
		Bio:    rand.Int()%2==1,
	}
	err := createdProduct.Upsert()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.NotEqual(t, "", createdProduct.ID, "Product ID should have been genetared at upsert time")
	// Reload product
	doc, err := firestoreClient.Collection(firestoreProductCollection).Doc(createdProduct.ID).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	assert.Equal(t, createdProduct.Name, docData["name"])
	assert.Equal(t, createdProduct.Bio, docData["bio"])
	// Cleanup test data
	cleanupStoreTestData(t, createdProduct.ID)
}

func TestProductRead(t *testing.T) {
	// Setup test data
	id, testData := generateProductTestData(t)

	loadedProduct := Product{
		ID: id,
	}
	err := loadedProduct.Load()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, testData["name"], loadedProduct.Name, "Didn't find the right product name")
	assert.Equal(t, testData["bio"], loadedProduct.Bio, "Didn't find the right bio label")

	// Cleanup test data
	cleanupStoreTestData(t, id)
}

func TestProductUpdate(t *testing.T) {
	ctx := context.Background()
	// Setup test data
	id, testData := generateProductTestData(t)

	// Update store data
	product := Product{
		ID:      id,
		Name:    randomstring.HumanFriendlyString(10),
		Bio:     ! testData["bio"].(bool),
	}
	product.Upsert()

	// Reload data
	doc, err := firestoreClient.Collection(firestoreProductCollection).Doc(id).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	assert.NotEqual(t, docData["name"], testData["name"])
	assert.NotEqual(t, docData["bio"], testData["bio"])

	// Cleanup test data
	cleanupProductTestData(t, id)
}

func TestProductDelete(t *testing.T) {
	ctx := context.Background()
	// Setup test data
	id, _ := generateProductTestData(t)

	// Delete data
	toDeleteProduct := Product{
		ID: id,
	}
	toDeleteProduct.Delete()
	_, err := firestoreClient.Collection(firestoreProductCollection).Doc(id).Get(ctx)
	if grpc.Code(err) != codes.NotFound {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestProductSearch(t *testing.T) {
	// Setup test data
	id, testData := generateProductTestData(t)

	searchedByNameProduct := Product{
		Name: testData["name"].(string),
	}
	productList, err := searchedByNameProduct.Search()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, 1, len(productList), "Didn't find the right amount of product")
	assert.Equal(t, testData["name"].(string), productList[0].Name, "Didn't find the right product name")
	assert.Equal(t, testData["bio"].(bool), productList[0].Bio, "Didn't find the right product bio label")

	// Cleanup test data
	cleanupProductTestData(t, id)
}
