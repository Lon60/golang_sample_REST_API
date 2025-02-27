package user

import "gorm.io/gorm"

// Repository wraps a gorm.DB for User operations.
type Repository struct {
	DB *gorm.DB
}

// NewRepository creates a new instance of User repository.
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

// CreateUser inserts a new user into the database.
func (r *Repository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

// GetByEmail retrieves a user by email.
func (r *Repository) GetByEmail(email string) (*User, error) {
	var user User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
} 