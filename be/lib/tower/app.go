package tower

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4/auth"
	"git.jetbrains.space/artdecoction/wt/tower/lib/config"
	"git.jetbrains.space/artdecoction/wt/tower/lib/logs"
	"go.uber.org/zap"
)

type App struct {
	Logger          *zap.Logger
	AuthClient      *auth.Client
	FirestoreClient *firestore.Client
}

func NewTowerApp() *App {
	logger := logs.NewLogger(config.Get().LoggerFormat)
	printStartInfo(logger)

	return &App{
		logger,
		nil,
		nil,
	}
}

func CleanupApp(app *App) {
	if app.Logger != nil {
		logs.SyncLogger(app.Logger)
	}
}

func printStartInfo(logger *zap.Logger) {
	logger.Info("Tower server starting", zap.String("version", config.Get().Version))
}
