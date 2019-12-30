package main

import (
	"flag"
	"github.com/balerter/balerter/internal/config"
	scriptsManager "github.com/balerter/balerter/internal/script/manager"
	"go.uber.org/zap"
	"log"
	"os"
)

var (
	version = "undefined"
)

func main() {
	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Printf("error init zap logger, %v", err)
		os.Exit(1)
	}

	logger.Info("balerter start", zap.String("version", version))

	cfg := config.New()
	if err := cfg.Init(); err != nil {
		logger.Error("error init config", zap.Error(err))
		os.Exit(1)
	}
	if err := cfg.Validate(); err != nil {
		logger.Error("error validate config", zap.Error(err))
		os.Exit(1)
	}
	logger.Debug("loaded configuration", zap.Any("config", cfg))

	logger.Info("init scripts manager")
	scriptsMgr := scriptsManager.New()
	if err := scriptsMgr.Init(cfg.Scripts.Sources); err != nil {
		logger.Error("error init scripts manager", zap.Error(err))
		os.Exit(1)
	}

	//

	logger.Info("terminate")
}
