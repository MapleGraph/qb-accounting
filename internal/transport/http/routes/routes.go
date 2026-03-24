package routes

import (
	"context"
	"net/http"

	"qb-accounting/internal/config"
	"qb-accounting/internal/transport/http/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// HealthProber is used by health endpoints; routes do not depend on qb-core details.
type HealthProber interface {
	IsLive() bool
	IsReady(ctx context.Context) bool
	IsStarted() bool
}

// RouterDependencies holds all dependencies needed for routing.
type RouterDependencies struct {
	Config       *config.Config
	Handlers     *handlers.Handlers
	HealthProber HealthProber
}

// SetupRoutes configures all application routes.
func SetupRoutes(deps RouterDependencies) *gin.Engine {
	gin.SetMode(deps.Config.Server.Mode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Health probes
	r.GET("/health", func(c *gin.Context) { HealthCheckHandler(c, deps) })
	r.GET("/healthz", func(c *gin.Context) { LivenessHandler(c, deps) })
	r.GET("/readyz", func(c *gin.Context) { ReadinessHandler(c, deps) })
	r.GET("/startupz", func(c *gin.Context) { StartupHandler(c, deps) })

	// Swagger
	r.GET("/api/public/accounting/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		public := api.Group("/public/accounting")
		{
			v1 := public.Group("/v1")
			{
				deps.Handlers.BookHandler.RegisterRoutes(v1)
				deps.Handlers.FiscalYearHandler.RegisterRoutes(v1)
				deps.Handlers.AccountingPeriodHandler.RegisterRoutes(v1)
				deps.Handlers.AccountGroupHandler.RegisterRoutes(v1)
				deps.Handlers.AccountHandler.RegisterRoutes(v1)
				deps.Handlers.VoucherSequenceHandler.RegisterRoutes(v1)
				deps.Handlers.JournalBatchHandler.RegisterRoutes(v1)
				deps.Handlers.JournalHandler.RegisterRoutes(v1)
				deps.Handlers.PostingRuleHandler.RegisterRoutes(v1)
				deps.Handlers.PostingRequestHandler.RegisterRoutes(v1)
				deps.Handlers.OpenItemHandler.RegisterRoutes(v1)
			}
		}

		protected := api.Group("/accounting")
		{
			v1 := protected.Group("/v1")
			{
				deps.Handlers.BookHandler.RegisterRoutes(v1)
				deps.Handlers.FiscalYearHandler.RegisterRoutes(v1)
				deps.Handlers.AccountingPeriodHandler.RegisterRoutes(v1)
				deps.Handlers.AccountGroupHandler.RegisterRoutes(v1)
				deps.Handlers.AccountHandler.RegisterRoutes(v1)
				deps.Handlers.VoucherSequenceHandler.RegisterRoutes(v1)
				deps.Handlers.JournalBatchHandler.RegisterRoutes(v1)
				deps.Handlers.JournalHandler.RegisterRoutes(v1)
				deps.Handlers.PostingRuleHandler.RegisterRoutes(v1)
				deps.Handlers.PostingRequestHandler.RegisterRoutes(v1)
				deps.Handlers.OpenItemHandler.RegisterRoutes(v1)
			}
		}
	}

	return r
}

// HealthCheckHandler returns the service health status; 200 when ready (required components healthy).
func HealthCheckHandler(c *gin.Context, deps RouterDependencies) {
	if deps.HealthProber != nil && deps.HealthProber.IsReady(c.Request.Context()) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "qb-accounting",
			"version": "1.0.0",
		})
		return
	}
	c.JSON(http.StatusServiceUnavailable, gin.H{
		"status":  "unhealthy",
		"service": "qb-accounting",
		"version": "1.0.0",
	})
}

// LivenessHandler implements Kubernetes liveness: process is running.
func LivenessHandler(c *gin.Context, deps RouterDependencies) {
	if deps.HealthProber != nil && deps.HealthProber.IsLive() {
		c.Status(http.StatusOK)
		return
	}
	c.Status(http.StatusServiceUnavailable)
}

// ReadinessHandler implements Kubernetes readiness: required components healthy, can serve traffic.
func ReadinessHandler(c *gin.Context, deps RouterDependencies) {
	if deps.HealthProber != nil && deps.HealthProber.IsReady(c.Request.Context()) {
		c.Status(http.StatusOK)
		return
	}
	c.Status(http.StatusServiceUnavailable)
}

// StartupHandler implements Kubernetes startup: initialization complete.
func StartupHandler(c *gin.Context, deps RouterDependencies) {
	if deps.HealthProber != nil && deps.HealthProber.IsStarted() {
		c.Status(http.StatusOK)
		return
	}
	c.Status(http.StatusServiceUnavailable)
}
