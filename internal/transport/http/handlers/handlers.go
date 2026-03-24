package handlers

import "qb-accounting/internal/services"

type Handlers struct {
	BookHandler             *BookHandler
	FiscalYearHandler       *FiscalYearHandler
	AccountingPeriodHandler *AccountingPeriodHandler
	AccountGroupHandler     *AccountGroupHandler
	AccountHandler          *AccountHandler
	VoucherSequenceHandler  *VoucherSequenceHandler
	JournalBatchHandler     *JournalBatchHandler
	JournalHandler          *JournalHandler
	PostingRuleHandler      *PostingRuleHandler
	PostingRequestHandler   *PostingRequestHandler
	OpenItemHandler         *OpenItemHandler
}

func NewHandlers(svc *services.Services) *Handlers {
	return &Handlers{
		BookHandler:             NewBookHandler(svc.BookService),
		FiscalYearHandler:       NewFiscalYearHandler(svc.FiscalYearService),
		AccountingPeriodHandler: NewAccountingPeriodHandler(svc.AccountingPeriodService),
		AccountGroupHandler:     NewAccountGroupHandler(svc.AccountGroupService),
		AccountHandler:          NewAccountHandler(svc.AccountService),
		VoucherSequenceHandler:  NewVoucherSequenceHandler(svc.VoucherSequenceService),
		JournalBatchHandler:     NewJournalBatchHandler(svc.JournalBatchService),
		JournalHandler:          NewJournalHandler(svc.JournalService),
		PostingRuleHandler:      NewPostingRuleHandler(svc.PostingRuleService),
		PostingRequestHandler:   NewPostingRequestHandler(svc.PostingRequestService),
		OpenItemHandler:         NewOpenItemHandler(svc.OpenItemService),
	}
}
