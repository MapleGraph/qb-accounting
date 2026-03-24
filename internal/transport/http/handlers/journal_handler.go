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

// JournalHandler handles journal HTTP requests.
type JournalHandler struct {
	svc      services.JournalService
	validate *validator.Validate
}

// NewJournalHandler creates a new journal handler.
func NewJournalHandler(svc services.JournalService) *JournalHandler {
	return &JournalHandler{svc: svc, validate: validator.New()}
}

// RegisterRoutes registers journal routes.
func (h *JournalHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/journals")
	{
		g.POST("", h.CreateJournal)
		g.POST("/post", h.PostJournal)
		g.POST("/reverse", h.ReverseJournal)
		g.GET("/book/:book_id", h.ListJournalsByBook)
		g.GET("/period/:period_id", h.ListJournalsByPeriod)
		g.GET("/:id", h.GetJournal)
		g.DELETE("/:id", h.DeleteJournal)
	}
}

// CreateJournal godoc
// @Summary Create a journal
// @Description Create a new journal (typically draft)
// @Tags journals
// @Accept json
// @Produce json
// @Param journal body dto.CreateJournalRequest true "Journal data with lines"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journals [post]
func (h *JournalHandler) CreateJournal(c *gin.Context) {
	var req dto.CreateJournalRequest
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
	result, err := h.svc.CreateJournal(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create journal", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Journal created successfully", result)
}

// GetJournal godoc
// @Summary Get a journal
// @Description Get a journal by ID; set include_lines=true to load journal lines
// @Tags journals
// @Produce json
// @Param id path string true "Journal ID"
// @Param include_lines query bool false "Include journal lines"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journals/{id} [get]
func (h *JournalHandler) GetJournal(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	includeLines := c.Query("include_lines") == "true"
	result, err := h.svc.GetJournal(c.Request.Context(), id, includeLines)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get journal", err)
		return
	}
	if result == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Journal not found", errors.New("journal not found"))
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Journal retrieved successfully", result)
}

// ListJournalsByBook godoc
// @Summary List journals by book
// @Description List journals for an accounting book
// @Tags journals
// @Produce json
// @Param book_id path string true "Book ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journals/book/{book_id} [get]
func (h *JournalHandler) ListJournalsByBook(c *gin.Context) {
	bookID := c.Param("book_id")
	if bookID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "book_id is required", errors.New("missing book_id parameter"))
		return
	}
	result, err := h.svc.ListJournalsByBook(c.Request.Context(), bookID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list journals", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Journals retrieved successfully", result)
}

// ListJournalsByPeriod godoc
// @Summary List journals by period
// @Description List journals for an accounting period
// @Tags journals
// @Produce json
// @Param period_id path string true "Accounting period ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journals/period/{period_id} [get]
func (h *JournalHandler) ListJournalsByPeriod(c *gin.Context) {
	periodID := c.Param("period_id")
	if periodID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "period_id is required", errors.New("missing period_id parameter"))
		return
	}
	result, err := h.svc.ListJournalsByPeriod(c.Request.Context(), periodID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list journals", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Journals retrieved successfully", result)
}

// PostJournal godoc
// @Summary Post a journal
// @Description Post a draft journal to the ledger
// @Tags journals
// @Accept json
// @Produce json
// @Param journal body dto.PostJournalRequest true "Post journal request"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journals/post [post]
func (h *JournalHandler) PostJournal(c *gin.Context) {
	var req dto.PostJournalRequest
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
	result, err := h.svc.PostJournal(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to post journal", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Journal posted successfully", result)
}

// ReverseJournal godoc
// @Summary Reverse a journal
// @Description Create a reversing journal entry for a posted journal
// @Tags journals
// @Accept json
// @Produce json
// @Param journal body dto.ReverseJournalRequest true "Reverse journal request"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journals/reverse [post]
func (h *JournalHandler) ReverseJournal(c *gin.Context) {
	var req dto.ReverseJournalRequest
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
	result, err := h.svc.ReverseJournal(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to reverse journal", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Journal reversed successfully", result)
}

// DeleteJournal godoc
// @Summary Delete a journal
// @Description Delete a journal
// @Tags journals
// @Param id path string true "Journal ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /journals/{id} [delete]
func (h *JournalHandler) DeleteJournal(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	if err := h.svc.DeleteJournal(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete journal", err)
		return
	}
	c.Status(http.StatusNoContent)
}
