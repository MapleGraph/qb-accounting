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

// BookHandler handles accounting book HTTP requests.
type BookHandler struct {
	svc      services.BookService
	validate *validator.Validate
}

// NewBookHandler creates a new book handler.
func NewBookHandler(svc services.BookService) *BookHandler {
	return &BookHandler{
		svc:      svc,
		validate: validator.New(),
	}
}

// RegisterRoutes registers book routes under the given router group.
func (h *BookHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/books")
	{
		g.POST("", h.CreateBook)
		g.GET("/company/:company_id", h.ListBooksByCompany)
		g.GET("/:id", h.GetBook)
		g.PUT("/:id", h.UpdateBook)
		g.DELETE("/:id", h.DeleteBook)
	}
}

// CreateBook godoc
// @Summary Create a book
// @Description Create a new accounting book for a company
// @Tags books
// @Accept json
// @Produce json
// @Param book body dto.CreateBookRequest true "Book data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req dto.CreateBookRequest
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
	result, err := h.svc.CreateBook(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create book", err)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Book created successfully", result)
}

// GetBook godoc
// @Summary Get a book
// @Description Get an accounting book by ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	result, err := h.svc.GetBook(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get book", err)
		return
	}
	if result == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Book not found", errors.New("book not found"))
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Book retrieved successfully", result)
}

// ListBooksByCompany godoc
// @Summary List books by company
// @Description List all accounting books for a company
// @Tags books
// @Produce json
// @Param company_id path string true "Company ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /books/company/{company_id} [get]
func (h *BookHandler) ListBooksByCompany(c *gin.Context) {
	companyID := c.Param("company_id")
	if companyID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "company_id is required", errors.New("missing company_id parameter"))
		return
	}
	result, err := h.svc.ListBooksByCompany(c.Request.Context(), companyID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list books", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Books retrieved successfully", result)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update an accounting book
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body dto.UpdateBookRequest true "Book fields"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	var req dto.UpdateBookRequest
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
	result, err := h.svc.UpdateBook(c.Request.Context(), id, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update book", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Book updated successfully", result)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete an accounting book
// @Tags books
// @Param id path string true "Book ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required", errors.New("missing id parameter"))
		return
	}
	if err := h.svc.DeleteBook(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete book", err)
		return
	}
	c.Status(http.StatusNoContent)
}
