package user_impl

import (
	"base-gin/internal/domain/user/entity"
	"base-gin/internal/infrastructure/database"
	"errors"
	"sync"
	"time"
)

type UserRepo struct {
	db     *database.DB
	users  map[int]*entity.User
	mutex  sync.RWMutex
	nextID int
}

func NewUserRepository(db *database.DB) *UserRepo {
	repo := &UserRepo{
		db:     db,
		users:  make(map[int]*entity.User),
		nextID: 1,
	}

	// 初始化一些Mock数据
	repo.initMockData()

	return repo
}

func (r *UserRepo) initMockData() {
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

func (r *UserRepo) FindByID(id int) (*entity.User, error) {
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

func (r *UserRepo) FindByEmail(email string) (*entity.User, error) {
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

func (r *UserRepo) FindAll() ([]*entity.User, error) {
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

func (r *UserRepo) Save(user *entity.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	user.ID = r.nextID
	r.nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	r.users[user.ID] = user
	return nil
}

func (r *UserRepo) Update(user *entity.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("用户不存在")
	}

	user.UpdatedAt = time.Now()
	r.users[user.ID] = user
	return nil
}

func (r *UserRepo) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("用户不存在")
	}

	delete(r.users, id)
	return nil
}
