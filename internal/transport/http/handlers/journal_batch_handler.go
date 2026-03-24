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

// JournalBatchHandler handles journal batch HTTP requests.
type JournalBatchHandler struct {
	svc      services.JournalBatchService
	validate *validator.Validate
}

// NewJournalBatchHandler creates a new journal batch handler.
func NewJournalBatchHandler(svc services.JournalBatchService) *JournalBatchHandler {
	return &JournalBatchHandler{svc: svc, validate: validator.New()}
}

// RegisterRoutes registers journal batch routes.
func (h *JournalBatchHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/journal-batches")
	{
		g.POST("", h.CreateJournalBatch)
		g.GET("/book/:book_id", h.ListJournalBatchesByBook)
		g.GET("/:id", h.GetJournalBatch)
		g.PUT("/:id", h.UpdateJournalBatch)
		g.DELETE("/:id", h.DeleteJournalBatch)
	}
}

// CreateJournalBatch godoc
// @Summary Create a journal batch
// @Description Create a new journal batch
// @Tags journal-batches
// @Accept json
// @Produce json
// @Param journal_batch body dto.CreateJournalBatchRequest true "Journal batch data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journal-batches [post]
func (h *JournalBatchHandler) CreateJournalBatch(c *gin.Context) {
	var req dto.CreateJournalBatchRequest
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
	result, err := h.svc.CreateJournalBatch(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create journal batch", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Journal batch created successfully", result)
}

// GetJournalBatch godoc
// @Summary Get a journal batch
// @Description Get a journal batch by ID
// @Tags journal-batches
// @Produce json
// @Param id path string true "Journal batch ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journal-batches/{id} [get]
func (h *JournalBatchHandler) GetJournalBatch(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	result, err := h.svc.GetJournalBatch(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get journal batch", err)
		return
	}
	if result == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Journal batch not found", errors.New("journal batch not found"))
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Journal batch retrieved successfully", result)
}

// ListJournalBatchesByBook godoc
// @Summary List journal batches by book
// @Description List journal batches for a book
// @Tags journal-batches
// @Produce json
// @Param book_id path string true "Book ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journal-batches/book/{book_id} [get]
func (h *JournalBatchHandler) ListJournalBatchesByBook(c *gin.Context) {
	bookID := c.Param("book_id")
	if bookID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "book_id is required", errors.New("missing book_id parameter"))
		return
	}
	result, err := h.svc.ListJournalBatchesByBook(c.Request.Context(), bookID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list journal batches", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Journal batches retrieved successfully", result)
}

// UpdateJournalBatch godoc
// @Summary Update a journal batch
// @Description Update a journal batch
// @Tags journal-batches
// @Accept json
// @Produce json
// @Param id path string true "Journal batch ID"
// @Param journal_batch body dto.UpdateJournalBatchRequest true "Journal batch fields"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journal-batches/{id} [put]
func (h *JournalBatchHandler) UpdateJournalBatch(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	var req dto.UpdateJournalBatchRequest
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
	result, err := h.svc.UpdateJournalBatch(c.Request.Context(), id, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update journal batch", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Journal batch updated successfully", result)
}

// DeleteJournalBatch godoc
// @Summary Delete a journal batch
// @Description Delete a journal batch
// @Tags journal-batches
// @Param id path string true "Journal batch ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journal-batches/{id} [delete]
func (h *JournalBatchHandler) DeleteJournalBatch(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	if err := h.svc.DeleteJournalBatch(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete journal batch", err)
		return
	}
	c.Status(http.StatusNoContent)
}
