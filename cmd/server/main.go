package main

import (
	"context"
	"example.com/open-gateway/internal/app"
	"example.com/open-gateway/internal/config"
	httpx "example.com/open-gateway/internal/delivery/http"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config.yaml")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	application := app.New(cfg)
	router := httpx.NewRouter(application.Logger)

	server := &http.Server{
		Addr:              cfg.Server.Addr,
		Handler:           router,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
		IdleTimeout:       cfg.Server.IdleTimeout,
	}

	go func() {
		http.ListenAndServe("127.0.0.1:6060", nil)
		//application.Logger.Info("listening", "addr", cfg.Server.Addr)
		//if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		//	log.Fatalf("server: %v", err)
		//}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	application.Logger.Info("shutting down")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown: %v", err)
	}
}