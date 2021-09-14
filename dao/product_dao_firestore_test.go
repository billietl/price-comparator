package dao

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyproto/randomstring"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"price-comparator/model"
	"price-comparator/testUtils"
)

func init() {
	randomstring.Seed()
	initFirestore(context.Background())
}

func generateProductTestData(t *testing.T) (product *model.Product) {
	product = testUtils.GenerateRandomProduct()
	dao := NewProductDAOFirestore()

	ctx := context.Background()
	doc, _, err := firestoreClient.Collection(firestoreProductCollection).Add(ctx, dao.fromModel(product))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	product.ID = doc.ID

	return
}

func TestProductDAOFirestoreCreate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	productDAO := NewProductDAOFirestore()

	// Upsert new product
	createdProduct := testUtils.GenerateRandomProduct()

	err := productDAO.Upsert(ctx, createdProduct)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	defer firestoreClient.Collection(firestoreProductCollection).Doc(createdProduct.ID).Delete(ctx)

	// Reload product
	doc, err := firestoreClient.Collection(firestoreProductCollection).Doc(createdProduct.ID).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	assert.Equal(t, createdProduct.Name, docData["name"])
	assert.Equal(t, createdProduct.Bio, docData["bio"])
	assert.Equal(t, createdProduct.Vrac, docData["vrac"])
}

func TestProductDAOFirestoreRead(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	// Setup test data
	testProduct := generateProductTestData(t)
	defer firestoreClient.Collection(firestoreProductCollection).Doc(testProduct.ID).Delete(ctx)

	productDAO := NewProductDAOFirestore()

	loadedProduct, err := productDAO.Load(ctx, testProduct.ID)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, testProduct.Name, loadedProduct.Name, "Didn't find the right product name")
	assert.Equal(t, testProduct.Bio, loadedProduct.Bio, "Didn't find the right bio label")
	assert.Equal(t, testProduct.Vrac, loadedProduct.Vrac, "Didn't find the right vrac label")
	assert.NotEqual(t, "", loadedProduct.ID, "Loaded store didn't have ID")
}

func TestProductDAOFirestoreUpdate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	productDAO := NewProductDAOFirestore()

	// Setup test data
	testProduct := generateProductTestData(t)
	defer firestoreClient.Collection(firestoreProductCollection).Doc(testProduct.ID).Delete(ctx)

	product, err := productDAO.Load(ctx, testProduct.ID)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	product.Name = randomstring.HumanFriendlyString(10)
	product.Bio = !product.Bio
	product.Vrac = !product.Vrac
	productDAO.Upsert(ctx, product)

	// Reload data
	doc, err := firestoreClient.Collection(firestoreProductCollection).Doc(testProduct.ID).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	assert.NotEqual(t, docData["name"], testProduct.Name)
	assert.NotEqual(t, docData["bio"], testProduct.Bio)
	assert.NotEqual(t, docData["vrac"], testProduct.Vrac)
}

func TestProductDAOFirestoreDelete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	productDAO := NewProductDAOFirestore()

	// Setup test data
	testProduct := generateProductTestData(t)

	// Delete data
	err := productDAO.Delete(ctx, testProduct.ID)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	_, err = firestoreClient.Collection(firestoreProductCollection).Doc(testProduct.ID).Get(ctx)
	if grpc.Code(err) != codes.NotFound {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestProductDAOFirestoreSearch(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	productDAO := NewProductDAOFirestore()

	// Setup test data
	testProduct := generateProductTestData(t)
	defer firestoreClient.Collection(firestoreProductCollection).Doc(testProduct.ID).Delete(ctx)

	searchedByNameProduct := model.Product{
		Name: testProduct.Name,
	}
	productList, err := productDAO.Search(ctx, &searchedByNameProduct, false, false)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if !assert.Equal(t, 1, len(*productList), "Didn't find the right amount of products") {
		t.Fail()
		return
	}
	assert.Equal(t, testProduct.Name, (*productList)[0].Name, "Didn't find the right product name")
	assert.Equal(t, testProduct.Bio, (*productList)[0].Bio, "Didn't find the right product bio label")
	assert.Equal(t, testProduct.Vrac, (*productList)[0].Vrac, "Didn't find the right product vrac label")
}
