package run

import (
	"context"
	"fmt"
	"log/slog"

	api "github.com/WtoJglg4/prizma-help/api/service"
	"github.com/WtoJglg4/prizma-help/internal/config"
	"github.com/WtoJglg4/prizma-help/internal/httpsrv"
	"github.com/WtoJglg4/prizma-help/internal/server"
)

type Options struct {
	Cfg    config.Config
	Logger *slog.Logger
}

func Run(ctx context.Context, opts Options) error {
	cfg := opts.Cfg
	logger := opts.Logger
	httpSrv := httpsrv.New(httpsrv.Options{
		Addr:   cfg.HTTPAddr,
		Logger: logger.With("logging-entity", "httpsrv"),
	})
	handler, err := api.NewServer(server.Server{})
	if err != nil {
		return fmt.Errorf("create api handler: %w", err)
	}
	httpSrv.Register("/get", handler)
	if err := httpSrv.Start(ctx); err != nil {
		return fmt.Errorf("httpsrv start: %w", err)
	}
	// Waiting for SIGINT or SIGTERM to stop the server
	<-ctx.Done()
	return httpSrv.Stop(ctx)
}
