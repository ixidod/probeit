// Package cmd well well well
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ixidod/probeit/internal/prober"
)

func Run(target string, interval time.Duration, timeout time.Duration) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	// watch for ctrl-c or SIGTERM
	// when signal arives cancel the context
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit
		fmt.Println("\n stopping......")
		cancel()
	}()
	p := prober.NewHTTPProber(target, &http.Client{})
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			probeCtx, probeCancel := context.WithTimeout(ctx, timeout)
			result := p.Probe(probeCtx)
			probeCancel() //do not wait for defer
			printResult(p.Target(), result)
		case <-ctx.Done():
			return
		}
	}

}

func printResult(target string, r prober.Result) {
	if r.Err != nil {
		fmt.Fprintf(os.Stderr, "ERR   -      %s - %v\n", target, r.Err)
		return
	}
	fmt.Printf("%d %s %s\n", r.StatusCode, formatLatency(r.Latency), target)
}

func formatLatency(d time.Duration) string {
	return fmt.Sprintf("%4dms", d.Milliseconds())
}
