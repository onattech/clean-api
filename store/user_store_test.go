package store_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/onattech/invest/models"
	"github.com/onattech/invest/store"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserStore(t *testing.T) {
	ctx := context.Background()

	// Start a PostgreSQL container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:17.0-alpine3.20",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	pgContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start PostgreSQL container: %v", err)
	}
	defer pgContainer.Terminate(ctx)

	// Get the container's host and port
	host, err := pgContainer.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %v", err)
	}
	port, err := pgContainer.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("Failed to get container port: %v", err)
	}

	// Build the DSN
	dsn := fmt.Sprintf("host=%s user=testuser password=testpass dbname=testdb port=%s sslmode=disable", host, port.Port())

	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	us := store.NewUserStore(db)

	t.Run("Create User", func(t *testing.T) {
		mockUser := &models.User{
			Name:     "Test",
			Email:    "test@gmail.com",
			Password: "password",
		}

		err := us.Create(context.Background(), mockUser)
		assert.NoError(t, err)

		// Verify that the user was created
		var user models.User
		err = db.First(&user, "email = ?", mockUser.Email).Error
		assert.NoError(t, err)
		assert.Equal(t, mockUser.Name, user.Name)
		assert.Equal(t, mockUser.Email, user.Email)
		log.Printf("user: %+v\n", user)

	})

	t.Run("Get User by Email", func(t *testing.T) {
		email := "test@gmail.com"
		user, err := us.GetByEmail(context.Background(), email)
		assert.NoError(t, err)
		assert.Equal(t, email, user.Email)
	})

	t.Run("Get User by ID", func(t *testing.T) {
		// Create a new user
		mockUser := &models.User{
			Name:     "Another Test",
			Email:    "another@test.com",
			Password: "password",
		}
		err := us.Create(context.Background(), mockUser)
		assert.NoError(t, err)

		user, err := us.GetByID(context.Background(), mockUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.ID, user.ID)
	})
}
