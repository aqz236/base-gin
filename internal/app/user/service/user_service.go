package service

import (
	"base-gin/internal/domain/user/entity"
	"base-gin/internal/domain/user/repository"
	domainService "base-gin/internal/domain/user/service"
	"base-gin/internal/domain/user/vo"
)

type UserService struct {
	userRepo          repository.UserRepository
	userDomainService *domainService.UserDomainService
}

func NewUserService(userRepo repository.UserRepository, userDomainService *domainService.UserDomainService) *UserService {
	return &UserService{
		userRepo:          userRepo,
		userDomainService: userDomainService,
	}
}

func (s *UserService) GetUser(id int) (*vo.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &vo.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserService) GetAllUsers() ([]*vo.UserResponse, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]*vo.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, &vo.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return responses, nil
}

func (s *UserService) CreateUser(req *vo.UserCreateRequest) (*vo.UserResponse, error) {
	// 使用领域服务验证
	if err := s.userDomainService.ValidateUserForCreation(req.Name, req.Email, req.Password); err != nil {
		return nil, err
	}

	// 创建用户实体
	user, err := entity.NewUser(req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	// 保存用户
	if err := s.userRepo.Save(user); err != nil {
		return nil, err
	}

	return &vo.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserService) UpdateUser(id int, req *vo.UserUpdateRequest) (*vo.UserResponse, error) {
	// 使用领域服务验证
	if err := s.userDomainService.ValidateUserForUpdate(id, req.Name, req.Email); err != nil {
		return nil, err
	}

	// 获取用户
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// 更新用户信息
	if err := user.UpdateName(req.Name); err != nil {
		return nil, err
	}

	if err := user.UpdateEmail(req.Email); err != nil {
		return nil, err
	}

	// 保存更新
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return &vo.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.Delete(id)
}
