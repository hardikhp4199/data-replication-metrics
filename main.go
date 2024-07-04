package main

import (
	"datareplication_metricsexporter/metrics"
	"datareplication_metricsexporter/storage/logging"
)

func main() {
	// starting application
	logging.DoLoggingLevelBasedLogs(logging.Info, "started metricsexporter", nil)
	metrics.StartApplication()
}
