package main

import (
	"os"

	"github.com/GrishaSkurikhin/EffectiveMobileTesting/internal/config"
	"github.com/GrishaSkurikhin/EffectiveMobileTesting/internal/lib/logger/slogpretty"
	"golang.org/x/exp/slog"
)

const (
	configPath = "config/.env"

	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func init() {
	config.MustLoad(configPath)
}

func main() {
	cfg := config.New()

	log := setupLogger(cfg.Env)
	log.Info(
		"starting app",
		slog.String("env", cfg.Env),
		slog.String("version", "1"),
	)
	log.Debug("debug messages are enabled")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
