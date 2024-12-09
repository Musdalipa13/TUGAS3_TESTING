package repository

import (
	"context"
	"fmt"
	"golang-database-user/config"
	"golang-database-user/model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser_Success(t *testing.T) {
	// Membuka koneksi database
	dbConn, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	// Inisialisasi repository
	userRepo := NewUserRepositoryImpl(dbConn)
	roleRepo := NewRoleRepositoryImpl(dbConn)

	ctx := context.Background()

	// Mencari role berdasarkan kode
	role, err := roleRepo.FindMstRole(ctx, "ROLE001")
	if err != nil {
		t.Fatalf("Failed to find role: %v", err)
	}

	// Data user baru
	newUser := model.MstUser{
		IdUser:      uuid.NewString(),
		Name:        "Test ipa",
		Email:       "test.ipa@example.com",
		Password:    "musdalipa",
		PhoneNumber: "083404879779",
		Role:        role,
	}

	// Menyisipkan user baru
	createdUser, err := userRepo.InsertUser(ctx, newUser)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	// Assertion untuk memastikan data sesuai
	assert.NotNil(t, createdUser, "Created user should not be nil")
	assert.Equal(t, newUser.IdUser, createdUser.IdUser, "User IDs should match")
	assert.Equal(t, newUser.Email, createdUser.Email, "Emails should match")
	assert.Equal(t, newUser.Role, createdUser.Role, "Roles should match")
	assert.Equal(t, newUser.Name, createdUser.Name, "Names should match")
}

func TestInsertUser_Fail(t *testing.T) {
	// Membuka koneksi database
	dbConn, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	// Inisialisasi repository
	userRepo := NewUserRepositoryImpl(dbConn)
	// roleRepo := NewRoleRepositoryImpl(dbConn)

	ctx := context.Background()

	// Data user baru dengan nilai properti tidak valid (contoh: email kosong)
	newUser := model.MstUser{}

	// Menyisipkan user baru
	_, err = userRepo.InsertUser(ctx, newUser)

	// Assertion untuk memastikan kegagalan
	assert.Error(t, err, "An error should occur when inserting a user with invalid data")
	assert.Contains(t, err.Error(), "email", "Error message should mention the invalid email")
}

func TestUpdateUser_Success(t *testing.T) {
	// Membuka koneksi database
	dbConn, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	// Inisialisasi repository
	userRepo := NewUserRepositoryImpl(dbConn)


	ctx := context.Background()

	// UserId yang ingin diperbarui
	userId := "4d4d68ff-8edf-41a0-9e70-7c0cbe92f4b2"
	user := model.MstUser{
		Name:       "ipa",
		Email:      "ipa@gmail.com",
		Password:   "musdalipa",
		PhoneNumber: "083404879779",
	}

	// Memperbarui user
	updatedUser, err := userRepo.UpdateUser(ctx, user, userId)


	// Assertion untuk memastikan data yang diperbarui sesuai dengan yang diinginkan
	assert.NotNil(t, updatedUser, "Updated user should not be nil")
	assert.Nil(t, err)
}

func TestUpdateUser_Fail(t *testing.T) {
    // Membuka koneksi database
    dbConn, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        t.Fatalf("Failed to connect to the database: %v", err)
    }
    defer dbConn.Close()

    // Inisialisasi repository
    userRepo := NewUserRepositoryImpl(dbConn)

    ctx := context.Background()

    // UserId yang tidak valid (kosong)
    invalidUserId := ""
    user := model.MstUser{
        Name:       "ipa",
        Email:      "ipa@gmail.com",
        Password:   "musdalipa",
        PhoneNumber: "083404879779",
    }

    // Mencoba memperbarui user dengan ID yang tidak valid
    _, err = userRepo.UpdateUser(ctx, user, invalidUserId)

    assert.Error(t, err, "An error should occur when trying to update a user with an invalid ID")
	assert.Contains(t, err.Error(), "invalid user ID", "Error message should mention the invalid ID")
	
}

func TestReadUser_Success(t *testing.T) {
    dbConn, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        t.Fatalf("Failed to connect to the database: %v", err)
    }
    defer dbConn.Close()

    // Inisialisasi repository
    userRepo := NewUserRepositoryImpl(dbConn)

    ctx := context.Background()

    // Mengambil data user
    user, err := userRepo.ReadUser(ctx)

    // Assertion untuk memastikan tidak ada error
    assert.Nil(t, err, "Error should be nil when reading user")

    // Pastikan data yang dibaca tidak kosong
    assert.NotNil(t, user, "User should not be nil")

}

func TestReadUser_Fail(t *testing.T) {
    // Simulasi kegagalan membuka koneksi database (misalnya koneksi gagal)
    dbConn, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        t.Fatalf("Failed to connect to the database: %v", err)
    }
    defer dbConn.Close()

    userRepo := NewUserRepositoryImpl(dbConn)

    ctx := context.Background()

    // Mencoba untuk membaca user yang tidak ada di database
    users, err := userRepo.ReadUser(ctx)

    // Memastikan ada error yang terjadi
    assert.NotNil(t, err, "An error should occur when trying to read a user with invalid data or query failure")

    // Memastikan error yang terjadi sesuai dengan yang diharapkan
    assert.Contains(t, err.Error(), "no user found", "Error message should mention 'no user found' or relevant error")
    assert.Empty(t, users, "Users should be empty when no records are found")
}

func TestDeleteUser_Success(t *testing.T) {
	dbConn, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        t.Fatalf("Failed to connect to the database: %v", err)
    }
    defer dbConn.Close()

    // Inisialisasi repository
    userRepo := NewUserRepositoryImpl(dbConn)

    ctx := context.Background()

	userId := "4d4d68ff-8edf-41a0-9e70-7c0cbe92f4b2"

	if userId == "" {
		fmt.Errorf("ID user tidak boleh kosong %v", userId)
	}

	err = userRepo.DeleteUser(ctx, userId)
	
	assert.Nil(t, err)
	
}

func TestDeleteUser_Fail(t *testing.T) {
	dbConn, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        t.Fatalf("Failed to connect to the database: %v", err)
    }
    defer dbConn.Close()

    // Inisialisasi repository
    userRepo := NewUserRepositoryImpl(dbConn)

    ctx := context.Background()

	userId := ""

	if userId == "" {
		fmt.Errorf("ID user tidak boleh kosong %v", userId)
	}

	err = userRepo.DeleteUser(ctx, userId)
	
	assert.NotNil(t, err)
	
}














