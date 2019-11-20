package mpjitsivideobridge

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	plugin JitsiVideobridgePlugin
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	u, _ := url.Parse(server.URL)
	plugin = JitsiVideobridgePlugin{
		KeyPrefix:   "jitsi-videobridge",
		LabelPrefix: "JVB",
		Host:        u.Hostname(),
		Port:        u.Port(),
	}
}

func teardown() {
	server.Close()
}

func TestMetricKeyPrefix(t *testing.T) {
	setup()
	defer teardown()

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
	setup()
	defer teardown()

	mux.HandleFunc("/colibri/stats", func(w http.ResponseWriter, r *http.Request) {
		// Set response headers
		w.Header().Set("Content-Type", "application/json")

		// Set 200 OK as HTTP status code.
		w.WriteHeader(http.StatusOK)

		// Set response body
		// This example is response body that from https://github.com/jitsi/jitsi-videobridge/blob/master/doc/statistics.md
		fmt.Fprintln(w, `
		{
			"audiochannels": 0,
			"bit_rate_download": 0,
			"bit_rate_upload": 0,
			"conference_sizes": [ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 ],
			"conferences": 0,
			"cpu_usage": 0.2358490566037736,
			"current_timestamp": "2019-03-14 11:02:15.184",
			"graceful_shutdown": false,
			"jitter_aggregate": 0,
			"largest_conference": 0,
			"loss_rate_download": 0,
			"loss_rate_upload": 0,
			"packet_rate_download": 0,
			"packet_rate_upload": 0,
			"participants": 0,
			"region": "eu-west-1",
			"relay_id": "10.0.0.5:4096",
			"rtp_loss": 0,
			"rtt_aggregate": 0,
			"threads": 59,
			"total_bytes_received": 257628359,
			"total_bytes_received_octo": 0,
			"total_bytes_sent": 257754048,
			"total_bytes_sent_octo": 0,
			"total_colibri_web_socket_messages_received": 0,
			"total_colibri_web_socket_messages_sent": 0,
			"total_conference_seconds": 470,
			"total_conferences_completed": 1,
			"total_conferences_created": 1,
			"total_data_channel_messages_received": 602,
			"total_data_channel_messages_sent": 600,
			"total_failed_conferences": 0,
			"total_ice_failed": 0,
			"total_ice_succeeded": 2,
			"total_ice_succeeded_tcp": 0,
			"total_loss_controlled_participant_seconds": 847,
			"total_loss_degraded_participant_seconds": 1,
			"total_loss_limited_participant_seconds": 0,
			"total_memory": 8257,
			"total_packets_dropped_octo": 0,
			"total_packets_received": 266644,
			"total_packets_received_octo": 0,
			"total_packets_sent": 266556,
			"total_packets_sent_octo": 0,
			"total_partially_failed_conferences": 0,
			"total_participants": 2,
			"used_memory": 4404,
			"videochannels": 0,
			"videostreams": 0
		  }
		`)
	})

	stats, err := plugin.FetchMetrics()

	if err != nil {
		t.Error(err)
	}

	if stats == nil {
		t.Errorf("empty response: %#v", stats)
	}

	got := stats
	want := map[string]float64{
		"audiochannels":     0,
		"bit_rate_download": 0,
		"bit_rate_upload":   0,
		// "conference_sizes": [ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 ],
		"conferences": 0,
		"cpu_usage":   0.2358490566037736,
		// "current_timestamp": "2019-03-14 11:02:15.184",
		// "graceful_shutdown": false,
		"jitter_aggregate":     0,
		"largest_conference":   0,
		"loss_rate_download":   0,
		"loss_rate_upload":     0,
		"packet_rate_download": 0,
		"packet_rate_upload":   0,
		"participants":         0,
		// "region": "eu-west-1",
		// "relay_id": "10.0.0.5:4096",
		// "rtp_loss": 0,
		"rtt_aggregate":                              0,
		"threads":                                    59,
		"total_bytes_received":                       257628359,
		"total_bytes_received_octo":                  0,
		"total_bytes_sent":                           257754048,
		"total_bytes_sent_octo":                      0,
		"total_colibri_web_socket_messages_received": 0,
		"total_colibri_web_socket_messages_sent":     0,
		"total_conference_seconds":                   470,
		"total_conferences_completed":                1,
		"total_conferences_created":                  1,
		"total_data_channel_messages_received":       602,
		"total_data_channel_messages_sent":           600,
		"total_failed_conferences":                   0,
		"total_ice_failed":                           0,
		"total_ice_succeeded":                        2,
		"total_ice_succeeded_tcp":                    0,
		"total_loss_controlled_participant_seconds":  847,
		"total_loss_degraded_participant_seconds":    1,
		"total_loss_limited_participant_seconds":     0,
		"total_memory":                               8257,
		"total_packets_dropped_octo":                 0,
		"total_packets_received":                     266644,
		"total_packets_received_octo":                0,
		"total_packets_sent":                         266556,
		"total_packets_sent_octo":                    0,
		"total_partially_failed_conferences":         0,
		"total_participants":                         2,
		"used_memory":                                4404,
		"videochannels":                              0,
		"videostreams":                               0,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Metrics got %#v, want %#v", got, want)
	}
}
