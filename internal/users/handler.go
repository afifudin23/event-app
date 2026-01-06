package users

import (
	"event-app/internal/common"
	"event-app/internal/users/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

type Params struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.Service.FindAll()
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToListResponse(users)))
}

func (h *Handler) GetUserByID(c *gin.Context) {
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	user, err := h.Service.FindByID(params.ID)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToDetailResponse(user)))
}

func (h *Handler) CreateUser(c *gin.Context) {
	var payload dto.UserRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	user, err := h.Service.Create(payload)

	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusCreated, common.SuccessResponse(dto.ToSuccessResponse(user.ID)))
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var params Params
	var payload dto.UserRequest

	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	user, err := h.Service.Update(params.ID, payload)

	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToSuccessResponse(user.ID)))
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.ValidationError(details))
		return
	}

	_, err := h.Service.Delete(params.ID)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(dto.ToSuccessResponse(params.ID)))
}
