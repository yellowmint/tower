package tower

import (
	"git.jetbrains.space/artdecoction/wt/tower/lib/config"
	"git.jetbrains.space/artdecoction/wt/tower/lib/fbase"
	"git.jetbrains.space/artdecoction/wt/tower/lib/logs"
	"go.uber.org/zap"
)

type App struct {
	Logger          *zap.Logger
	FirebaseClients *fbase.Clients
}

func NewTowerApp() *App {
	logger, err := logs.NewLogger(config.Get().LoggerFormat)
	if err != nil {
		panic(err)
	}

	printStartInfo(logger)

	err = fbase.ConfigureEmulator(config.Get().FirebaseEmulator)
	if err != nil {
		panic(err)
	}

	firebaseClients, err := fbase.NewClients(config.Get().GcpProjectId)
	if err != nil {
		panic(err)
	}

	return &App{
		logger,
		firebaseClients,
	}
}

func CleanupApp(app *App) {
	if app.Logger != nil {
		_ = logs.SyncLogger(app.Logger)
	}

	if app.FirebaseClients != nil {
		_ = fbase.Cleanup(app.FirebaseClients)
	}
}

func printStartInfo(logger *zap.Logger) {
	logger.Info("Tower server starting", zap.String("version", config.Get().Version))
}
