package user_impl

import (
	"base-gin/internal/domain/user/entity"
	"base-gin/internal/infrastructure/database"
	"base-gin/internal/infrastructure/database/models"
	"errors"

	"gorm.io/gorm"
)

// GormUserRepository GORM实现的用户仓储
type GormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository 创建新的GORM用户仓储
func NewGormUserRepository(database *database.DB) *GormUserRepository {
	return &GormUserRepository{
		db: database.GetGormDB(),
	}
}

func (r *GormUserRepository) FindByID(id int) (*entity.User, error) {
	var userModel models.UserModel

	if err := r.db.First(&userModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	return userModel.ToEntity(), nil
}

func (r *GormUserRepository) FindByEmail(email string) (*entity.User, error) {
	var userModel models.UserModel

	if err := r.db.Where("email = ?", email).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	return userModel.ToEntity(), nil
}

func (r *GormUserRepository) FindAll() ([]*entity.User, error) {
	var userModels []models.UserModel

	if err := r.db.Find(&userModels).Error; err != nil {
		return nil, err
	}

	users := make([]*entity.User, 0, len(userModels))
	for _, userModel := range userModels {
		users = append(users, userModel.ToEntity())
	}

	return users, nil
}

func (r *GormUserRepository) Save(user *entity.User) error {
	// 检查邮箱是否已存在（只检查未删除的记录）
	var existingUser models.UserModel
	if err := r.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("邮箱已存在")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	userModel := models.NewUserModelFromEntity(user)

	if err := r.db.Create(userModel).Error; err != nil {
		return err
	}

	// 更新实体的ID
	user.ID = int(userModel.ID)
	user.CreatedAt = userModel.CreatedAt
	user.UpdatedAt = userModel.UpdatedAt

	return nil
}

func (r *GormUserRepository) Update(user *entity.User) error {
	userModel := models.NewUserModelFromEntity(user)

	result := r.db.Model(&userModel).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"name":       user.Name,
		"email":      user.Email,
		"password":   user.Password,
		"updated_at": user.UpdatedAt,
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("用户不存在")
	}

	return nil
}

func (r *GormUserRepository) Delete(id int) error {
	// 使用软删除（默认行为）
	result := r.db.Delete(&models.UserModel{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("用户不存在")
	}

	return nil
}

// HardDelete 提供硬删除选项（如果需要的话）
func (r *GormUserRepository) HardDelete(id int) error {
	// 使用 Unscoped() 进行硬删除（真实删除）
	result := r.db.Unscoped().Delete(&models.UserModel{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("用户不存在")
	}

	return nil
}
