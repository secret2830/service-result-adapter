package main

import (
	"os"

	"github.com/smartcontractkit/chainlink/core/logger"
	"go.uber.org/zap/zapcore"
)

func init() {
	logger.SetLogger(logger.CreateProductionLogger("", false, zapcore.DebugLevel, false))
}

func main() {
	logger.Info("Starting service result adapter")

	listenAddr := os.Getenv("SRA_LISTEN_ADDR")

	RunWebServer(listenAddr, handle)
}
