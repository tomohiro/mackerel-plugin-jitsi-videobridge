package mpjitsivideobridge

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

// JitsiVideobridgePlugin as Mackerel agent plugin for Jitsi Videobridge
type JitsiVideobridgePlugin struct {
	KeyPrefix   string
	LabelPrefix string
	Host        string
	Port        string
}

// Stats represents a statistics.
//
// https://github.com/jitsi/jitsi-videobridge/blob/master/doc/statistics.md
// https://github.com/jitsi/jitsi-videobridge/blob/master/src/main/java/org/jitsi/videobridge/stats/VideobridgeStatistics.java
type Stats struct {
	Audiochannels                         float64   `json:"audiochannels"`
	BitRateDownload                       float64   `json:"bit_rate_download"`
	BitRateUpload                         float64   `json:"bit_rate_upload"`
	ConferenceSizes                       []float64 `json:"conference_sizes"`
	Conferences                           float64   `json:"conferences"`
	CPUUsage                              float64   `json:"cpu_usage"`
	CurrentTimestamp                      string    `json:"current_timestamp"`
	GracefulShutdown                      bool      `json:"graceful_shutdown"`
	JitterAggregate                       float64   `json:"jitter_aggregate"`
	LargestConference                     float64   `json:"largest_conference"`
	LossRateDownload                      float64   `json:"loss_rate_download"`
	LossRateUpload                        float64   `json:"loss_rate_upload"`
	PacketRateDownload                    float64   `json:"packet_rate_download"`
	PacketRateUpload                      float64   `json:"packet_rate_upload"`
	Participants                          float64   `json:"participants"`
	Region                                string    `json:"region"`
	RelayID                               string    `json:"relay_id"`
	RtpLoss                               float64   `json:"rtp_loss"`
	RttAggregate                          float64   `json:"rtt_aggregate"`
	Threads                               float64   `json:"threads"`
	TotalBytesReceived                    float64   `json:"total_bytes_received"`
	TotalBytesReceivedOcto                float64   `json:"total_bytes_received_octo"`
	TotalBytesSent                        float64   `json:"total_bytes_sent"`
	TotalBytesSentOcto                    float64   `json:"total_bytes_sent_octo"`
	TotalColibriWebSocketMessagesReceived float64   `json:"total_colibri_web_socket_messages_received"`
	TotalColibriWebSocketMessagesSent     float64   `json:"total_colibri_web_socket_messages_sent"`
	TotalConferenceSeconds                float64   `json:"total_conference_seconds"`
	TotalConferencesCompleted             float64   `json:"total_conferences_completed"`
	TotalConferencesCreated               float64   `json:"total_conferences_created"`
	TotalDataChannelMessagesReceived      float64   `json:"total_data_channel_messages_received"`
	TotalDataChannelMessagesSent          float64   `json:"total_data_channel_messages_sent"`
	TotalFailedConferences                float64   `json:"total_failed_conferences"`
	TotalIceFailed                        float64   `json:"total_ice_failed"`
	TotalIceSucceeded                     float64   `json:"total_ice_succeeded"`
	TotalIceSucceededTCP                  float64   `json:"total_ice_succeeded_tcp"`
	TotalLossControlledParticipantSeconds float64   `json:"total_loss_controlled_participant_seconds"`
	TotalLossDegradedParticipantSeconds   float64   `json:"total_loss_degraded_participant_seconds"`
	TotalLossLimitedParticipantSeconds    float64   `json:"total_loss_limited_participant_seconds"`
	TotalMemory                           float64   `json:"total_memory"`
	TotalPacketsDroppedOcto               float64   `json:"total_packets_dropped_octo"`
	TotalPacketsReceived                  float64   `json:"total_packets_received"`
	TotalPacketsReceivedOcto              float64   `json:"total_packets_received_octo"`
	TotalPacketsSent                      float64   `json:"total_packets_sent"`
	TotalPacketsSentOcto                  float64   `json:"total_packets_sent_octo"`
	TotalPartiallyFailedConferences       float64   `json:"total_partially_failed_conferences"`
	TotalParticipants                     float64   `json:"total_participants"`
	UsedMemory                            float64   `json:"used_memory"`
	Videochannels                         float64   `json:"videochannels"`
	Videostreams                          float64   `json:"videostreams"`
}

// MetricKeyPrefix returns prefix of Jitsi Videobridge metrics
func (p JitsiVideobridgePlugin) MetricKeyPrefix() string {
	return p.KeyPrefix
}

// GraphDefinition returns graph definition
func (p JitsiVideobridgePlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := p.LabelPrefix
	return map[string]mp.Graphs{
		"cpu": {
			Label: fmt.Sprintf("%s: cpu", labelPrefix),
			Unit:  mp.UnitPercentage,
			Metrics: []mp.Metrics{
				{Name: "cpu_usage", Label: "Usage", Scale: 100},
			},
		},
		"memory": {
			Label: fmt.Sprintf("%s: memory", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "total_memory", Label: "Total"},
				{Name: "used_memory", Label: "Used"},
			},
		},
		"thread": {
			Label: fmt.Sprintf("%s: JVM threads", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "threads", Label: "Threads"},
			},
		},
		"conference": {
			Label: fmt.Sprintf("%s: Conference", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "conferences", Label: "Ongoing"},
				{Name: "total_conferences_completed", Label: "Completed", Diff: true},
				{Name: "total_conferences_created", Label: "Created", Diff: true},
				{Name: "total_failed_conferences", Label: "Failed", Diff: true},
				{Name: "total_partially_failed_conferences", Label: "Partially Failed", Diff: true},
			},
		},
		"conference_time": {
			Label: fmt.Sprintf("%s: Conference Time (sec)", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "total_conference_seconds", Label: "Completed"},
				{Name: "total_loss_controlled_participant_seconds", Label: "Loss Controlled"},
				{Name: "total_loss_degraded_participant_seconds", Label: "Loss Degraded"},
				{Name: "total_loss_limited_participant_seconds", Label: "Loss Limited"},
			},
		},
	}
}

// FetchMetrics fetches metrics from Jitsi Videobridge Colibri REST float64erface
func (p JitsiVideobridgePlugin) FetchMetrics() (map[string]float64, error) {
	url := fmt.Sprintf("http://%v:%v/colibri/stats", p.Host, p.Port)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get a request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad response code: %v", res.StatusCode)
	}

	stats := Stats{}
	if err = json.NewDecoder(res.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to decode a responded JSON. %w", err)
	}

	return transformStatsToMetrics(&stats), nil
}

func transformStatsToMetrics(s *Stats) map[string]float64 {
	metrics := make(map[string]float64)

	// CPU
	metrics["cpu_usage"] = s.CPUUsage

	// Memory
	metrics["total_memory"] = s.TotalMemory
	metrics["used_memory"] = s.UsedMemory

	// Java
	metrics["threads"] = s.Threads

	// Conferences
	metrics["conferences"] = s.Conferences
	metrics["total_conferences_completed"] = s.TotalConferencesCompleted
	metrics["total_conferences_created"] = s.TotalConferencesCreated
	metrics["total_failed_conferences"] = s.TotalFailedConferences
	metrics["total_partially_failed_conferences"] = s.TotalPartiallyFailedConferences

	// Conference Lengths
	metrics["total_conference_seconds"] = s.TotalConferenceSeconds
	metrics["total_loss_controlled_participant_seconds"] = s.TotalLossControlledParticipantSeconds
	metrics["total_loss_degraded_participant_seconds"] = s.TotalLossDegradedParticipantSeconds
	metrics["total_loss_limited_participant_seconds"] = s.TotalLossLimitedParticipantSeconds

	// Participants
	metrics["participants"] = s.Participants
	metrics["total_participants"] = s.TotalParticipants
	metrics["largest_conference"] = s.LargestConference

	// Channels / Streams
	metrics["audiochannels"] = s.Audiochannels
	metrics["videochannels"] = s.Videochannels
	metrics["videostreams"] = s.Videostreams

	// ICE connection total statuses
	metrics["total_ice_succeeded"] = s.TotalIceSucceeded
	metrics["total_ice_succeeded_tcp"] = s.TotalIceSucceededTCP
	metrics["total_ice_failed"] = s.TotalIceFailed

	// Jitter (Experimentail)
	metrics["jitter_aggregate"] = s.JitterAggregate

	// RTT
	metrics["rtt_aggregate"] = s.RttAggregate

	// Videobridge bit rate download / upload
	metrics["bit_rate_download"] = s.BitRateDownload
	metrics["bit_rate_upload"] = s.BitRateUpload

	// Videobridge packet rate download / upload
	metrics["packet_rate_download"] = s.PacketRateDownload
	metrics["packet_rate_upload"] = s.PacketRateUpload

	// RTP packet loss rate download / upload
	metrics["loss_rate_download"] = s.LossRateDownload
	metrics["loss_rate_upload"] = s.LossRateUpload

	// Bytes received / sent
	metrics["total_bytes_received"] = s.TotalBytesReceived
	metrics["total_bytes_received_octo"] = s.TotalBytesReceivedOcto
	metrics["total_bytes_sent"] = s.TotalBytesSent
	metrics["total_bytes_sent_octo"] = s.TotalBytesSentOcto

	// Packet total received / sent / dropped
	metrics["total_packets_received"] = s.TotalPacketsReceived
	metrics["total_packets_received_octo"] = s.TotalPacketsReceivedOcto
	metrics["total_packets_sent"] = s.TotalPacketsSent
	metrics["total_packets_sent_octo"] = s.TotalPacketsSentOcto
	metrics["total_packets_dropped_octo"] = s.TotalPacketsDroppedOcto

	// Datachannel messages total received / sent
	metrics["total_data_channel_messages_received"] = s.TotalDataChannelMessagesReceived
	metrics["total_data_channel_messages_sent"] = s.TotalDataChannelMessagesSent

	// Colibri WebSocket messages total received / sent
	metrics["total_colibri_web_socket_messages_received"] = s.TotalColibriWebSocketMessagesReceived
	metrics["total_colibri_web_socket_messages_sent"] = s.TotalColibriWebSocketMessagesSent

	return metrics
}

// Do the plugin
func Do() {
	optKeyPrefix := flag.String("metric-key-prefix", "jitsi-videobridge", "Metric key prefix")
	optLabelPrefix := flag.String("metric-label-prefix", "JVB", "Metric label prefix")
	optHost := flag.String("host", "127.0.0.1", "Hostname or IP address of Jitsi Videobridge Colibri REST interface")
	optPort := flag.String("port", "80", "Port of Jitsi Videobridge Colibri REST interface")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	flag.Parse()

	p := JitsiVideobridgePlugin{
		KeyPrefix:   *optKeyPrefix,
		LabelPrefix: *optLabelPrefix,
		Host:        *optHost,
		Port:        *optPort,
	}

	helper := mp.NewMackerelPlugin(p)
	helper.Tempfile = *optTempfile
	helper.Run()
}
