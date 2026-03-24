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

// OpenItemHandler handles open item HTTP requests.
type OpenItemHandler struct {
	svc      services.OpenItemService
	validate *validator.Validate
}

// NewOpenItemHandler creates a new open item handler.
func NewOpenItemHandler(svc services.OpenItemService) *OpenItemHandler {
	return &OpenItemHandler{svc: svc, validate: validator.New()}
}

// RegisterRoutes registers open item routes.
func (h *OpenItemHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/open-items")
	{
		g.POST("", h.CreateOpenItem)
		g.GET("/book/:book_id", h.ListOpenItemsByBook)
		g.GET("/party/:party_id", h.ListOpenItemsByParty)
		g.POST("/allocate", h.AllocateOpenItems)
		g.POST("/adjustments", h.CreateAdjustment)
		g.GET("/:id", h.GetOpenItem)
		g.PUT("/:id", h.UpdateOpenItem)
		g.DELETE("/:id", h.DeleteOpenItem)
	}
}

// CreateOpenItem godoc
// @Summary Create an open item
// @Tags open-items
// @Accept json
// @Produce json
// @Param body body dto.CreateOpenItemRequest true "Open item"
// @Success 201 {object} utils.Response
// @Router /open-items [post]
func (h *OpenItemHandler) CreateOpenItem(c *gin.Context) {
	var req dto.CreateOpenItemRequest
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
	result, err := h.svc.CreateOpenItem(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create open item", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Open item created successfully", result)
}

// GetOpenItem godoc
// @Summary Get an open item
// @Tags open-items
// @Param id path string true "Open item ID"
// @Success 200 {object} utils.Response
// @Router /open-items/{id} [get]
func (h *OpenItemHandler) GetOpenItem(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	result, err := h.svc.GetOpenItem(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get open item", err)
		return
	}
	if result == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Open item not found", errors.New("open item not found"))
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Open item retrieved successfully", result)
}

// ListOpenItemsByBook godoc
// @Summary List open items by book
// @Tags open-items
// @Param book_id path string true "Book ID"
// @Success 200 {object} utils.Response
// @Router /open-items/book/{book_id} [get]
func (h *OpenItemHandler) ListOpenItemsByBook(c *gin.Context) {
	bookID := c.Param("book_id")
	if bookID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "book_id is required", errors.New("missing book_id parameter"))
		return
	}
	result, err := h.svc.GetOpenItemsByBook(c.Request.Context(), bookID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list open items", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Open items retrieved successfully", result)
}

// ListOpenItemsByParty godoc
// @Summary List open items by party
// @Tags open-items
// @Param party_id path string true "Party ID"
// @Success 200 {object} utils.Response
// @Router /open-items/party/{party_id} [get]
func (h *OpenItemHandler) ListOpenItemsByParty(c *gin.Context) {
	partyID := c.Param("party_id")
	if partyID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "party_id is required", errors.New("missing party_id parameter"))
		return
	}
	result, err := h.svc.GetOpenItemsByParty(c.Request.Context(), partyID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list open items", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Open items retrieved successfully", result)
}

// UpdateOpenItem godoc
// @Summary Update an open item
// @Tags open-items
// @Accept json
// @Produce json
// @Param id path string true "Open item ID"
// @Param body body dto.UpdateOpenItemRequest true "Updates"
// @Success 200 {object} utils.Response
// @Router /open-items/{id} [put]
func (h *OpenItemHandler) UpdateOpenItem(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	var req dto.UpdateOpenItemRequest
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
	result, err := h.svc.UpdateOpenItem(c.Request.Context(), id, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update open item", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Open item updated successfully", result)
}

// DeleteOpenItem godoc
// @Summary Delete an open item
// @Tags open-items
// @Param id path string true "Open item ID"
// @Success 204 "No Content"
// @Router /open-items/{id} [delete]
func (h *OpenItemHandler) DeleteOpenItem(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	if err := h.svc.DeleteOpenItem(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete open item", err)
		return
	}
	c.Status(http.StatusNoContent)
}

// AllocateOpenItems godoc
// @Summary Allocate between open items
// @Tags open-items
// @Accept json
// @Produce json
// @Param body body dto.AllocationRequest true "Allocation"
// @Success 201 {object} utils.Response
// @Router /open-items/allocate [post]
func (h *OpenItemHandler) AllocateOpenItems(c *gin.Context) {
	var req dto.AllocationRequest
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
	result, err := h.svc.CreateAllocation(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to allocate open items", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Allocation created successfully", result)
}

// CreateAdjustment godoc
// @Summary Create an open item adjustment
// @Tags open-items
// @Accept json
// @Produce json
// @Param body body dto.AdjustmentRequest true "Adjustment"
// @Success 201 {object} utils.Response
// @Router /open-items/adjustments [post]
func (h *OpenItemHandler) CreateAdjustment(c *gin.Context) {
	var req dto.AdjustmentRequest
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
	result, err := h.svc.CreateAdjustment(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create adjustment", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Adjustment created successfully", result)
}
