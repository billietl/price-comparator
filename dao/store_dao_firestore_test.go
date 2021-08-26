package dao

import (
	"context"
	"price-comparator/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyproto/randomstring"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func init() {
	randomstring.Seed()
	initFirestore(context.Background())
}

func generateStoreTestData(t *testing.T) (store *model.Store) {
	store = model.GenerateRandomStore()
	dao := NewStoreDAOFirestore()

	ctx := context.Background()
	doc, _, err := firestoreClient.Collection(firestoreStoreCollection).Add(ctx, dao.fromModel(store))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	store.ID = doc.ID

	return
}

func TestStoreDAOFirestoreCreate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	storeDAO := NewStoreDAOFirestore()

	// Upsert new store
	createdStore := model.NewStore(
		randomstring.HumanFriendlyString(10),
		randomstring.HumanFriendlyString(10),
		randomstring.HumanFriendlyString(5),
	)

	err := storeDAO.Upsert(ctx, createdStore)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	defer firestoreClient.Collection(firestoreStoreCollection).Doc(createdStore.ID).Delete(ctx)

	// Reload data
	doc, err := firestoreClient.Collection(firestoreStoreCollection).Doc(createdStore.ID).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	assert.Equal(t, createdStore.Name, docData["name"])
	assert.Equal(t, createdStore.City, docData["city"])
	assert.Equal(t, createdStore.Zipcode, docData["zipcode"])
}

func TestStoreDAOFirestoreRead(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	// Setup test data
	testStore := generateStoreTestData(t)
	defer firestoreClient.Collection(firestoreStoreCollection).Doc(testStore.ID).Delete(ctx)

	storeDAO := NewStoreDAOFirestore()

	loadedStore, err := storeDAO.Load(ctx, testStore.ID)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, testStore.Name, loadedStore.Name, "Didn't find the right store name")
	assert.Equal(t, testStore.City, loadedStore.City, "Didn't find the right store city")
	assert.Equal(t, testStore.Zipcode, loadedStore.Zipcode, "Didn't find the right store zipcode")
	assert.NotEqual(t, "", loadedStore.ID, "Loaded store didn't have ID")
}

func TestStoreDAOFirestoreUpdate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	storeDAO := NewStoreDAOFirestore()

	// Setup test data
	testStore := generateStoreTestData(t)
	defer firestoreClient.Collection(firestoreStoreCollection).Doc(testStore.ID).Delete(ctx)

	store, err := storeDAO.Load(ctx, testStore.ID)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	store.Name = randomstring.HumanFriendlyString(10)
	store.City = randomstring.HumanFriendlyString(10)
	store.Zipcode = randomstring.HumanFriendlyString(5)
	storeDAO.Upsert(ctx, store)

	// Reload data
	doc, err := firestoreClient.Collection(firestoreStoreCollection).Doc(testStore.ID).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	assert.NotEqual(t, testStore.Name, docData["name"])
	assert.NotEqual(t, testStore.City, docData["city"])
	assert.NotEqual(t, testStore.Zipcode, docData["zipcode"])
}

func TestStoreDAOFirestoreDelete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	storeDAO := NewStoreDAOFirestore()

	// Setup test data
	testStore := generateStoreTestData(t)

	err := storeDAO.Delete(ctx, testStore.ID)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	_, err = firestoreClient.Collection(firestoreStoreCollection).Doc(testStore.ID).Get(ctx)
	if grpc.Code(err) != codes.NotFound {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestStoreDAOFirestoreSearch(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	storeDAO := NewStoreDAOFirestore()

	// Setup test data
	testStore := generateStoreTestData(t)
	defer firestoreClient.Collection(firestoreStoreCollection).Doc(testStore.ID).Delete(ctx)

	searchedByNameStore := model.Store{
		Name: testStore.Name,
	}
	storeList, err := storeDAO.Search(ctx, &searchedByNameStore)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, 1, len(*storeList), "Didn't find the right amount of stores")
	assert.Equal(t, testStore.Name, (*storeList)[0].Name, "Didn't find the right store name")
	assert.Equal(t, testStore.City, (*storeList)[0].City, "Didn't find the right store city")
	assert.Equal(t, testStore.Zipcode, (*storeList)[0].Zipcode, "Didn't find the right store zipcode")
}
