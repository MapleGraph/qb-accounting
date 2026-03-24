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

// VoucherSequenceHandler handles voucher sequence HTTP requests.
type VoucherSequenceHandler struct {
	svc      services.VoucherSequenceService
	validate *validator.Validate
}

// NewVoucherSequenceHandler creates a new voucher sequence handler.
func NewVoucherSequenceHandler(svc services.VoucherSequenceService) *VoucherSequenceHandler {
	return &VoucherSequenceHandler{svc: svc, validate: validator.New()}
}

// RegisterRoutes registers voucher sequence routes.
func (h *VoucherSequenceHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/voucher-sequences")
	{
		g.POST("", h.CreateVoucherSequence)
		g.GET("/:id", h.GetVoucherSequence)
		g.PUT("/:id", h.UpdateVoucherSequence)
		g.DELETE("/:id", h.DeleteVoucherSequence)
	}
}

// CreateVoucherSequence godoc
// @Summary Create a voucher sequence
// @Description Create a new voucher numbering sequence
// @Tags voucher-sequences
// @Accept json
// @Produce json
// @Param voucher_sequence body dto.CreateVoucherSequenceRequest true "Voucher sequence data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /voucher-sequences [post]
func (h *VoucherSequenceHandler) CreateVoucherSequence(c *gin.Context) {
	var req dto.CreateVoucherSequenceRequest
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
	result, err := h.svc.CreateVoucherSequence(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create voucher sequence", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Voucher sequence created successfully", result)
}

// GetVoucherSequence godoc
// @Summary Get a voucher sequence
// @Description Get a voucher sequence by ID
// @Tags voucher-sequences
// @Produce json
// @Param id path string true "Voucher sequence ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /voucher-sequences/{id} [get]
func (h *VoucherSequenceHandler) GetVoucherSequence(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	result, err := h.svc.GetVoucherSequence(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get voucher sequence", err)
		return
	}
	if result == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Voucher sequence not found", errors.New("voucher sequence not found"))
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Voucher sequence retrieved successfully", result)
}

// UpdateVoucherSequence godoc
// @Summary Update a voucher sequence
// @Description Update a voucher sequence
// @Tags voucher-sequences
// @Accept json
// @Produce json
// @Param id path string true "Voucher sequence ID"
// @Param voucher_sequence body dto.UpdateVoucherSequenceRequest true "Voucher sequence fields"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /voucher-sequences/{id} [put]
func (h *VoucherSequenceHandler) UpdateVoucherSequence(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	var req dto.UpdateVoucherSequenceRequest
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
	result, err := h.svc.UpdateVoucherSequence(c.Request.Context(), id, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update voucher sequence", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Voucher sequence updated successfully", result)
}

// DeleteVoucherSequence godoc
// @Summary Delete a voucher sequence
// @Description Delete a voucher sequence
// @Tags voucher-sequences
// @Param id path string true "Voucher sequence ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /voucher-sequences/{id} [delete]
func (h *VoucherSequenceHandler) DeleteVoucherSequence(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	if err := h.svc.DeleteVoucherSequence(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete voucher sequence", err)
		return
	}
	c.Status(http.StatusNoContent)
}
