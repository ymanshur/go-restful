package repository

import (
	"go-restful/internal/dto"
	"go-restful/internal/model"

	"gorm.io/gorm"
)

type (
	User interface {
		Save(payload *dto.CreateUserRequest) error
		FindById(userId uint) (*model.User, error)
		ExistById(userId uint) (bool, error)
		UpdateById(user *model.User, payload *dto.UpdateUserRequest) (*model.User, error)
		DeleteById(userId uint) error
		FindAll() ([]*model.User, error)
		FindByEmail(email *string) (*model.User, error)
	}
	user struct {
		db *gorm.DB
	}
)

func NewUser(db *gorm.DB) User {
	return &user{
		db,
	}
}

func (u *user) Save(payload *dto.CreateUserRequest) error {
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

func (u *user) FindById(userId uint) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Omit("password").Where(model.User{ID: userId}).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) ExistById(userId uint) (bool, error) {
	var count int64
	if err := u.db.Where(model.User{ID: userId}).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (u *user) UpdateById(user *model.User, payload *dto.UpdateUserRequest) (*model.User, error) {
	updatedUser := new(model.User)
	if err := u.db.Model(&user).Omit("password").Updates(&payload).Take(&updatedUser).Error; err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *user) DeleteById(userId uint) error {
	if err := u.db.Delete(&model.User{ID: userId}).Error; err != nil {
		return err
	}
	return nil
}

func (u *user) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := u.db.Omit("password").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *user) FindByEmail(email *string) (*model.User, error) {
	user := new(model.User)
	if err := u.db.
		Omit("password").Where("email = ?", email).
		Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
