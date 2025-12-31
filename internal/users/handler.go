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
	c.JSON(http.StatusOK, common.SuccessResponse(ToListResponse(users)))
}

func (h *Handler) GetUserByID(c *gin.Context) {
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.BadRequestError(details))
		return
	}

	user, err := h.Service.FindByID(params.ID)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(ToResponse(*user)))
}

func (h *Handler) CreateUser(c *gin.Context) {
	var payload dto.UserRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.BadRequestError(details))
		return
	}

	success, err := h.Service.Create(User{
		Fullname: payload.Fullname,
		Email:    payload.Email,
		Password: payload.Password,
	})
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusCreated, common.SuccessResponse(ToCreateResponse(success)))
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.BadRequestError(details))
		return
	}
	var payload dto.UserRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.BadRequestError(details))
		return
	}

	success, err := h.Service.Update(params.ID, User{
		Fullname: payload.Fullname,
		Email:    payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(ToUpdateResponse(success)))
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		details := common.ErrorValidation(err)
		common.ErrorHandler(c, common.BadRequestError(details))
		return
	}

	success, err := h.Service.Delete(params.ID)
	if err != nil {
		common.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(ToDeleteResponse(success)))
}
