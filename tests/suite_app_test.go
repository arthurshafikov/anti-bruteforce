package tests

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/thewolf27/anti-bruteforce/internal/app"
	"github.com/thewolf27/anti-bruteforce/internal/bucket"
	"github.com/thewolf27/anti-bruteforce/internal/models"
	"github.com/thewolf27/anti-bruteforce/internal/server/http"
	"github.com/thewolf27/anti-bruteforce/internal/storage"
	"github.com/thewolf27/anti-bruteforce/pkg/logger"
)

type AppSuite struct {
	suite.Suite
	cancelContext context.CancelFunc
	App           *app.App
	ServerEngine  *gin.Engine
}

func (appS *AppSuite) SetupSuite() {
	ctx, cancel := context.WithCancel(context.Background())
	appS.cancelContext = cancel

	logger := logger.NewLogger("DEBUG")
	storage := storage.NewStorage(os.Getenv("DSN"))
	storage.Connect(ctx)

	bucket := bucket.NewLeakyBucket(ctx, models.AuthorizeLimits{
		LimitAttemptsForLogin:    int64(limitAttemptsForLogin),
		LimitAttemptsForPassword: int64(limitAttemptsForPassword),
		LimitAttemptsForIP:       int64(limitAttemptsForIP),
	})

	appS.App = app.NewApp(ctx, logger, storage, bucket)
	handler := http.NewHandler(appS.App)

	server := http.NewServer(":8999", handler)
	server.InitRoutes()
	appS.ServerEngine = server.Engine
}

func (appS *AppSuite) TearDownTest() {
	appS.App.Storage.ResetDatabase()
	appS.App.ResetBucket()
}

func (appS *AppSuite) TearDownSuite() {
	appS.cancelContext()
}
