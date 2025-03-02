package authentication

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

func (r *Repository) GetByEmail(email string) (*User, error) {
	var user User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
