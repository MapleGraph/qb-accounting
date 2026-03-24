package remote

import (
	"context"
	"fmt"

	pbNotification "qb-accounting/internal/proto/notification"

	qbgrpc "github.com/MapleGraph/qb-core/v2/pkg/grpc"
)

const ServiceNameNotification = "notification"

// NotificationService defines the interface for notification service operations
type NotificationService interface {
	SendSMS(ctx context.Context, req *SendSMSRequest) (*SendSMSResponse, error)
	SendEmail(ctx context.Context, req *SendEmailRequest) (*SendEmailResponse, error)
}

// SendSMSRequest represents a request to send an SMS
type SendSMSRequest struct {
	To         string
	TemplateID string
	Variables  map[string]string
}

// SendSMSResponse represents the response from sending an SMS
type SendSMSResponse struct {
	MessageID string
	Success   bool
	Error     string
}

// SendEmailRequest represents a request to send an email
type SendEmailRequest struct {
	To           []string
	CC           []string
	BCC          []string
	Subject      string
	TemplateID   string
	Variables    map[string]string
	IsCustomMail bool
	MailContent  string
}

// SendEmailResponse represents the response from sending an email
type SendEmailResponse struct {
	MessageID string
	Success   bool
	Error     string
}

type notificationRepository struct {
	handler qbgrpc.ClientHandler
}

// NewNotificationRepository creates a new notification repository using a qb-core gRPC handler.
func NewNotificationRepository(handler qbgrpc.ClientHandler) NotificationService {
	return &notificationRepository{handler: handler}
}

func (r *notificationRepository) client(ctx context.Context) (pbNotification.NotificationServiceClient, error) {
	if r.handler == nil {
		return nil, fmt.Errorf("notification gRPC handler is not available")
	}
	conn, err := r.handler.GetConnectionWithRetry(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get notification gRPC connection: %w", err)
	}
	return pbNotification.NewNotificationServiceClient(conn), nil
}

// SendSMS sends an SMS via the notification service
func (r *notificationRepository) SendSMS(ctx context.Context, req *SendSMSRequest) (*SendSMSResponse, error) {
	client, err := r.client(ctx)
	if err != nil {
		return nil, err
	}

	grpcReq := &pbNotification.SendSMSRequest{
		To:         req.To,
		TemplateId: req.TemplateID,
		Variables:  req.Variables,
	}

	resp, err := client.SendSMS(ctx, grpcReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call SendSMS: %w", err)
	}

	return &SendSMSResponse{
		MessageID: resp.MessageId,
		Success:   resp.Success,
		Error:     resp.Error,
	}, nil
}

// SendEmail sends an email via the notification service
func (r *notificationRepository) SendEmail(ctx context.Context, req *SendEmailRequest) (*SendEmailResponse, error) {
	client, err := r.client(ctx)
	if err != nil {
		return nil, err
	}

	grpcReq := &pbNotification.SendEmailRequest{
		To:           req.To,
		Cc:           req.CC,
		Bcc:          req.BCC,
		Subject:      req.Subject,
		TemplateId:   req.TemplateID,
		Variables:    req.Variables,
		IsCustomMail: req.IsCustomMail,
		MailContent:  req.MailContent,
	}

	resp, err := client.SendEmail(ctx, grpcReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call SendEmail: %w", err)
	}

	return &SendEmailResponse{
		MessageID: resp.MessageId,
		Success:   resp.Success,
		Error:     resp.Error,
	}, nil
}
