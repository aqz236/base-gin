package repository

import "base-gin/internal/domain/user/entity"

type UserRepository interface {
	FindByID(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	Save(user *entity.User) error
	Update(user *entity.User) error
	Delete(id int) error
}
