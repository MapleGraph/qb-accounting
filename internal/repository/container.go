package repository

import (
	qbpostgres "github.com/MapleGraph/qb-core/v2/pkg/postgres"

	"qb-accounting/internal/repository/remote"
)

// RepositoryContainer holds all repository instances and remote clients.
type RepositoryContainer struct {
	BookRepo                     BookRepository
	FiscalYearRepo               FiscalYearRepository
	AccountingPeriodRepo         AccountingPeriodRepository
	AccountGroupRepo             AccountGroupRepository
	AccountRepo                  AccountRepository
	VoucherSequenceRepo          VoucherSequenceRepository
	JournalBatchRepo             JournalBatchRepository
	JournalRepo                  JournalRepository
	JournalLineRepo              JournalLineRepository
	PostingRuleVersionRepo       PostingRuleVersionRepository
	PostingRequestRepo           PostingRequestRepository
	PostingRequestSnapshotRepo   PostingRequestSnapshotRepository
	JournalSourceLinkRepo        JournalSourceLinkRepository
	OpenItemRepo                 OpenItemRepository
	OpenItemAllocationRepo       OpenItemAllocationRepository
	OpenItemAdjustmentRepo       OpenItemAdjustmentRepository

	SetupService        remote.SetupService
	EmployeeService     remote.EmployeeService
	NotificationService remote.NotificationService
	CatalogueService    remote.CatalogueService

	DB qbpostgres.DBHandler
}

// NewRepositoryContainer creates a new repository container with all repositories.
func NewRepositoryContainer(
	bookRepo BookRepository,
	fiscalYearRepo FiscalYearRepository,
	accountingPeriodRepo AccountingPeriodRepository,
	accountGroupRepo AccountGroupRepository,
	accountRepo AccountRepository,
	voucherSequenceRepo VoucherSequenceRepository,
	journalBatchRepo JournalBatchRepository,
	journalRepo JournalRepository,
	journalLineRepo JournalLineRepository,
	postingRuleVersionRepo PostingRuleVersionRepository,
	postingRequestRepo PostingRequestRepository,
	postingRequestSnapshotRepo PostingRequestSnapshotRepository,
	journalSourceLinkRepo JournalSourceLinkRepository,
	openItemRepo OpenItemRepository,
	openItemAllocationRepo OpenItemAllocationRepository,
	openItemAdjustmentRepo OpenItemAdjustmentRepository,
	setupService remote.SetupService,
	employeeService remote.EmployeeService,
	notificationService remote.NotificationService,
	catalogueService remote.CatalogueService,
	db qbpostgres.DBHandler,
) *RepositoryContainer {
	return &RepositoryContainer{
		BookRepo:                   bookRepo,
		FiscalYearRepo:             fiscalYearRepo,
		AccountingPeriodRepo:       accountingPeriodRepo,
		AccountGroupRepo:           accountGroupRepo,
		AccountRepo:                accountRepo,
		VoucherSequenceRepo:        voucherSequenceRepo,
		JournalBatchRepo:           journalBatchRepo,
		JournalRepo:                journalRepo,
		JournalLineRepo:            journalLineRepo,
		PostingRuleVersionRepo:     postingRuleVersionRepo,
		PostingRequestRepo:         postingRequestRepo,
		PostingRequestSnapshotRepo: postingRequestSnapshotRepo,
		JournalSourceLinkRepo:      journalSourceLinkRepo,
		OpenItemRepo:               openItemRepo,
		OpenItemAllocationRepo:     openItemAllocationRepo,
		OpenItemAdjustmentRepo:     openItemAdjustmentRepo,
		SetupService:               setupService,
		EmployeeService:            employeeService,
		NotificationService:        notificationService,
		CatalogueService:           catalogueService,
		DB:                         db,
	}
}
