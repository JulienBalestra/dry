package ticknow

import (
	"context"
	"time"
)

// TickNow is a convenient wrapper over time.TickNow to immediately tick
type TickNow struct {
	C chan struct{}
}

// New creates and starts a ticker with an immediate tick
func NewTickNow(ctx context.Context, duration time.Duration) *TickNow {
	t := &TickNow{
		C: make(chan struct{}, 1),
	}
	t.C <- struct{}{}
	go func() {
		for {
			tick, cancel := context.WithTimeout(ctx, duration)
			select {
			case <-ctx.Done():
				cancel()
				close(t.C)
				return

			case <-tick.Done():
				cancel()
				t.C <- struct{}{}
				return
			}
		}
	}()
	return t
}
