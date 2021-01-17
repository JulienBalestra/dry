package promregister

import (
	"github.com/prometheus/client_golang/prometheus"
)

func Register(collectors ...prometheus.Collector) error {
	for _, c := range collectors {
		err := prometheus.Register(c)
		if err != nil {
			return err
		}
	}
	return nil
}
