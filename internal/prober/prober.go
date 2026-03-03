// Package prober comment well well well
package prober

import (
	"context"
	"time"
)

type Result struct {
	StatusCode int
	Latency    time.Duration
	Err        error
}

type Prober interface {
	Target() string
	Probe(ctx context.Context) Result
}
