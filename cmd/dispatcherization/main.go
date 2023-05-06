// Package main ...
package main

import (
	"github.com/koha90/dispatcherization/internal/config"
	"github.com/koha90/dispatcherization/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	logger.Info("logger implements")
	logger.Info("config implements on port:", cfg.Listen.PORT)
}
