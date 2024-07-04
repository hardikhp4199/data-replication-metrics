package metrics

import (
	"datareplication_metricsexporter/config"
	"datareplication_metricsexporter/core"
	"datareplication_metricsexporter/server"
	"datareplication_metricsexporter/storage/kafka"
	"datareplication_metricsexporter/storage/logging"
	"datareplication_metricsexporter/util/common"
	"encoding/json"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	topic         = config.GetString("Kafka.Topics.SourceTopic")
	consumerGroup = config.GetString("Kafka.ConsumerGroup")
	batch         = config.GetString("KafkaBatchSize")
)

func StartApplication() {
	// start promethues listen
	go server.ConnectPrometheus()
	// get prometheus object
	metrics := common.GetPrometheusObjects()
	// process of producing metrics
	processMetrics(metrics)
}

func processMetrics(metrics map[string]*prometheus.HistogramVec) {
	//  consume all messages from kafka
	metricsList, err_consume := consumeMessages()
	if err_consume != nil {
		logging.DoLoggingLevelBasedLogs(logging.Error, "", err_consume)
	} else {
		for _, v := range metricsList {
			metrics["RtConsume"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.RtConsume))
			metrics["CountBatch"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.CountBatch))
			metrics["CountConsume"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.CountConsume))
			metrics["CountUnique"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.CountUnique))
			metrics["CountUpdate"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.CountUpdate))
			metrics["CountDelete"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.CountDelete))
			metrics["CountSkip"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.CountSkip))
			metrics["RtCbTotal"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.RtCbTotal))
			metrics["RtUpdateTotal"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.RtUpdateTotal))
			metrics["RtDeleteTotal"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.RtDeleteTotal))
			metrics["RtCommitOffset"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.RtCommitOffset))
			metrics["RtTotal"].WithLabelValues(v.Type, v.Topic).Observe(float64(v.RtTotal))

			jsonData, err := json.Marshal(&v)

			if err != nil {
				logging.DoLoggingLevelBasedLogs(logging.Error, "", logging.EnrichErrorWithStackTrace(err))
			} else {
				logging.DoLoggingLevelBasedLogs(logging.Info, "metrics: "+string(jsonData), nil)
			}

		}
	}
	// commit kafka messages
	kafka.CommitOffset(topic, consumerGroup)
	processMetrics(metrics)
}

// for consume messages that return metrics list
func consumeMessages() (metricsList []core.MetricsObservation, kafkaError error) {
	batch, _ := strconv.Atoi(batch)
	result, err_consume := kafka.Consume(topic, consumerGroup, batch)
	if err_consume != nil {
		kafkaError = logging.EnrichErrorWithStackTrace(err_consume)
	} else {
		for _, msg := range result.Messages {
			var metricsObservation core.MetricsObservation
			err := json.Unmarshal(msg.Value, &metricsObservation)
			if err != nil {
				kafkaError = logging.EnrichErrorWithStackTrace(err)
			}
			// listing of messages
			metricsList = append(metricsList, metricsObservation)
		}
	}
	return metricsList, kafkaError
}
