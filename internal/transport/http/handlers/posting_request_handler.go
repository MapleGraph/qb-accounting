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

// PostingRequestHandler handles posting request HTTP requests.
type PostingRequestHandler struct {
	svc      services.PostingRequestService
	validate *validator.Validate
}

// NewPostingRequestHandler creates a new posting request handler.
func NewPostingRequestHandler(svc services.PostingRequestService) *PostingRequestHandler {
	return &PostingRequestHandler{svc: svc, validate: validator.New()}
}

// RegisterRoutes registers posting request routes.
func (h *PostingRequestHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/posting-requests")
	{
		g.POST("", h.CreatePostingRequest)
		g.GET("/:id", h.GetPostingRequest)
	}
}

// CreatePostingRequest godoc
// @Summary Create a posting request
// @Description Enqueue a posting request from a source document
// @Tags posting-requests
// @Accept json
// @Produce json
// @Param posting_request body dto.CreatePostingRequest true "Posting request data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /posting-requests [post]
func (h *PostingRequestHandler) CreatePostingRequest(c *gin.Context) {
	var req dto.CreatePostingRequest
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
	result, err := h.svc.CreatePostingRequest(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create posting request", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Posting request created successfully", result)
}

// GetPostingRequest godoc
// @Summary Get a posting request
// @Description Get a posting request by ID
// @Tags posting-requests
// @Produce json
// @Param id path string true "Posting request ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /posting-requests/{id} [get]
func (h *PostingRequestHandler) GetPostingRequest(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	result, err := h.svc.GetPostingRequest(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get posting request", err)
		return
	}
	if result == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Posting request not found", errors.New("posting request not found"))
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Posting request retrieved successfully", result)
}
