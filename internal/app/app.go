package app

import (
	"context"
	"flag"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"maxblog-me-template/internal/conf"
	"maxblog-me-template/internal/core"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type options struct {
	ConfigDir  string
	ConfigFile string
}

type Option func(*options)

func SetConfigDir(configDir string) Option {
	return func(opts *options) {
		opts.ConfigDir = configDir
	}
}

func SetConfigFile(configFile string) Option {
	return func(opts *options) {
		opts.ConfigFile = configFile
	}
}

func InitConfig(opts *options) {
	cfg := conf.GetInstanceOfConfig()
	cfg.Load(opts.ConfigDir, opts.ConfigFile)
	logger.WithFields(logger.Fields{
		"path": opts.ConfigDir + "/" + opts.ConfigFile,
	}).Info(core.Config_File_Load_Succeeded)
	core.SetUpstreamAddr(cfg.Upstream.MaxblogFETemplate.Host, cfg.Upstream.MaxblogFETemplate.Port)
	core.SetDownstreamAddr(cfg.Downstream.MaxblogBETemplate.Host, cfg.Downstream.MaxblogBETemplate.Port)
}

func InitServer(ctx context.Context, handler http.Handler) func() {
	cfg := conf.GetInstanceOfConfig()
	host := flag.String("host", cfg.Server.Host, "Enter host")
	port := flag.Int("port", cfg.Server.Port, "Enter port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	go func() {
		logger.WithContext(ctx).Infof("Server is running at %s", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(core.Server_Serve_Failed, err)
		}
	}()
	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.Server.ShutdownTimeout))
		defer cancel()
		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.WithContext(ctx).Errorf(err.Error())
		}
		logger.Info(core.Server_Stoped)
	}
}

func Init(ctx context.Context, opts ...Option) func() {
	// initialising options
	options := options{}
	for _, opt := range opts {
		opt(&options)
	}
	// init conf
	InitConfig(&options)
	// init injector
	injector, _ := InitInjector()
	// init server
	serverClean := InitServer(ctx, injector.Engine)
	return func() {
		serverClean()
	}
}

func Launch(ctx context.Context, opts ...Option) {
	logger.Info(core.Server_Launch_Start)
	clean := Init(ctx, opts...)
	cfg := conf.GetInstanceOfConfig()
	logger.WithFields(logger.Fields{
		"app_name": cfg.App.AppName,
		"version":  cfg.App.Version,
		"pid":      os.Getpid(),
		"host":     cfg.Server.Host,
		"port":     cfg.Server.Port,
	}).Info(core.Server_Started)
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
LOOP:
	for {
		sig := <-sc
		logger.WithContext(ctx).Infof("%s [%s]", core.Server_Interrupt_Received, sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break LOOP
		case syscall.SIGHUP:
		default:
			break LOOP
		}
	}
	defer logger.WithContext(ctx).Infof(core.Server_Shutting_Down)
	defer time.Sleep(time.Second)
	defer os.Exit(state)
	defer clean()
}
