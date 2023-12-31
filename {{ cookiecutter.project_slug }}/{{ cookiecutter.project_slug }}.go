package {{ cookiecutter.project_slug }}

import (
	"fmt"
	"log/slog"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	LogFormat string `long:"log-format" choice:"text" choice:"json" default:"text" description:"Log format"`
	Verbose   []bool `short:"v" long:"verbose" description:"Show verbose debug information, each -v bumps log level"`
	logLevel  slog.Level
}

func Execute() int {
	if err := parseFlags(); err != nil {
		slog.Error("error parsing flags", "error", err)
		return 1
	}

	if err := setLogLevel(); err != nil {
		slog.Error("error setting log level", "error", err)
		return 1
	}

	if err := setupLogger(); err != nil {
		slog.Error("error setting up logger", "error", err)
		return 1
	}

	if err := run(); err != nil {
		slog.Error("run failed", "error", err)
		return 1
	}

	return 0
}

func parseFlags() error {
	_, err := flags.Parse(&opts)
	return fmt.Errorf("parse flags failed: %w", err)
}

func run() error {
	slog.Debug("Debug", "currrent level", opts.logLevel)
	slog.Info("Info", "currrent level", opts.logLevel)
	slog.Warn("Warn", "currrent level", opts.logLevel)
	slog.Error("Error", "currrent level", opts.logLevel)

	return nil
}
