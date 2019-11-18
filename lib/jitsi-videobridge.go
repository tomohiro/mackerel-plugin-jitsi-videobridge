package mpjitsivideobridge

import (
	"flag"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

// JitsiVideobridgePlugin as Mackerel agent plugin for Jitsi Videobridge
type JitsiVideobridgePlugin struct {
	Prefix string
	Host   string
	Port   string
}

// MetricKeyPrefix returns prefix of Jitsi Videobridge metrics
func (p JitsiVideobridgePlugin) MetricKeyPrefix() string {
	return p.Prefix
}

// GraphDefinition returns graph definition
func (p JitsiVideobridgePlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(p.MetricKeyPrefix())
	return map[string]mp.Graphs{
		"Audio Channels": {
			Label: labelPrefix,
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "audiochannels", Label: "Audio Channels"},
			},
		},
	}
}

// FetchMetrics fetches metrics from Jitsi Videobridge Colibri REST interface
func (p JitsiVideobridgePlugin) FetchMetrics() (map[string]float64, error) {
	metrics := make(map[string]float64)
	return metrics, nil
}

// Do the plugin
func Do() {
	optPrefix := flag.String("metric-key-prefix", "jitsi-videobridge", "Metric key prefix")
	optHost := flag.String("host", "127.0.0.1", "Hostname or IP address of Jitsi Videobridge Colibri REST interface")
	optPort := flag.String("port", "80", "Port of Jitsi Videobridge Colibri REST interface")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	flag.Parse()

	p := JitsiVideobridgePlugin{
		Prefix: *optPrefix,
		Host:   *optHost,
		Port:   *optPort,
	}

	helper := mp.NewMackerelPlugin(p)
	helper.Tempfile = *optTempfile
	helper.Run()
}
