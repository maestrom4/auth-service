package mongodb_test

// Currently fixing this
// import (
// 	"context"
// 	"testing"

// 	"auth-service/internal/models"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // Mock objects and methods ...
// type MockCollection struct {
// 	mock.Mock
// }

// type MockSingleResult struct {
// 	mock.Mock
// }

// func (m *MockSingleResult) Decode(v interface{}) error {
// 	args := m.Called(v)
// 	return args.Error(0)
// }

// func (m *MockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
// 	args := m.Called(ctx, filter, opts)
// 	return args.Get(0).(*mongo.SingleResult)
// }

// var objectID primitive.ObjectID
// var err error

// objectID, err = primitive.ObjectIDFromHex(id)
// if err != nil {
//     // Handle the error
//     fmt.Println("Error:", err)
//     return
// }

// // Use the objectID
// fmt.Println("ObjectID:", objectID)

// func TestGetUserByID(t *testing.T) {
// 	// Mock the collection

// 	collection := new(MockCollection)
// 	ctx := context.TODO()
// 	id := primitive.NewObjectID().Hex()

// 	// Set up the expected user
// 	expectedUser := &models.User{
// 		ID:             id,
// 		Username:       "",
// 		HashedPassword: "",
// 		CreatedAt:      "",
// 		LastLogin:      "",
// 	}

// 	singleResult := new(MockSingleResult)
// 	collection.On("FindOne", ctx, bson.M{"_id": primitive.ObjectIDFromHex(id)}).Return(singleResult)

// 	singleResult.On("Decode", mock.AnythingOfType("*models.User")).Run(func(args mock.Arguments) {
// 		arg := args.Get(0).(*models.User)
// 		*arg = *expectedUser
// 	}).Return(nil)

// 	repo := UserRepository{Collection: collection}
// 	user, err := repo.GetUserByID(ctx, id)

// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedUser, user)
// }
