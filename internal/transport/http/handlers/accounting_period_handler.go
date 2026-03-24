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

// AccountingPeriodHandler handles accounting period HTTP requests.
type AccountingPeriodHandler struct {
	svc      services.AccountingPeriodService
	validate *validator.Validate
}

// NewAccountingPeriodHandler creates a new accounting period handler.
func NewAccountingPeriodHandler(svc services.AccountingPeriodService) *AccountingPeriodHandler {
	return &AccountingPeriodHandler{svc: svc, validate: validator.New()}
}

// RegisterRoutes registers accounting period routes.
func (h *AccountingPeriodHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/accounting-periods")
	{
		g.POST("", h.CreateAccountingPeriod)
		g.GET("/book/:book_id", h.ListAccountingPeriodsByBook)
		g.GET("/fiscal-year/:fiscal_year_id", h.ListAccountingPeriodsByFiscalYear)
		g.GET("/:id", h.GetAccountingPeriod)
		g.PUT("/:id", h.UpdateAccountingPeriod)
		g.DELETE("/:id", h.DeleteAccountingPeriod)
	}
}

// CreateAccountingPeriod godoc
// @Summary Create an accounting period
// @Description Create a new accounting period
// @Tags accounting-periods
// @Accept json
// @Produce json
// @Param body body dto.CreateAccountingPeriodRequest true "Period data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /accounting-periods [post]
func (h *AccountingPeriodHandler) CreateAccountingPeriod(c *gin.Context) {
	var req dto.CreateAccountingPeriodRequest
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
	result, err := h.svc.CreateAccountingPeriod(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create accounting period", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Accounting period created successfully", result)
}

// GetAccountingPeriod godoc
// @Summary Get an accounting period
// @Tags accounting-periods
// @Produce json
// @Param id path string true "Period ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /accounting-periods/{id} [get]
func (h *AccountingPeriodHandler) GetAccountingPeriod(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	result, err := h.svc.GetAccountingPeriod(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get accounting period", err)
		return
	}
	if result == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Accounting period not found", errors.New("accounting period not found"))
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Accounting period retrieved successfully", result)
}

// ListAccountingPeriodsByBook godoc
// @Summary List accounting periods by book
// @Tags accounting-periods
// @Produce json
// @Param book_id path string true "Book ID"
// @Success 200 {object} utils.Response
// @Router /accounting-periods/book/{book_id} [get]
func (h *AccountingPeriodHandler) ListAccountingPeriodsByBook(c *gin.Context) {
	bookID := c.Param("book_id")
	if bookID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "book_id is required", errors.New("missing book_id parameter"))
		return
	}
	result, err := h.svc.ListAccountingPeriodsByBook(c.Request.Context(), bookID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list accounting periods", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Accounting periods retrieved successfully", result)
}

// ListAccountingPeriodsByFiscalYear godoc
// @Summary List accounting periods by fiscal year
// @Tags accounting-periods
// @Produce json
// @Param fiscal_year_id path string true "Fiscal year ID"
// @Success 200 {object} utils.Response
// @Router /accounting-periods/fiscal-year/{fiscal_year_id} [get]
func (h *AccountingPeriodHandler) ListAccountingPeriodsByFiscalYear(c *gin.Context) {
	fiscalYearID := c.Param("fiscal_year_id")
	if fiscalYearID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "fiscal_year_id is required", errors.New("missing fiscal_year_id parameter"))
		return
	}
	result, err := h.svc.ListAccountingPeriodsByFiscalYear(c.Request.Context(), fiscalYearID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list accounting periods", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Accounting periods retrieved successfully", result)
}

// UpdateAccountingPeriod godoc
// @Summary Update an accounting period
// @Tags accounting-periods
// @Accept json
// @Produce json
// @Param id path string true "Period ID"
// @Param body body dto.UpdateAccountingPeriodRequest true "Updates"
// @Success 200 {object} utils.Response
// @Router /accounting-periods/{id} [put]
func (h *AccountingPeriodHandler) UpdateAccountingPeriod(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	var req dto.UpdateAccountingPeriodRequest
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
	result, err := h.svc.UpdateAccountingPeriod(c.Request.Context(), id, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update accounting period", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Accounting period updated successfully", result)
}

// DeleteAccountingPeriod godoc
// @Summary Delete an accounting period
// @Tags accounting-periods
// @Param id path string true "Period ID"
// @Success 204 "No Content"
// @Router /accounting-periods/{id} [delete]
func (h *AccountingPeriodHandler) DeleteAccountingPeriod(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	if err := h.svc.DeleteAccountingPeriod(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete accounting period", err)
		return
	}
	c.Status(http.StatusNoContent)
}
