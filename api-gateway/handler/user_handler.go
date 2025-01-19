package handler

import (
	"log"
	"net/http"
	"project/api-gateway/model"
	"project/api-gateway/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	service service.Service
	logger  *zap.Logger
}

func NewUserController(service service.Service, logger *zap.Logger) *UserController {
	return &UserController{service, logger}
}

func (ctrl *UserController) Update(c *gin.Context) {
	username := c.MustGet("username")

	GoodResponseWithData(c, "profile updated", http.StatusOK, gin.H{"user_id": username})
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	query := c.Query("filter")

	resGrpc, err := ctrl.service.User.GetAllUsers(query)
	if err != nil {
		log.Println(err)
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Get All Users Success", http.StatusOK, resGrpc.Users)
}

func (ctrl *UserController) UpdateProfile(c *gin.Context) {
	email := c.MustGet("email").(string)

	userProfile := model.User{Email: email}
	if err := c.ShouldBindJSON(&userProfile); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(userProfile)

	resGrpc, err := ctrl.service.User.UpdateUser(userProfile)
	if err != nil {
		log.Println(err)
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	GoodResponseWithData(c, resGrpc.Message, http.StatusOK, nil)
}
