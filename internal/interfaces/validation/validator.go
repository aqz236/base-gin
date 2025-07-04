package validation

import (
	"errors"
	"net/mail"
	"strings"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return errors.New("邮箱不能为空")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("邮箱格式不正确")
	}

	return nil
}

func (v *Validator) ValidateName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("姓名不能为空")
	}

	if len(name) < 2 || len(name) > 50 {
		return errors.New("姓名长度必须在2-50个字符之间")
	}

	return nil
}

func (v *Validator) ValidatePassword(password string) error {
	if password == "" {
		return errors.New("密码不能为空")
	}

	if len(password) < 6 {
		return errors.New("密码长度不能少于6位")
	}

	return nil
}
