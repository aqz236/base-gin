package service

import (
	"base-gin/internal/domain/user/entity"
	"base-gin/internal/domain/user/repository"
	"errors"
)

type UserDomainService struct {
	userRepo repository.UserRepository
}

func NewUserDomainService(userRepo repository.UserRepository) *UserDomainService {
	return &UserDomainService{
		userRepo: userRepo,
	}
}

// CheckEmailUnique 检查邮箱是否唯一
func (s *UserDomainService) CheckEmailUnique(email string, excludeID int) error {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil // 邮箱不存在，可以使用
	}

	if user.ID != excludeID {
		return errors.New("邮箱已被使用")
	}

	return nil
}

// ValidateUserForCreation 验证用户创建
func (s *UserDomainService) ValidateUserForCreation(name, email, password string) error {
	// 检查邮箱唯一性
	if err := s.CheckEmailUnique(email, 0); err != nil {
		return err
	}

	// 创建用户实体进行验证
	_, err := entity.NewUser(name, email, password)
	return err
}

// ValidateUserForUpdate 验证用户更新
func (s *UserDomainService) ValidateUserForUpdate(id int, name, email string) error {
	// 检查用户是否存在
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 检查邮箱唯一性（排除自己）
	if err := s.CheckEmailUnique(email, id); err != nil {
		return err
	}

	// 验证字段
	if err := user.UpdateName(name); err != nil {
		return err
	}

	if err := user.UpdateEmail(email); err != nil {
		return err
	}

	return nil
}
