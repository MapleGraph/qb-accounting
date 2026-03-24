package services

import (
	"qb-accounting/internal/repository"
)

type Services struct {
	BookService             BookService
	FiscalYearService       FiscalYearService
	AccountingPeriodService AccountingPeriodService
	AccountGroupService     AccountGroupService
	AccountService          AccountService
	VoucherSequenceService  VoucherSequenceService
	JournalBatchService     JournalBatchService
	JournalService          JournalService
	PostingRuleService      PostingRuleService
	PostingRequestService   PostingRequestService
	OpenItemService         OpenItemService
	AccountingGRPCService   AccountingGRPCService
}

func NewServices(repos *repository.RepositoryContainer) *Services {
	grpcSvc := NewAccountingGRPCService(repos)
	return &Services{
		BookService:             NewBookService(repos),
		FiscalYearService:       NewFiscalYearService(repos),
		AccountingPeriodService: NewAccountingPeriodService(repos),
		AccountGroupService:     NewAccountGroupService(repos),
		AccountService:          NewAccountService(repos),
		VoucherSequenceService:  NewVoucherSequenceService(repos),
		JournalBatchService:     NewJournalBatchService(repos),
		JournalService:          NewJournalService(repos),
		PostingRuleService:      NewPostingRuleService(repos),
		PostingRequestService:   NewPostingRequestService(repos),
		OpenItemService:         NewOpenItemService(repos),
		AccountingGRPCService:   grpcSvc,
	}
}
