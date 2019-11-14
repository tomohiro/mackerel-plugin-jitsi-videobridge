package mpjitsivideobridge

import (
	"testing"
)

func TestMetricKeyPrefix(t *testing.T) {
	var p JitsiVideobridgePlugin

	expected := "jitsi-videobridge"
	actual := p.MetricKeyPrefix()

	if actual != expected {
		t.Errorf("MetricKeyPrefix: %v should be %v", actual, expected)
	}
}
