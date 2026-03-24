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

// AccountHandler handles account HTTP requests.
type AccountHandler struct {
	svc      services.AccountService
	validate *validator.Validate
}

// NewAccountHandler creates a new account handler.
func NewAccountHandler(svc services.AccountService) *AccountHandler {
	return &AccountHandler{svc: svc, validate: validator.New()}
}

// RegisterRoutes registers account routes.
func (h *AccountHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/accounts")
	{
		g.POST("", h.CreateAccount)
		g.GET("/book/:book_id", h.ListAccountsByBook)
		g.GET("/company/:company_id", h.ListAccountsByCompany)
		g.GET("/:id", h.GetAccount)
		g.PUT("/:id", h.UpdateAccount)
		g.DELETE("/:id", h.DeleteAccount)
	}
}

// CreateAccount godoc
// @Summary Create an account
// @Description Create a new ledger account
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body dto.CreateAccountRequest true "Account data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /accounts [post]
func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req dto.CreateAccountRequest
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
	result, err := h.svc.CreateAccount(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create account", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Account created successfully", result)
}

// GetAccount godoc
// @Summary Get an account
// @Description Get an account by ID
// @Tags accounts
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /accounts/{id} [get]
func (h *AccountHandler) GetAccount(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	result, err := h.svc.GetAccount(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get account", err)
		return
	}
	if result == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Account not found", errors.New("account not found"))
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Account retrieved successfully", result)
}

// ListAccountsByBook godoc
// @Summary List accounts by book
// @Description List accounts for an accounting book
// @Tags accounts
// @Produce json
// @Param book_id path string true "Book ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /accounts/book/{book_id} [get]
func (h *AccountHandler) ListAccountsByBook(c *gin.Context) {
	bookID := c.Param("book_id")
	if bookID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "book_id is required", errors.New("missing book_id parameter"))
		return
	}
	result, err := h.svc.ListAccountsByBook(c.Request.Context(), bookID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list accounts", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Accounts retrieved successfully", result)
}

// ListAccountsByCompany godoc
// @Summary List accounts by company
// @Description List accounts for a company
// @Tags accounts
// @Produce json
// @Param company_id path string true "Company ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /accounts/company/{company_id} [get]
func (h *AccountHandler) ListAccountsByCompany(c *gin.Context) {
	companyID := c.Param("company_id")
	if companyID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "company_id is required", errors.New("missing company_id parameter"))
		return
	}
	result, err := h.svc.ListAccountsByCompany(c.Request.Context(), companyID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list accounts", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Accounts retrieved successfully", result)
}

// UpdateAccount godoc
// @Summary Update an account
// @Description Update an account
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param account body dto.UpdateAccountRequest true "Account fields"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /accounts/{id} [put]
func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	var req dto.UpdateAccountRequest
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
	result, err := h.svc.UpdateAccount(c.Request.Context(), id, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update account", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Account updated successfully", result)
}

// DeleteAccount godoc
// @Summary Delete an account
// @Description Delete an account
// @Tags accounts
// @Param id path string true "Account ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /accounts/{id} [delete]
func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	if err := h.svc.DeleteAccount(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete account", err)
		return
	}
	c.Status(http.StatusNoContent)
}
