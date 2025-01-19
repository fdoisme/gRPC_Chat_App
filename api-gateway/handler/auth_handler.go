package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"project/api-gateway/database"
	"project/api-gateway/model"
	"project/api-gateway/service"
)

type AuthController struct {
	service service.Service
	logger  *zap.Logger
	cacher  database.Cacher
}

func NewAuthController(service service.Service, logger *zap.Logger, cacher database.Cacher) *AuthController {
	return &AuthController{service, logger, cacher}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := ctrl.service.User.CreateUser(user)
	if err != nil {
		BadResponse(c, "Email Already Registered", http.StatusBadRequest)
		return
	}

	res, err := ctrl.service.Auth.Register(user)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	emailData := EmailData{ID: uuid.New(), OTP: res.Otp}
	_, err = ctrl.service.Email.Send(user.Email, "Chateo OTP", "otp", emailData)
	if err != nil {
		ctrl.logger.Error("failed to send email", zap.Error(err))
		BadResponse(c, "failed to send email", http.StatusInternalServerError)
		return
	}

	GoodResponseWithData(c, "registration success. otp sent", http.StatusOK, nil)
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := ctrl.service.Auth.Login(user)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	emailData := EmailData{ID: uuid.New(), OTP: res.Otp}
	_, err = ctrl.service.Email.Send(user.Email, "Chateo OTP", "otp", emailData)
	if err != nil {
		ctrl.logger.Error("failed to send email", zap.Error(err))
		BadResponse(c, "failed to send email", http.StatusInternalServerError)
		return
	}

	GoodResponseWithData(c, "login success. otp sent", http.StatusOK, nil)
}

func (ctrl *AuthController) ValidateOtp(c *gin.Context) {
	otpID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		BadResponse(c, "invalid otp", http.StatusUnprocessableEntity)
	}

	otp := model.Otp{ID: otpID}
	if err = c.ShouldBindJSON(&otp); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := ctrl.service.Auth.ValidateOtp(otp)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if res == nil {
		BadResponse(c, "not found", http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "otp is valid", http.StatusOK, gin.H{"token": res})
}

type EmailData struct {
	ID  uuid.UUID
	OTP string
}
