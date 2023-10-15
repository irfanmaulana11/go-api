package dbrepositories

import (
	"context"
	"go-api/app/models"
)

func (r *DBRepository) CreateUser(ctx context.Context, user models.User) error {
	var err error

	tx := r.db.Begin()
	if err = tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func (r *DBRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var (
		user models.User
		err  error
	)

	tx := r.db.Begin()
	if err = tx.Find(&user, "email = ?", email).Error; err != nil {
		tx.Rollback()
		return user, err
	}
	tx.Commit()

	return user, err
}

func (r *DBRepository) UpdateUser(ctx context.Context, data models.User) error {
	var err error

	tx := r.db.Begin()
	err = tx.Table("users").Omit("id").Where("id = ?", data.ID).Updates(data).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
