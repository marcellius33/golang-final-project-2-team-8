package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(userId uint) (*models.User, error)
	UpdateUser(userId uint, user *models.User) (*models.User, error)
	DeleteUser(userId uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	return user, r.db.Create(user).Error
}

func (r *userRepository) FindUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindUserByID(userId uint) (*models.User, error) {
	user := models.User{}
	err := r.db.Where("id = ?", userId).First(&user).Error
	return &user, err
}

func (r *userRepository) UpdateUser(userId uint, updateUser *models.User) (*models.User, error) {
	user := updateUser
	err := r.db.Model(&user).Where("id=?", userId).Updates(updateUser).Error
	return user, err
}

func (r *userRepository) DeleteUser(userId uint) error {
	var order models.User

	err := r.db.Where("id=?", userId).Delete(&order).Error
	return err
}