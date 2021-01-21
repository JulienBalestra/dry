package ticknow

import (
	"context"
	"time"
)

// TickNow is a convenient wrapper over time.TickNow to immediately tick
type TickNow struct {
	C      chan time.Time
	ctx    context.Context
	cancel context.CancelFunc
}

type WithContext struct {
	C chan time.Time
}

func NewTickNow(duration time.Duration) *TickNow {
	ctx, cancel := context.WithCancel(context.TODO())
	t := &TickNow{
		C:      make(chan time.Time, 1),
		ctx:    ctx,
		cancel: cancel,
	}
	go func() {
		ticker := time.NewTicker(duration)
		t.C <- time.Now()
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return

			case tick := <-ticker.C:
				t.C <- tick
			}
		}
	}()
	return t
}

func (t *TickNow) Stop() {
	t.cancel()
}

func NewTickNowWithContext(ctx context.Context, duration time.Duration) *WithContext {
	t := &WithContext{
		C: make(chan time.Time, 1),
	}
	go func() {
		ticker := time.NewTicker(duration)
		t.C <- time.Now()
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return

			case tick := <-ticker.C:
				t.C <- tick
			}
		}
	}()
	return t
}
