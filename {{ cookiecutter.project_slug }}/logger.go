package {{ cookiecutter.project_slug }}

import (
	"log/slog"
	"os"

	"github.com/taylormonacelli/littlecow"
)

func getLogger(logLevel slog.Level, logFormat string) (*slog.Logger, error) {
	opts := littlecow.NewHandlerOptions(logLevel, littlecow.RemoveTimestampAndTruncateSource)

	var handler slog.Handler
	handler = slog.NewTextHandler(os.Stderr, opts)
	if logFormat == "json" {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	}

	return slog.New(handler), nil
}
