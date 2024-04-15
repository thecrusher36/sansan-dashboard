package repository

import (
	"context"
	"testing"

	rolev1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRoleRepository_GetRoleList(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRoleRepository(db)
	ctx := context.TODO()

	// Add test data
	testData := []*rolev1.Role{
		{RoleName: "Admin", CreatedAt: timestamppb.Now()},
		{RoleName: "User", CreatedAt: timestamppb.Now()},
	}

	for _, role := range testData {
		_, err := repo.AddRole(ctx, role)
		assert.NoError(t, err, "Failed to add role")
	}

	// Test GetRoleList
	roles, err := repo.GetRoleList(ctx, nil)
	assert.NoError(t, err, "Failed to get role list")
	assert.NotNil(t, roles, "Returned role list is nil")
	assert.Len(t, roles, len(testData), "Unexpected length of role list")

	// Clean up test data
	for _, role := range roles {
		_, err := repo.RemoveRole(ctx, role)
		assert.NoError(t, err, "Failed to remove role")
	}
}

func TestRoleRepository_GetRole(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRoleRepository(db)
	ctx := context.TODO()

	// Add test data
	testRole := &rolev1.Role{RoleName: "Test Role", CreatedAt: timestamppb.Now()}
	createdRole, err := repo.AddRole(ctx, testRole)
	assert.NoError(t, err, "Failed to add role")

	// Test GetRole
	fetchedRole, err := repo.GetRole(ctx, &rolev1.Role{Id: createdRole.Id})
	assert.NoError(t, err, "Failed to get role")
	assert.NotNil(t, fetchedRole, "Returned role is nil")
	assert.Equal(t, createdRole.Id, fetchedRole.Id, "Unexpected role ID")

	// Clean up test data
	_, err = repo.RemoveRole(ctx, createdRole)
	assert.NoError(t, err, "Failed to remove role")
}

func TestRoleRepository_AddRole(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRoleRepository(db)
	ctx := context.TODO()

	// Test AddRole
	testRole := &rolev1.Role{RoleName: "Test Role", CreatedAt: timestamppb.Now()}
	createdRole, err := repo.AddRole(ctx, testRole)
	assert.NoError(t, err, "Failed to add role")
	assert.NotNil(t, createdRole, "Returned role is nil")

	// Clean up test data
	_, err = repo.RemoveRole(ctx, createdRole)
	assert.NoError(t, err, "Failed to remove role")
}

func TestRoleRepository_EditRole(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRoleRepository(db)
	ctx := context.TODO()

	// Add test data
	testRole := &rolev1.Role{RoleName: "Test Role", CreatedAt: timestamppb.Now()}
	createdRole, err := repo.AddRole(ctx, testRole)
	assert.NoError(t, err, "Failed to add role")

	// Modify test data
	createdRole.RoleName = "Modified Role"

	// Test EditRole
	modifiedRole, err := repo.EditRole(ctx, createdRole)
	assert.NoError(t, err, "Failed to edit role")
	assert.NotNil(t, modifiedRole, "Returned modified role is nil")
	assert.Equal(t, createdRole.Id, modifiedRole.Id, "Unexpected role ID")
	assert.Equal(t, createdRole.RoleName, modifiedRole.RoleName, "Unexpected role name")

	// Clean up test data
	_, err = repo.RemoveRole(ctx, modifiedRole)
	assert.NoError(t, err, "Failed to remove role")
}

func TestRoleRepository_RemoveRole(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRoleRepository(db)
	ctx := context.TODO()

	// Add test data
	testRole := &rolev1.Role{RoleName: "Test Role", CreatedAt: timestamppb.Now()}
	createdRole, err := repo.AddRole(ctx, testRole)
	assert.NoError(t, err, "Failed to add role")

	// Test RemoveRole
	removedRole, err := repo.RemoveRole(ctx, createdRole)
	assert.NoError(t, err, "Failed to remove role")
	assert.NotNil(t, removedRole, "Returned removed role is nil")

	// Try fetching removed role
	_, err = repo.GetRole(ctx, removedRole)
	assert.Error(t, err, "Error expected when fetching removed role")
}
