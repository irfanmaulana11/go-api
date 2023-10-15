package handlers

import (
	"errors"
	"go-api/app/models"
	"go-api/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *APIGOHandler) UserRegister(c *gin.Context) {
	reqBody := models.User{}

	if err := c.BindJSON(&reqBody); err != nil {
		if err != nil {
			helpers.RespondErrorJSON(c, http.StatusInternalServerError, errors.New("can't bind request to struct"))
			return
		}
	}

	if valErrors, err := h.APIGOUsecase.UserRegister(c, reqBody); err != nil {
		if valErrors != nil {
			c.JSON(http.StatusUnprocessableEntity, valErrors)
			return
		}

		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	helpers.RespondJSON(c, http.StatusCreated, "User Created", nil)
}

func (h *APIGOHandler) UserLogin(c *gin.Context) {
	var (
		reqBody  = models.UserLogin{}
		respBody = models.UserLoginResponse{}
	)
	if err := c.BindJSON(&reqBody); err != nil {
		if err != nil {
			helpers.RespondErrorJSON(c, http.StatusInternalServerError, errors.New("can't bind request to struct"))
			return
		}
	}

	token, err := h.APIGOUsecase.UserLogin(c, reqBody)
	if err != nil {
		helpers.RespondJSON(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	respBody.Token = *token

	helpers.RespondJSON(c, http.StatusOK, respBody, nil)
}

func (h *APIGOHandler) GetUserDetail(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.Split(auth, "Bearer ")[1]

	claims, err := helpers.GetTokenClaims(token)
	if err != nil {
		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	user, err := h.APIGOUsecase.GetUserDetail(c, claims.Email)
	if err != nil {
		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	helpers.RespondJSON(c, http.StatusOK, user, nil)
}

func (h *APIGOHandler) UserUpdate(c *gin.Context) {
	reqBody := models.User{}

	auth := c.GetHeader("Authorization")
	token := strings.Split(auth, "Bearer ")[1]

	claims, err := helpers.GetTokenClaims(token)
	if err != nil {
		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	if err := c.BindJSON(&reqBody); err != nil {
		if err != nil {
			helpers.RespondErrorJSON(c, http.StatusInternalServerError, errors.New("can't bind request to struct"))
			return
		}
	}

	if valErrors, err := h.APIGOUsecase.UserUpdate(c, claims.Email, reqBody); err != nil {
		if valErrors != nil {
			c.JSON(http.StatusUnprocessableEntity, valErrors)
			return
		}

		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	helpers.RespondJSON(c, http.StatusOK, "User Updated", nil)
}
