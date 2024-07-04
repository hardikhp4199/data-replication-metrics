package server

import (
	"datareplication_metricsexporter/config"
	"datareplication_metricsexporter/storage/logging"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var port = config.GetInt("HttpServer.Port")

// prometheus server start
func ConnectPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	httpPort := ":" + strconv.Itoa(port)
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		logging.DoLoggingLevelBasedLogs(logging.Error, "", logging.EnrichErrorWithStackTrace(err))
	}
}
