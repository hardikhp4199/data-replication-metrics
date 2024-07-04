package core

import "time"

type MetricsObservation struct {
	RtConsume      int64
	CountBatch     int
	CountConsume   int
	CountUnique    int
	CountUpdate    int
	CountDelete    int
	CountSkip      int
	RtCbTotal      int64
	RtUpdateTotal  int64
	RtDeleteTotal  int64
	RtCommitOffset int64
	RtTotal        int64
	Topic          string
	Type           string
}

var LastActiveTime time.Time
