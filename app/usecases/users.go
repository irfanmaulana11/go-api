package usecases

import (
	"errors"
	"fmt"
	"go-api/app/models"
	"go-api/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func (u *APIGOUsecase) UserRegister(c *gin.Context, user models.User) (*helpers.ValidationErrors, error) {
	userRules := &models.UserRules{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	valid, err := govalidator.ValidateStruct(userRules)
	if !valid && err != nil {
		errors := helpers.ValidationError(userRules, err)
		return errors, err
	}

	hashPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.IsActive = true
	user.Password = hashPassword

	if err := u.DBRepository.CreateUser(c, user); err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *APIGOUsecase) UserLogin(c *gin.Context, user models.UserLogin) (*string, error) {
	userData, err := u.DBRepository.GetUserByEmail(c, user.Email)
	if err != nil {
		return nil, err
	}

	valid := helpers.CheckPassword(user.Password, userData.Password)
	if !valid {
		return nil, errors.New("user name and password invalid")
	}

	token, err := helpers.CreateToken(userData.Email)
	if err != nil {
		return nil, errors.New("failed generating token")
	}

	return &token, nil
}

func (u *APIGOUsecase) GetUserDetail(c *gin.Context, email string) (*models.User, error) {
	userData, err := u.DBRepository.GetUserByEmail(c, email)
	if err != nil {
		return nil, err
	}

	return &userData, nil
}

func (u *APIGOUsecase) UserUpdate(c *gin.Context, email string, user models.User) (*helpers.ValidationErrors, error) {
	/*
		changing email data will affect the login process,
		re-login with the latest email
	*/

	userRules := &models.UserRules{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	userData, err := u.DBRepository.GetUserByEmail(c, email)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Can't find data : %s", err.Error()))
	}

	valid, err := govalidator.ValidateStruct(userRules)
	if !valid && err != nil {
		errors := helpers.ValidationError(userRules, err)
		return errors, err
	}

	hashPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	userData.Email = user.Email
	userData.Name = user.Name
	userData.Password = hashPassword
	user.UpdatedAt = time.Now()

	if err := u.DBRepository.UpdateUser(c, userData); err != nil {
		return nil, err
	}

	return nil, nil
}
