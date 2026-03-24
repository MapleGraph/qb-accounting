package remote

import (
	"context"
	"fmt"

	pbEmployee "qb-accounting/internal/proto/employee"

	qbgrpc "github.com/MapleGraph/qb-core/v2/pkg/grpc"
)

const ServiceNameEmployee = "employee"

// EmployeeService defines the interface for employee service operations
type EmployeeService interface {
	GetUser(ctx context.Context, uuid string) (*User, error)
	AddUser(ctx context.Context, req *AddUserRequest) (*User, error)
}

// User represents a user from the employee service
type User struct {
	UUID           string
	OrganizationID string
	DepartmentID   string
	UserType       string
	FirstName      string
	LastName       string
	Email          string
	Mobile         string
	IsActive       int32
	Roles          []string
}

// AddUserRequest represents a request to add a new user
type AddUserRequest struct {
	OrganizationID string
	UserType       string
	FirstName      string
	LastName       string
	Email          string
	Mobile         string
	Password       string
	Role           string
}

type employeeRepository struct {
	handler qbgrpc.ClientHandler
}

// NewEmployeeRepository creates a new employee repository using a qb-core gRPC handler.
func NewEmployeeRepository(handler qbgrpc.ClientHandler) EmployeeService {
	return &employeeRepository{handler: handler}
}

func (r *employeeRepository) client(ctx context.Context) (pbEmployee.UserServiceClient, error) {
	if r.handler == nil {
		return nil, fmt.Errorf("employee gRPC handler is not available")
	}
	conn, err := r.handler.GetConnectionWithRetry(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee gRPC connection: %w", err)
	}
	return pbEmployee.NewUserServiceClient(conn), nil
}

// GetUser fetches a user by UUID from the employee service
func (r *employeeRepository) GetUser(ctx context.Context, uuid string) (*User, error) {
	client, err := r.client(ctx)
	if err != nil {
		return nil, err
	}

	req := &pbEmployee.GetUserRequest{
		Uuid: uuid,
	}

	resp, err := client.GetUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call GetUser: %w", err)
	}

	if resp.User == nil {
		return nil, nil
	}

	return &User{
		UUID:           resp.User.Uuid,
		OrganizationID: resp.User.OrganizationId,
		DepartmentID:   resp.User.DepartmentId,
		UserType:       resp.User.UserType,
		FirstName:      resp.User.FirstName,
		LastName:       resp.User.LastName,
		Email:          resp.User.Email,
		Mobile:         resp.User.Mobile,
		IsActive:       resp.User.IsActive,
		Roles:          resp.User.Roles,
	}, nil
}

// AddUser adds a new user via the employee service
func (r *employeeRepository) AddUser(ctx context.Context, req *AddUserRequest) (*User, error) {
	client, err := r.client(ctx)
	if err != nil {
		return nil, err
	}

	grpcReq := &pbEmployee.AddUserRequest{
		OrganizationId: req.OrganizationID,
		UserType:       req.UserType,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Mobile:         req.Mobile,
		Password:       req.Password,
		Role:           req.Role,
	}

	resp, err := client.AddUser(ctx, grpcReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call AddUser: %w", err)
	}

	if resp.User == nil {
		return nil, nil
	}

	return &User{
		UUID:           resp.User.Uuid,
		OrganizationID: resp.User.OrganizationId,
		DepartmentID:   resp.User.DepartmentId,
		UserType:       resp.User.UserType,
		FirstName:      resp.User.FirstName,
		LastName:       resp.User.LastName,
		Email:          resp.User.Email,
		Mobile:         resp.User.Mobile,
		IsActive:       resp.User.IsActive,
		Roles:          resp.User.Roles,
	}, nil
}
