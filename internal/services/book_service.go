package services

import (
	"context"
	"fmt"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"
)

// BookService defines accounting book operations for HTTP handlers (DTO responses).
type BookService interface {
	CreateBook(ctx context.Context, req *dto.CreateBookRequest) (*dto.BookResponse, error)
	GetBook(ctx context.Context, id string) (*dto.BookResponse, error)
	ListBooksByCompany(ctx context.Context, companyID string) ([]*dto.BookResponse, error)
	UpdateBook(ctx context.Context, id string, req *dto.UpdateBookRequest) (*dto.BookResponse, error)
	DeleteBook(ctx context.Context, id string) error
}

type bookService struct {
	repos *repository.RepositoryContainer
}

func NewBookService(repos *repository.RepositoryContainer) BookService {
	return &bookService{repos: repos}
}

func (s *bookService) CreateBook(ctx context.Context, req *dto.CreateBookRequest) (*dto.BookResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	status := models.BookStatusActive
	b := &models.Book{
		CompanyID:             req.CompanyID,
		Code:                  req.Code,
		Name:                  req.Name,
		BookType:              models.BookType(req.BookType),
		BaseCurrencyCode:      req.BaseCurrencyCode,
		ReportingCurrencyCode: req.ReportingCurrencyCode,
		Status:                status,
	}
	if err := s.repos.BookRepo.Create(ctx, b); err != nil {
		return nil, err
	}
	created, err := s.repos.BookRepo.GetByID(ctx, b.ID)
	if err != nil {
		return nil, err
	}
	return bookModelToDTO(created), nil
}

func (s *bookService) GetBook(ctx context.Context, id string) (*dto.BookResponse, error) {
	b, err := s.repos.BookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return bookModelToDTO(b), nil
}

func (s *bookService) ListBooksByCompany(ctx context.Context, companyID string) ([]*dto.BookResponse, error) {
	list, err := s.repos.BookRepo.GetByCompanyID(ctx, companyID)
	if err != nil {
		return nil, err
	}
	return booksModelToDTO(list), nil
}

func (s *bookService) UpdateBook(ctx context.Context, id string, req *dto.UpdateBookRequest) (*dto.BookResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	existing, err := s.repos.BookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("book not found")
	}
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.ReportingCurrencyCode != nil {
		existing.ReportingCurrencyCode = req.ReportingCurrencyCode
	}
	if req.Status != nil {
		existing.Status = models.BookStatus(*req.Status)
	}
	if err := s.repos.BookRepo.Update(ctx, existing); err != nil {
		return nil, err
	}
	updated, err := s.repos.BookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return bookModelToDTO(updated), nil
}

func (s *bookService) DeleteBook(ctx context.Context, id string) error {
	return s.repos.BookRepo.Delete(ctx, id)
}
