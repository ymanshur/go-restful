package repository

import (
	"go-restful/internal/app/dto"
	"go-restful/internal/model"

	"gorm.io/gorm"
)

type (
	User interface {
		Save(payload *dto.CreateUserRequest) error
		FindById(userId uint) (*model.User, error)
		UpdateById(userId uint, user *model.User) (*model.User, error)
		DeleteById(userId uint) error
		FindAll() ([]*model.User, error)
		FindByEmail(email *string) (*model.User, error)
	}
	userRepo struct {
		db *gorm.DB
	}
)

func NewUser(db *gorm.DB) User {
	return &userRepo{
		db,
	}
}

func (u *userRepo) Save(payload *dto.CreateUserRequest) error {
	newUser := model.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}
	if err := u.db.Save(&newUser).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepo) FindById(userId uint) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Omit("password").Where(model.User{ID: userId}).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) UpdateById(userId uint, user *model.User) (*model.User, error) {
	if err := u.db.Model(&model.User{ID: userId}).Updates(&user).Error; err != nil {
		return nil, err
	}

	updatedUser := new(model.User)
	if err := u.db.Omit("password").Find(&updatedUser, userId).Error; err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *userRepo) DeleteById(userId uint) error {
	if err := u.db.Delete(&model.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepo) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := u.db.Omit("password").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepo) FindByEmail(email *string) (*model.User, error) {
	user := new(model.User)
	if err := u.db.
		Omit("password").Where("email = ?", email).
		Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
