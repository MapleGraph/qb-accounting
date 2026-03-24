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

// FiscalYearHandler handles fiscal year HTTP requests.
type FiscalYearHandler struct {
	svc      services.FiscalYearService
	validate *validator.Validate
}

// NewFiscalYearHandler creates a new fiscal year handler.
func NewFiscalYearHandler(svc services.FiscalYearService) *FiscalYearHandler {
	return &FiscalYearHandler{svc: svc, validate: validator.New()}
}

// RegisterRoutes registers fiscal year routes.
func (h *FiscalYearHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/fiscal-years")
	{
		g.POST("", h.CreateFiscalYear)
		g.GET("/book/:book_id", h.ListFiscalYearsByBook)
		g.GET("/:id", h.GetFiscalYear)
		g.PUT("/:id", h.UpdateFiscalYear)
		g.DELETE("/:id", h.DeleteFiscalYear)
	}
}

// CreateFiscalYear godoc
// @Summary Create a fiscal year
// @Description Create a new fiscal year for a book
// @Tags fiscal-years
// @Accept json
// @Produce json
// @Param fiscal_year body dto.CreateFiscalYearRequest true "Fiscal year data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /fiscal-years [post]
func (h *FiscalYearHandler) CreateFiscalYear(c *gin.Context) {
	var req dto.CreateFiscalYearRequest
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
	result, err := h.svc.CreateFiscalYear(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create fiscal year", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Fiscal year created successfully", result)
}

// GetFiscalYear godoc
// @Summary Get a fiscal year
// @Description Get a fiscal year by ID
// @Tags fiscal-years
// @Produce json
// @Param id path string true "Fiscal year ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /fiscal-years/{id} [get]
func (h *FiscalYearHandler) GetFiscalYear(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	result, err := h.svc.GetFiscalYear(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get fiscal year", err)
		return
	}
	if result == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Fiscal year not found", errors.New("fiscal year not found"))
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Fiscal year retrieved successfully", result)
}

// ListFiscalYearsByBook godoc
// @Summary List fiscal years by book
// @Description List fiscal years for an accounting book
// @Tags fiscal-years
// @Produce json
// @Param book_id path string true "Book ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /fiscal-years/book/{book_id} [get]
func (h *FiscalYearHandler) ListFiscalYearsByBook(c *gin.Context) {
	bookID := c.Param("book_id")
	if bookID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "book_id is required", errors.New("missing book_id parameter"))
		return
	}
	result, err := h.svc.ListFiscalYearsByBook(c.Request.Context(), bookID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list fiscal years", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Fiscal years retrieved successfully", result)
}

// UpdateFiscalYear godoc
// @Summary Update a fiscal year
// @Description Update a fiscal year
// @Tags fiscal-years
// @Accept json
// @Produce json
// @Param id path string true "Fiscal year ID"
// @Param fiscal_year body dto.UpdateFiscalYearRequest true "Fiscal year fields"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /fiscal-years/{id} [put]
func (h *FiscalYearHandler) UpdateFiscalYear(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	var req dto.UpdateFiscalYearRequest
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
	result, err := h.svc.UpdateFiscalYear(c.Request.Context(), id, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update fiscal year", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Fiscal year updated successfully", result)
}

// DeleteFiscalYear godoc
// @Summary Delete a fiscal year
// @Description Delete a fiscal year
// @Tags fiscal-years
// @Param id path string true "Fiscal year ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /fiscal-years/{id} [delete]
func (h *FiscalYearHandler) DeleteFiscalYear(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	if err := h.svc.DeleteFiscalYear(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete fiscal year", err)
		return
	}
	c.Status(http.StatusNoContent)
}
