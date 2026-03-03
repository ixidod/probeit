package prober

import "time"

type Result struct {
	StatusCode int
	Latency    time.Duration
	Err        error
}
