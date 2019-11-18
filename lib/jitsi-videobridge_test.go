package mpjitsivideobridge

import (
	"testing"
)

var (
	plugin JitsiVideobridgePlugin
)

func setup() {
	plugin = JitsiVideobridgePlugin{
		Prefix: "jitsi-videobridge",
		Host:   "127.0.0.1",
		Port:   "80",
	}
}

func TestMetricKeyPrefix(t *testing.T) {
	setup()

	want := "jitsi-videobridge"
	got := plugin.MetricKeyPrefix()

	if want != got {
		t.Errorf("MetricKeyPrefix: %v should be %v", want, got)
	}
}

func TestGraphDefinition(t *testing.T) {
	t.Skip()
}

func TestFetchMetrics(t *testing.T) {
	t.Skip()
}
