package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"qb-accounting/internal/config"
	"qb-accounting/internal/repository"
	postgres_repo "qb-accounting/internal/repository/postgres"
	"qb-accounting/internal/repository/remote"
	"qb-accounting/internal/services"
	"qb-accounting/internal/transport/http/handlers"
	"qb-accounting/internal/transport/http/routes"

	qbinfra "github.com/MapleGraph/qb-core/v2/pkg/infrastructure"

	_ "qb-accounting/docs"
)

// @title           QB Accounting Service API
// @version         1.0.0
// @description     Enterprise accounting module for QueueBuster ERP
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.queuebuster.co/support
// @contact.email  support@queuebuster.co

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8086
// @BasePath  /api/public/accounting/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.LoadConfig()

	qbCfg := cfg.ToQBCoreConfig()
	qbInfra, err := qbinfra.NewInfrastructureWithOptions(*qbCfg, qbinfra.InfrastructureOptions{
		InitializeDatabases:  true,
		InitializeRedis:      true,
		InitializeClickHouse: false,
		InitializeGRPC:       true,
		DatabasesRequired: map[string]bool{
			"main": true,
		},
	})
	if err != nil {
		log.Printf("failed to initialize qb-core infrastructure: %v", err)
		os.Exit(1)
	}
	defer func() {
		if closeErr := qbInfra.Close(); closeErr != nil {
			log.Printf("error closing infrastructure: %v", closeErr)
		}
	}()

	db, err := qbInfra.DBManager.Handler("main")
	if err != nil {
		log.Printf("failed to get main database handler from qb-core infrastructure: %v", err)
		os.Exit(1)
	}

	var setupService remote.SetupService
	var employeeService remote.EmployeeService
	var notificationService remote.NotificationService
	var catalogueService remote.CatalogueService

	if qbInfra.GRPC != nil {
		if setupHandler, err := qbInfra.GRPC.Get(remote.ServiceNameSetup); err == nil {
			setupService = remote.NewSetupRepository(setupHandler)
			log.Printf("setup gRPC service initialized")
		} else if cfg.Server.GrpcServices.SetupService.Enabled {
			log.Printf("setup gRPC service not available: %v", err)
		}

		if employeeHandler, err := qbInfra.GRPC.Get(remote.ServiceNameEmployee); err == nil {
			employeeService = remote.NewEmployeeRepository(employeeHandler)
			log.Printf("employee gRPC service initialized")
		} else if cfg.Server.GrpcServices.EmployeeService.Enabled {
			log.Printf("employee gRPC service not available: %v", err)
		}

		if notificationHandler, err := qbInfra.GRPC.Get(remote.ServiceNameNotification); err == nil {
			notificationService = remote.NewNotificationRepository(notificationHandler)
			log.Printf("notification gRPC service initialized")
		} else if cfg.Server.GrpcServices.NotificationService.Enabled {
			log.Printf("notification gRPC service not available: %v", err)
		}

		if catalogueHandler, err := qbInfra.GRPC.Get(remote.ServiceNameCatalogue); err == nil {
			catalogueService = remote.NewCatalogueRepository(catalogueHandler)
			log.Printf("catalogue gRPC service initialized")
		} else if cfg.Server.GrpcServices.CatalogueService.Enabled {
			log.Printf("catalogue gRPC service not available: %v", err)
		}
	} else {
		log.Printf("gRPC registry not initialized, remote services unavailable")
	}

	bookRepo, err := postgres_repo.NewBookRepository(db)
	if err != nil {
		log.Printf("failed to create book repository: %v", err)
		os.Exit(1)
	}
	fiscalYearRepo, err := postgres_repo.NewFiscalYearRepository(db)
	if err != nil {
		log.Printf("failed to create fiscal year repository: %v", err)
		os.Exit(1)
	}
	accountingPeriodRepo, err := postgres_repo.NewAccountingPeriodRepository(db)
	if err != nil {
		log.Printf("failed to create accounting period repository: %v", err)
		os.Exit(1)
	}
	accountGroupRepo, err := postgres_repo.NewAccountGroupRepository(db)
	if err != nil {
		log.Printf("failed to create account group repository: %v", err)
		os.Exit(1)
	}
	accountRepo, err := postgres_repo.NewAccountRepository(db)
	if err != nil {
		log.Printf("failed to create account repository: %v", err)
		os.Exit(1)
	}
	voucherSequenceRepo, err := postgres_repo.NewVoucherSequenceRepository(db)
	if err != nil {
		log.Printf("failed to create voucher sequence repository: %v", err)
		os.Exit(1)
	}
	journalBatchRepo, err := postgres_repo.NewJournalBatchRepository(db)
	if err != nil {
		log.Printf("failed to create journal batch repository: %v", err)
		os.Exit(1)
	}
	journalRepo, err := postgres_repo.NewJournalRepository(db)
	if err != nil {
		log.Printf("failed to create journal repository: %v", err)
		os.Exit(1)
	}
	journalLineRepo, err := postgres_repo.NewJournalLineRepository(db)
	if err != nil {
		log.Printf("failed to create journal line repository: %v", err)
		os.Exit(1)
	}
	postingRuleVersionRepo, err := postgres_repo.NewPostingRuleVersionRepository(db)
	if err != nil {
		log.Printf("failed to create posting rule version repository: %v", err)
		os.Exit(1)
	}
	postingRequestRepo, err := postgres_repo.NewPostingRequestRepository(db)
	if err != nil {
		log.Printf("failed to create posting request repository: %v", err)
		os.Exit(1)
	}
	postingRequestSnapshotRepo, err := postgres_repo.NewPostingRequestSnapshotRepository(db)
	if err != nil {
		log.Printf("failed to create posting request snapshot repository: %v", err)
		os.Exit(1)
	}
	journalSourceLinkRepo, err := postgres_repo.NewJournalSourceLinkRepository(db)
	if err != nil {
		log.Printf("failed to create journal source link repository: %v", err)
		os.Exit(1)
	}
	openItemRepo, err := postgres_repo.NewOpenItemRepository(db)
	if err != nil {
		log.Printf("failed to create open item repository: %v", err)
		os.Exit(1)
	}
	openItemAllocationRepo, err := postgres_repo.NewOpenItemAllocationRepository(db)
	if err != nil {
		log.Printf("failed to create open item allocation repository: %v", err)
		os.Exit(1)
	}
	openItemAdjustmentRepo, err := postgres_repo.NewOpenItemAdjustmentRepository(db)
	if err != nil {
		log.Printf("failed to create open item adjustment repository: %v", err)
		os.Exit(1)
	}

	repoContainer := repository.NewRepositoryContainer(
		bookRepo,
		fiscalYearRepo,
		accountingPeriodRepo,
		accountGroupRepo,
		accountRepo,
		voucherSequenceRepo,
		journalBatchRepo,
		journalRepo,
		journalLineRepo,
		postingRuleVersionRepo,
		postingRequestRepo,
		postingRequestSnapshotRepo,
		journalSourceLinkRepo,
		openItemRepo,
		openItemAllocationRepo,
		openItemAdjustmentRepo,
		setupService,
		employeeService,
		notificationService,
		catalogueService,
		db,
	)

	svc := services.NewServices(repoContainer)
	hdlrs := handlers.NewHandlers(svc)

	routerDeps := routes.RouterDependencies{
		Config:       cfg,
		Handlers:     hdlrs,
		HealthProber: qbInfra,
	}
	router := routes.SetupRoutes(routerDeps)

	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Printf("HTTP server starting on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Printf("signal received, shutting down...")
	case err := <-errCh:
		log.Printf("server fatal error: %v", err)
		os.Exit(1)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP graceful shutdown failed: %v", err)
	}

	log.Printf("application shutdown complete")
}
