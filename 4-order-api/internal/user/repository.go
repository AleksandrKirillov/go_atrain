package user

import (
	"api/order/pkg/db"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindBySession(sessionId string) (*User, error) {
	var user User
	result := repo.Database.DB.Where("session_id = ?", sessionId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) FindByPhone(phone string) (*User, error) {
	var user User
	result := repo.Database.DB.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) Update(user *User) error {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
