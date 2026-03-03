package prober

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type HTTPProber struct {
	target string
	client *http.Client
}

// guardrail as usual ..
var _ Prober = (*HTTPProber)(nil)

// constructor
func NewHTTPProber(target string, client *http.Client) *HTTPProber {
	return &HTTPProber{target: target, client: client}
}

// implementation
func (h *HTTPProber) Target() string {
	return h.target
}
func (h *HTTPProber) Probe(ctx context.Context) Result {
	// time to go
	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, h.target, nil)
	if err != nil {
		return Result{
			Err: fmt.Errorf("building request: %w", err),
		}
	}
	res, err := h.client.Do(req)
	if err != nil {
		return Result{
			Latency: time.Since(start),
			Err:     fmt.Errorf("making request: %w", err),
		}
	}

	return Result{
		StatusCode: res.StatusCode,
		Latency:    time.Since(start),
	}
}
