package ticknow

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTickNow(t *testing.T) {
	const period = time.Millisecond * 100
	const factor = 3
	const timeout = period * factor
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	tick := NewTickNow(ctx, period)
	incr := 0
	for {
		select {
		case <-ctx.Done():
			cancel()
			assert.GreaterOrEqual(t, incr, factor)
			return
		case <-tick.C:
			incr++
		}
	}
}
