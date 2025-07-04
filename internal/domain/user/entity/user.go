package entity

import (
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name, email, password string) (*User, error) {
	user := &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("用户名不能为空")
	}

	if len(u.Name) < 2 || len(u.Name) > 50 {
		return errors.New("用户名长度必须在2-50个字符之间")
	}

	if u.Email == "" {
		return errors.New("邮箱不能为空")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("邮箱格式不正确")
	}

	if u.Password == "" {
		return errors.New("密码不能为空")
	}

	if len(u.Password) < 6 {
		return errors.New("密码长度不能少于6位")
	}

	return nil
}

func (u *User) UpdateName(name string) error {
	if name == "" {
		return errors.New("用户名不能为空")
	}
	if len(name) < 2 || len(name) > 50 {
		return errors.New("用户名长度必须在2-50个字符之间")
	}

	u.Name = name
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) UpdateEmail(email string) error {
	if email == "" {
		return errors.New("邮箱不能为空")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("邮箱格式不正确")
	}

	u.Email = email
	u.UpdatedAt = time.Now()
	return nil
}
