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

// PostingRuleHandler handles posting rule version HTTP requests.
type PostingRuleHandler struct {
	svc      services.PostingRuleService
	validate *validator.Validate
}

// NewPostingRuleHandler creates a new posting rule handler.
func NewPostingRuleHandler(svc services.PostingRuleService) *PostingRuleHandler {
	return &PostingRuleHandler{svc: svc, validate: validator.New()}
}

// RegisterRoutes registers posting rule routes.
func (h *PostingRuleHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/posting-rules")
	{
		g.POST("", h.CreatePostingRuleVersion)
		g.GET("/:id", h.GetPostingRuleVersion)
		g.PUT("/:id", h.UpdatePostingRuleVersion)
		g.DELETE("/:id", h.DeletePostingRuleVersion)
	}
}

// CreatePostingRule godoc
// @Summary Create a posting rule version
// @Description Create a new posting rule version for a source document type
// @Tags posting-rules
// @Accept json
// @Produce json
// @Param posting_rule body dto.CreatePostingRuleVersionRequest true "Posting rule version data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /posting-rules [post]
func (h *PostingRuleHandler) CreatePostingRuleVersion(c *gin.Context) {
	var req dto.CreatePostingRuleVersionRequest
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
	result, err := h.svc.CreatePostingRuleVersion(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create posting rule version", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Posting rule version created successfully", result)
}

// GetPostingRule godoc
// @Summary Get a posting rule version
// @Description Get a posting rule version by ID
// @Tags posting-rules
// @Produce json
// @Param id path string true "Posting rule version ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /posting-rules/{id} [get]
func (h *PostingRuleHandler) GetPostingRuleVersion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	result, err := h.svc.GetPostingRuleVersion(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get posting rule version", err)
		return
	}
	if result == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Posting rule version not found", errors.New("posting rule version not found"))
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Posting rule version retrieved successfully", result)
}

// UpdatePostingRule godoc
// @Summary Update a posting rule version
// @Description Update a posting rule version
// @Tags posting-rules
// @Accept json
// @Produce json
// @Param id path string true "Posting rule version ID"
// @Param posting_rule body dto.UpdatePostingRuleVersionRequest true "Posting rule version fields"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /posting-rules/{id} [put]
func (h *PostingRuleHandler) UpdatePostingRuleVersion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	var req dto.UpdatePostingRuleVersionRequest
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
	result, err := h.svc.UpdatePostingRuleVersion(c.Request.Context(), id, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update posting rule version", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Posting rule version updated successfully", result)
}

// DeletePostingRule godoc
// @Summary Delete a posting rule version
// @Description Delete a posting rule version
// @Tags posting-rules
// @Param id path string true "Posting rule version ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /posting-rules/{id} [delete]
func (h *PostingRuleHandler) DeletePostingRuleVersion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	if err := h.svc.DeletePostingRuleVersion(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete posting rule version", err)
		return
	}
	c.Status(http.StatusNoContent)
}
