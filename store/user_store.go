package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/onattech/invest/models"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) models.UserStore {
	return &userStore{
		db: db,
	}
}

func (us *userStore) Create(c context.Context, user *models.User) error {
	// Insert the user into the database using GORM's Create method
	return us.db.WithContext(c).Create(user).Error
}

func (us *userStore) Fetch(c context.Context) ([]models.User, error) {
	var users []models.User
	// Exclude the "password" field from the results
	err := us.db.WithContext(c).Select("id", "name", "email").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *userStore) GetByEmail(c context.Context, email string) (models.User, error) {
	var user models.User
	// Find a user by email
	err := us.db.WithContext(c).Where("email = ?", email).First(&user).Error
	return user, err
}

func (us *userStore) GetByID(c context.Context, id uuid.UUID) (models.User, error) {
	var user models.User
	// Find a user by ID
	err := us.db.WithContext(c).First(&user, "id = ?", id).Error
	return user, err
}
