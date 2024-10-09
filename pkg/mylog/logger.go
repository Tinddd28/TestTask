package mylog

import (
	"github.com/Tinddd28/TestTask/pkg/mylog/slogpretty"
	"log/slog"
	"os"
)

func SetupLogger() *slog.Logger {
	logger := SetupPrettyLogger()
	return logger
}

func SetupPrettyLogger() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
