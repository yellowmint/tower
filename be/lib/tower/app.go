package tower

import (
	"git.jetbrains.space/artdecoction/wt/tower/lib/config"
	"git.jetbrains.space/artdecoction/wt/tower/lib/logs"
	"go.uber.org/zap"
)

type App struct {
	Logger *zap.Logger
}

func NewTowerApp() App {
	logger := logs.NewLogger(config.Get().LoggerFormat)
	printStartInfo(logger)

	return App{
		logger,
	}
}

func CleanupApp(app App) {
	if app.Logger != nil {
		logs.SyncLogger(app.Logger)
	}
}

func printStartInfo(logger *zap.Logger) {
	logger.Info("Tower server starting", zap.String("version", config.Get().Version))
}
