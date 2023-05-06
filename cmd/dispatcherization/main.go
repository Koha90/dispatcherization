// Package main ...
package main

import "github.com/koha90/dispatcherization/pkg/logging"

func main() {
	logger := logging.GetLogger()
	logger.Info("logger implements")
}
