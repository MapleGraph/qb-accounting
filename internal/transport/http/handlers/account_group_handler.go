package handlers

import (
	"errors"
	"net/http"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/services"
	"qb-accounting/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AccountGroupHandler handles account group HTTP requests.
type AccountGroupHandler struct {
	svc      services.AccountGroupService
	validate *validator.Validate
}

// NewAccountGroupHandler creates a new account group handler.
func NewAccountGroupHandler(svc services.AccountGroupService) *AccountGroupHandler {
	return &AccountGroupHandler{svc: svc, validate: validator.New()}
}

// RegisterRoutes registers account group routes.
func (h *AccountGroupHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/account-groups")
	{
		g.POST("", h.CreateAccountGroup)
		g.GET("/book/:book_id", h.ListAccountGroupsByBook)
		g.GET("/:id", h.GetAccountGroup)
		g.PUT("/:id", h.UpdateAccountGroup)
		g.DELETE("/:id", h.DeleteAccountGroup)
	}
}

// CreateAccountGroup godoc
// @Summary Create an account group
// @Description Create a new account group under a book
// @Tags account-groups
// @Accept json
// @Produce json
// @Param account_group body dto.CreateAccountGroupRequest true "Account group data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /account-groups [post]
func (h *AccountGroupHandler) CreateAccountGroup(c *gin.Context) {
	var req dto.CreateAccountGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	if err := h.validate.Struct(&req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, utils.ValidationError(validationErrors))
			return
		}
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err)
		return
	}
	result, err := h.svc.CreateAccountGroup(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create account group", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Account group created successfully", result)
}

// GetAccountGroup godoc
// @Summary Get an account group
// @Description Get an account group by ID
// @Tags account-groups
// @Produce json
// @Param id path string true "Account group ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /account-groups/{id} [get]
func (h *AccountGroupHandler) GetAccountGroup(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	result, err := h.svc.GetAccountGroup(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get account group", err)
		return
	}
	if result == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Account group not found", errors.New("account group not found"))
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Account group retrieved successfully", result)
}

// ListAccountGroupsByBook godoc
// @Summary List account groups by book
// @Description List account groups for a book
// @Tags account-groups
// @Produce json
// @Param book_id path string true "Book ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /account-groups/book/{book_id} [get]
func (h *AccountGroupHandler) ListAccountGroupsByBook(c *gin.Context) {
	bookID := c.Param("book_id")
	if bookID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "book_id is required", errors.New("missing book_id parameter"))
		return
	}
	result, err := h.svc.ListAccountGroupsByBook(c.Request.Context(), bookID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list account groups", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Account groups retrieved successfully", result)
}

// UpdateAccountGroup godoc
// @Summary Update an account group
// @Description Update an account group
// @Tags account-groups
// @Accept json
// @Produce json
// @Param id path string true "Account group ID"
// @Param account_group body dto.UpdateAccountGroupRequest true "Account group fields"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /account-groups/{id} [put]
func (h *AccountGroupHandler) UpdateAccountGroup(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	var req dto.UpdateAccountGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	if err := h.validate.Struct(&req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, utils.ValidationError(validationErrors))
			return
		}
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err)
		return
	}
	result, err := h.svc.UpdateAccountGroup(c.Request.Context(), id, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update account group", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Account group updated successfully", result)
}

// DeleteAccountGroup godoc
// @Summary Delete an account group
// @Description Delete an account group
// @Tags account-groups
// @Param id path string true "Account group ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /account-groups/{id} [delete]
func (h *AccountGroupHandler) DeleteAccountGroup(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	if err := h.svc.DeleteAccountGroup(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete account group", err)
		return
	}
	c.Status(http.StatusNoContent)
}
