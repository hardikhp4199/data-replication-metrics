package common

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

//  promethues histogram object init
func GetPrometheusObjects() map[string]*prometheus.HistogramVec {
	metricsObj := make(map[string]*prometheus.HistogramVec)
	// consume batch  time
	consumeTime := getMetricsObject("dr_rt_consume", "This is drs metrics of consume batch time", []float64{10, 50, 100, 500}) // prometheus.DefBuckets = .005, .01, .025, .05, .1, .25, .5, 1, 2.5
	metricsObj["RtConsume"] = consumeTime

	// count of batch process
	countBatch := getMetricsObject("dr_count_batch", "This is drs metrics of count of batch ", []float64{100})
	metricsObj["CountBatch"] = countBatch

	//  consume messages metrics
	countConsume := getMetricsObject("dr_count_consume", "This is drs metrics of count of comsume batch", []float64{100})
	metricsObj["CountConsume"] = countConsume

	//  unique messages metrics
	countUnique := getMetricsObject("dr_count_unique", "This is drs metrics of count of unique batch ", []float64{100})
	metricsObj["CountUnique"] = countUnique

	// total count updated documents
	countUpdate := getMetricsObject("dr_count_update", "This is drs metrics of  count of delete messages", []float64{100})
	metricsObj["CountUpdate"] = countUpdate

	// total count deleted documents
	countDelete := getMetricsObject("dr_count_delete", "This is drs metrics of  count of delete messages", []float64{100})
	metricsObj["CountDelete"] = countDelete

	// total count skiped documents
	countSkip := getMetricsObject("dr_count_skip", "This is drs metrics ofcount of skip messages", []float64{100})
	metricsObj["CountSkip"] = countSkip

	//  total couchbase time
	rtCbTotalTime := getMetricsObject("dr_rt_cb_total", "This is drs metrics of  couchbase time ", []float64{2500, 5000, 7500, 10000, 15000, 20000})
	metricsObj["RtCbTotal"] = rtCbTotalTime

	//  total update time
	rtUpdate := getMetricsObject("dr_rt_update_total", "This is drs metrics of  couchbase update time ", []float64{2500, 5000, 7500, 10000, 15000, 20000})
	metricsObj["RtUpdateTotal"] = rtUpdate

	// total delete time
	rtDelete := getMetricsObject("dr_rt_delete_total", "This is drs metrics of  couchbase delete time ", []float64{2500, 5000, 7500, 10000, 15000, 20000})
	metricsObj["RtDeleteTotal"] = rtDelete

	//  commit time
	rtCommitOffset := getMetricsObject("dr_rt_commitOffset", "This is drs metrics of  commit offset time ", []float64{10, 50, 100})
	metricsObj["RtCommitOffset"] = rtCommitOffset

	//  entire batch time
	rtTotal := getMetricsObject("dr_rt_total", "This is drs metrics of  entire procees time ", []float64{500, 1000, 5000, 10000})
	metricsObj["RtTotal"] = rtTotal

	// unused metrics will be removed
	unregisterMetrics()
	return metricsObj
}

//  remove unused metrics
func unregisterMetrics() {
	prometheus.Unregister(collectors.NewGoCollector())
	prometheus.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
}

// return prometheus histogram object
func getMetricsObject(metricsName string, metricsHelpMessage string, bucketSize []float64) *prometheus.HistogramVec {
	metrics := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metricsName,
			Help:    metricsHelpMessage,
			Buckets: bucketSize,
		}, []string{"type", "topic"},
	)
	prometheus.MustRegister(metrics)
	return metrics
}
