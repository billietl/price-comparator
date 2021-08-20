package dao

import (
	"context"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyproto/randomstring"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"price-comparator/model"
)

func init() {
	randomstring.Seed()
	initFirestore(context.Background())
}

func generateProductTestData(t *testing.T) (id string, result map[string]interface{}) {
	result = map[string]interface{}{}
	result["name"] = randomstring.HumanFriendlyString(10)
	result["bio"] = rand.Int()%2 == 1

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

func TestProductDAOFirestoreCreate(t *testing.T) {
	ctx := context.Background()
	productDAO := NewProductDAOFirestore()

	// Upsert new product
	createdProduct := model.NewProduct(
		randomstring.HumanFriendlyString(10),
		rand.Int()%2 == 1,
	)

	err := productDAO.Upsert(ctx, createdProduct)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

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
	cleanupProductTestData(t, createdProduct.ID)
}

func TestProductDAOFirestoreRead(t *testing.T) {
	ctx := context.Background()
	// Setup test data
	id, testData := generateProductTestData(t)

	productDAO := NewProductDAOFirestore()

	loadedProduct, err := productDAO.Load(ctx, id)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, testData["name"], loadedProduct.Name, "Didn't find the right product name")
	assert.Equal(t, testData["bio"], loadedProduct.Bio, "Didn't find the right bio label")
	assert.NotEqual(t, "", loadedProduct.ID, "Loaded store didn't have ID")

	// Cleanup test data
	cleanupStoreTestData(t, id)
}

func TestProductDAOFirestoreUpdate(t *testing.T) {
	ctx := context.Background()
	productDAO := NewProductDAOFirestore()

	// Setup test data
	id, testData := generateProductTestData(t)

	product, err := productDAO.Load(ctx, id)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	product.Name = randomstring.HumanFriendlyString(10)
	product.Bio = !product.Bio
	productDAO.Upsert(ctx, product)

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

func TestProductDAOFirestoreDelete(t *testing.T) {
	ctx := context.Background()
	productDAO := NewProductDAOFirestore()

	// Setup test data
	id, _ := generateProductTestData(t)

	// Delete data
	err := productDAO.Delete(ctx, id)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	_, err = firestoreClient.Collection(firestoreStoreCollection).Doc(id).Get(ctx)
	if grpc.Code(err) != codes.NotFound {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestProductDAOFirestoreSearch(t *testing.T) {

	ctx := context.Background()
	productDAO := NewProductDAOFirestore()

	// Setup test data
	id, testData := generateProductTestData(t)

	searchedByNameProduct := model.Product{
		Name: testData["name"].(string),
	}
	productList, err := productDAO.Search(ctx, &searchedByNameProduct)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if ! assert.Equal(t, 1, len(*productList), "Didn't find the right amount of products") {
		t.Fail()
	}
	assert.Equal(t, testData["name"], (*productList)[0].Name, "Didn't find the right product name")
	assert.Equal(t, testData["bio"], (*productList)[0].Bio, "Didn't find the right product bio label")

	// Cleanup test data
	cleanupStoreTestData(t, id)
}
