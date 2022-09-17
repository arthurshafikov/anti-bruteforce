package tests

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/arthurshafikov/anti-bruteforce/internal/bucket"
	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/logger"
	"github.com/arthurshafikov/anti-bruteforce/internal/repository"
	"github.com/arthurshafikov/anti-bruteforce/internal/services"
	"github.com/arthurshafikov/anti-bruteforce/internal/transport/http"
	"github.com/arthurshafikov/anti-bruteforce/pkg/postgres"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"
)

type AppSuite struct {
	suite.Suite
	cancelContext context.CancelFunc
	Services      *services.Services
	DB            *sqlx.DB
	ServerEngine  *gin.Engine
}

func (appS *AppSuite) SetupSuite() {
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)
	appS.cancelContext = cancel

	logger := logger.NewLogger("DEBUG")

	resetBucketsTicker := time.NewTicker(time.Second * 60) // todo config
	group.Go(func() error {
		<-ctx.Done()
		resetBucketsTicker.Stop()

		return nil
	})
	bucket := bucket.NewLeakyBucket(resetBucketsTicker, core.AuthorizeLimits{
		LimitAttemptsForLogin:    int64(limitAttemptsForLogin),
		LimitAttemptsForPassword: int64(limitAttemptsForPassword),
		LimitAttemptsForIP:       int64(limitAttemptsForIP),
	})
	group.Go(func() error {
		bucket.Leak(ctx)

		return nil
	})

	appS.DB = postgres.NewSqlxDB(ctx, group, os.Getenv("DSN"))
	repository := repository.NewRepository(appS.DB)
	appS.Services = services.NewServices(&services.Dependencies{
		Logger:      logger,
		LeakyBucket: bucket,
		Repository:  repository,
	})
	handler := http.NewHandler(appS.Services)

	server := http.NewServer(handler)
	server.InitRoutes()
	appS.ServerEngine = server.Engine
}

func (appS *AppSuite) TearDownTest() {
	tables := []string{
		core.WhitelistIpsTable,
		core.BlacklistIpsTable,
	}
	_, err := appS.DB.Exec(fmt.Sprintf(`TRUNCATE TABLE %s`, strings.Join(tables, ", ")))
	require.NoError(appS.T(), err)
	appS.Services.Bucket.ResetBucket()
}

func (appS *AppSuite) TearDownSuite() {
	appS.cancelContext()
}
