package tower

import (
	"git.jetbrains.space/artdecoction/gt/dtrace/pkg/dtrace"
	"git.jetbrains.space/artdecoction/wt/tower/lib/config"
	"git.jetbrains.space/artdecoction/wt/tower/lib/fbase"
	"git.jetbrains.space/artdecoction/wt/tower/lib/logs"
	"git.jetbrains.space/artdecoction/wt/tower/lib/validate"
	"go.uber.org/zap"
)

type App struct {
	Logger          *zap.Logger
	FirebaseClients *fbase.Clients
	Tracer          *dtrace.Tracer
}

func NewTowerApp() *App {
	logger, err := logs.NewLogger(config.Get().LoggerFormat)
	if err != nil {
		panic(err)
	}

	printStartInfo(logger)

	tracer, err := dtrace.NewTracer(config.Get().TracerEnabled, config.Get().GcpProjectId)
	if err != nil {
		panic(err)
	}

	err = fbase.ConfigureEmulator(config.Get().FirebaseEmulator)
	if err != nil {
		panic(err)
	}

	firebaseClients, err := fbase.NewClients(config.Get().GcpProjectId)
	if err != nil {
		panic(err)
	}

	validate.Init()

	return &App{
		logger,
		firebaseClients,
		tracer,
	}
}

func CleanupApp(app *App) {
	if app.Logger != nil {
		_ = logs.SyncLogger(app.Logger)
	}

	if app.Tracer != nil {
		_ = app.Tracer.Cleanup()
	}

	if app.FirebaseClients != nil {
		_ = app.FirebaseClients.Cleanup()
	}
}

func printStartInfo(logger *zap.Logger) {
	logger.Info("Tower server starting", zap.String("version", config.Get().Version))
}
