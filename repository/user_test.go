package repository

import (
	"context"
	"testing"

	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestUserRepository_GetUserList(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.TODO()

	// Add test data
	testData := []*userv1.User{
		{UserName: "user1", CreatedAt: timestamppb.Now()},
		{UserName: "user2", CreatedAt: timestamppb.Now()},
	}

	for _, user := range testData {
		_, err := repo.AddUser(ctx, user)
		assert.NoError(t, err, "AddUser should not return an error")
	}

	// Test GetUserList
	users, err := repo.GetUserList(ctx, nil)
	assert.NoError(t, err, "GetUserList should not return an error")
	assert.NotNil(t, users, "GetUserList should return a list of users")
	assert.Len(t, users, len(testData), "GetUserList should return the correct number of users")

	// Clean up test data
	for _, user := range users {
		_, err := repo.RemoveUser(ctx, user)
		assert.NoError(t, err, "RemoveUser should not return an error")
	}
}

func TestUserRepository_GetUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.TODO()

	// Add test data
	testUser := &userv1.User{UserName: "testuser", CreatedAt: timestamppb.Now()}
	createdUser, err := repo.AddUser(ctx, testUser)
	assert.NoError(t, err, "AddUser should not return an error")

	// Test GetUser
	fetchedUser, err := repo.GetUser(ctx, &userv1.User{Id: createdUser.Id})
	assert.NoError(t, err, "GetUser should not return an error")
	assert.NotNil(t, fetchedUser, "GetUser should return a user")
	assert.Equal(t, createdUser.Id, fetchedUser.Id, "GetUser should return the correct user")

	// Clean up test data
	_, err = repo.RemoveUser(ctx, fetchedUser)
	assert.NoError(t, err, "RemoveUser should not return an error")
}

func TestUserRepository_AddUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.TODO()

	// Test AddUser
	testUser := &userv1.User{UserName: "testuser", CreatedAt: timestamppb.Now()}
	createdUser, err := repo.AddUser(ctx, testUser)
	assert.NoError(t, err, "AddUser should not return an error")
	assert.NotNil(t, createdUser, "Returned user is nil")

	// Clean up test data
	_, err = repo.RemoveUser(ctx, createdUser)
	assert.NoError(t, err, "RemoveUser should not return an error")
}

func TestUserRepository_EditUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.TODO()

	// Add test data
	testUser := &userv1.User{UserName: "testuser", CreatedAt: timestamppb.Now()}
	createdUser, err := repo.AddUser(ctx, testUser)
	assert.NoError(t, err, "AddUser should not return an error")

	// Modify test data
	createdUser.UserName = "modifieduser"

	// Test EditUser
	modifiedUser, err := repo.EditUser(ctx, createdUser)
	assert.NoError(t, err, "EditUser should not return an error")
	assert.NotNil(t, modifiedUser, "Returned user is nil")
	assert.Equal(t, createdUser.Id, modifiedUser.Id, "Returned user is nil")
	assert.Equal(t, createdUser.UserName, modifiedUser.UserName, "Returned user is nil")

	// Clean up test data
	_, err = repo.RemoveUser(ctx, modifiedUser)
	assert.NoError(t, err, "RemoveUser should not return an error")
}

func TestUserRepository_RemoveUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.TODO()

	// Add test data
	testUser := &userv1.User{UserName: "testuser", CreatedAt: timestamppb.Now()}
	createdUser, err := repo.AddUser(ctx, testUser)
	assert.NoError(t, err, "AddUser should not return an error")

	// Test RemoveUser
	removedUser, err := repo.RemoveUser(ctx, createdUser)
	assert.NoError(t, err, "RemoveUser should not return an error")
	assert.NotNil(t, removedUser, "Returned user is nil")

	// Try fetching removed user
	_, err = repo.GetUser(ctx, removedUser)
	assert.Error(t, err, "GetUser should return an error") // User should not be found
}
