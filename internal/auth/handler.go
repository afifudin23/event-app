package auth

import (
	"event-app/internal/auth/dto"
	"event-app/internal/common"
	"event-app/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
	Cfg     *config.Config
}

func NewHandler(service Service, cfg *config.Config) *Handler {
	return &Handler{Service: service, Cfg: cfg}
}

func (h *Handler) Login(c *gin.Context) {
	var payload dto.UserLoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}
	user, accessToken, err := h.Service.Login(payload)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(ToLoginResponse(*user, *accessToken)))
}

func (h *Handler) Register(c *gin.Context) {
	var payload dto.UserRegisterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}
	user, accessToken, err := h.Service.Register(payload)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(ToLoginResponse(*user, *accessToken)))
}
