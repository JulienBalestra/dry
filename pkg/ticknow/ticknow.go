package ticknow

import "time"

// TickNow is a convenient wrapper over time.TickNow to immediately tick
type TickNow struct {
	internalTicker *time.Ticker
	C              chan time.Time
}

// New creates and starts a ticker with an immediate tick
func NewTickNow(duration time.Duration) *TickNow {
	t := &TickNow{
		internalTicker: time.NewTicker(duration),
		C:              make(chan time.Time, 1),
	}
	// immediately tick
	t.C <- time.Now()
	go func() {
		for tick := range t.internalTicker.C {
			t.C <- tick
		}
	}()
	return t
}

// Stop the internal ticker and close the exposed chan
func (t *TickNow) Stop() {
	t.internalTicker.Stop()
	close(t.C)
}
