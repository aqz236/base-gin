package user_impl

import (
	"base-gin/internal/domain/user/entity"
	"base-gin/internal/infrastructure/database"
	"errors"
	"sync"
	"time"
)

type MockUserRepository struct {
	db     *database.DB
	users  map[int]*entity.User
	mutex  sync.RWMutex
	nextID int
}

func NewMockUserRepository(db *database.DB) *MockUserRepository {
	repo := &MockUserRepository{
		db:     db,
		users:  make(map[int]*entity.User),
		nextID: 1,
	}

	// 初始化一些Mock数据
	repo.initMockData()

	return repo
}

func (r *MockUserRepository) initMockData() {
	users := []*entity.User{
		{
			ID:        1,
			Name:      "张三",
			Email:     "zhangsan@example.com",
			Password:  "password123",
			CreatedAt: time.Now().Add(-24 * time.Hour),
			UpdatedAt: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:        2,
			Name:      "李四",
			Email:     "lisi@example.com",
			Password:  "password456",
			CreatedAt: time.Now().Add(-12 * time.Hour),
			UpdatedAt: time.Now().Add(-12 * time.Hour),
		},
		{
			ID:        3,
			Name:      "王五",
			Email:     "wangwu@example.com",
			Password:  "password789",
			CreatedAt: time.Now().Add(-6 * time.Hour),
			UpdatedAt: time.Now().Add(-6 * time.Hour),
		},
	}

	for _, user := range users {
		r.users[user.ID] = user
		if user.ID >= r.nextID {
			r.nextID = user.ID + 1
		}
	}
}

func (r *MockUserRepository) FindByID(id int) (*entity.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("用户不存在")
	}

	// 返回副本以避免外部修改
	userCopy := *user
	return &userCopy, nil
}

func (r *MockUserRepository) FindByEmail(email string) (*entity.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			// 返回副本以避免外部修改
			userCopy := *user
			return &userCopy, nil
		}
	}

	return nil, errors.New("用户不存在")
}

func (r *MockUserRepository) FindAll() ([]*entity.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]*entity.User, 0, len(r.users))
	for _, user := range r.users {
		// 返回副本以避免外部修改
		userCopy := *user
		users = append(users, &userCopy)
	}

	return users, nil
}

func (r *MockUserRepository) Save(user *entity.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	user.ID = r.nextID
	r.nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	r.users[user.ID] = user
	return nil
}

func (r *MockUserRepository) Update(user *entity.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("用户不存在")
	}

	user.UpdatedAt = time.Now()
	r.users[user.ID] = user
	return nil
}

func (r *MockUserRepository) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("用户不存在")
	}

	delete(r.users, id)
	return nil
}
