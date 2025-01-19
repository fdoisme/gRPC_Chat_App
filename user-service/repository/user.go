package repository

import (
	"errors"
	"log"
	"project/user-service/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(filter bool) ([]model.User, error)
	Insert(user *model.User) error
	UpdateProfile(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllUsers(filter bool) ([]model.User, error) {
	var users []model.User
	if filter {
		err := r.db.Where("is_online=?", filter).Find(&users).Error
		if err != nil {
			log.Println(err)
			return []model.User{}, errors.New("Internal Server Error")
		}
	} else {
		if err := r.db.Find(&users).Error; err != nil {
			log.Println(err)
			return []model.User{}, errors.New("Internal Server Error")
		}
	}
	return users, nil
}
func (repo *userRepository) Insert(user *model.User) error {
	err := repo.db.Create(&user).Error
	if err != nil {
		log.Println(err)
		return errors.New("Already Registered")
	}
	return nil
}
func (repo *userRepository) UpdateProfile(user *model.User) error {
	err := repo.db.Model(&model.User{}).
		Where("email ILIKE ?", user.Email).
		Updates(user).Error
	if err != nil {
		log.Println(err)
		return errors.New("Bad Request")
	}
	return nil
}
