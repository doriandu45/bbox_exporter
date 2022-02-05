// Copyright (C) 2021 Nicolas Lamirault <nicolas.lamirault@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exporter

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/nlamirault/bbox_exporter/bbox"
)

var (
	ftthState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_ftth_state"),
		"LinkState of the GEth FTTH port",
		nil, nil,
	)
	txBytesWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_bytes"),
		"TX bytes",
		nil, nil,
	)
	txPacketsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_packets"),
		"TX packets",
		nil, nil,
	)
	txPacketsErrorsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_packets_errors"),
		"TX packets in error",
		nil, nil,
	)
	txPacketsDiscardsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_packets_discards"),
		"TX packets discards",
		nil, nil,
	)
	txLineOccupationWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_line_occupation"),
		"TX line occupation",
		nil, nil,
	)
	txBandwidthWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_bandwidth"),
		"TX bandwith available",
		nil, nil,
	)
	txBandwidthMaxWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_transmitted_bandwidth_max"),
		"TX maximum bandwith available",
		nil, nil,
	)

	rxBytesWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_bytes"),
		"RX bytes",
		nil, nil,
	)
	rxPacketsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_packets"),
		"RX packets",
		nil, nil,
	)
	rxPacketsErrorsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_packets_errors"),
		"RX packets in error",
		nil, nil,
	)
	rxPacketsDiscardsWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_packets_discards"),
		"RX packets discards",
		nil, nil,
	)
	rxLineOccupationWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_line_occupation"),
		"RX line occupation",
		nil, nil,
	)
	rxBandwidthWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_bandwidth"),
		"RX bandwith available",
		nil, nil,
	)
	rxBandwidthMaxWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_received_bandwidth_max"),
		"RX bandwith available",
		nil, nil,
	)
	diagnosticsMinWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_min"),
		"Minimum response Time",
		[]string{"mode"}, nil,
	)
	diagnosticsMaxWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_max"),
		"Maximum response Time",
		[]string{"mode"}, nil,
	)
	diagnosticsAvgWan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_avg"),
		"Average response Time",
		[]string{"mode"}, nil,
	)
	diagnosticsNumberOfSuccess = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_success"),
		"Number of sucess",
		[]string{"mode"}, nil,
	)
	diagnosticsNumberOfError = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_error"),
		"Number of error",
		[]string{"mode"}, nil,
	)
	diagnosticsNumberOfTries = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "wan_diagnostics_tries"),
		"Number of tries",
		[]string{"mode"}, nil,
	)
	xDslLocalFEC = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_down_fec"),
		"Number of FEC errors for downstream",
		nil, nil,
	)
	xDslRemoteFEC = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_up_fec"),
		"Number of FEC errors for upstream",
		nil, nil,
	)
	xDslLocalCRC = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_down_crc"),
		"Number of CRC errors for downstream",
		nil, nil,
	)
	xDslRemoteCRC = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_up_crc"),
		"Number of CRC errors for upstream",
		nil, nil,
	)
	xDslLocalHEC = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_down_hec"),
		"Number of HEC errors for downstream",
		nil, nil,
	)
	xDslRemoteHEC = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_up_hec"),
		"Number of HEC errors for upstream",
		nil, nil,
	)
	xDslStatus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_status"),
		"Status of the xDsl connexion",
		nil, nil,
	)
	xDslModulation = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_modulation"),
		"Modulation of the xDsl link",
		[]string{"modulation"}, nil,
	)
	xDslShowtime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_showtime"),
		"Active time of the xDsl link in seconds",
		nil, nil,
	)
	xDslATUR = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_atu-r"),
		"Provider of the ATU-R chipset",
		[]string{"provider"}, nil,
	)
	xDslATUC = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_atu-c"),
		"Provider of the ATU-C chupset",
		[]string{"provider"}, nil,
	)
	xDslSyncCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_sync_count"),
		"Number of xDsl synchronisations  since last reboot",
		nil, nil,
	)
	xDslUpBitrate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_up_bitrate"),
		"Speed of the xDsl upstream in kB/s",
		nil, nil,
	)
	xDslUpNoise = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_up_noise"),
		"Uplink noise in cB (0.1 dB)",
		nil, nil,
	)
	xDslUpAttenuation = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_up_attenuation"),
		"Uplink attenuation in cB (0.1 dB)",
		nil, nil,
	)
	xDslUpPower = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_up_power"),
		"Uplink power in whatever unit it is. No, the BBox API doc don't tell anything about that",
		nil, nil,
	)
	xDslUpBoost = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_up_boost"),
		"Indicates which speed boosting technology is used for the upstream",
		[]string{"used"}, nil,
	)
	xDslUpInterleave = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_up_interleave"),
		"Uplink interleave delay (unit unknown)",
		nil, nil,
	)
	xDslDownBitrate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_down_bitrate"),
		"Speed of the xDsl downstream in kB/s",
		nil, nil,
	)
	xDslDownNoise = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_down_noise"),
		"Downlink noise in cB (0.1 dB)",
		nil, nil,
	)
	xDslDownAttenuation = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_down_attenuation"),
		"Downlink attenuation in cB (0.1 dB)",
		nil, nil,
	)
	xDslDownPower = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_down_power"),
		"Downlink power in whatever unit it is. No, the BBox API doc don't tell anything about that",
		nil, nil,
	)
	xDslDownBoost = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_down_boost"),
		"Indicates which speed boosting technology is used for the downstream",
		[]string{"used"}, nil,
	)
	xDslDownInterleave = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "xdsl_down_interleave"),
		"Downlink interleave delay (unit unknown)",
		nil, nil,
	)
)

func describeWanMetrics(ch chan<- *prometheus.Desc) {
	ch <- ftthState
	ch <- txBytesWan
	ch <- txPacketsWan
	ch <- txPacketsErrorsWan
	ch <- txPacketsDiscardsWan
	ch <- txLineOccupationWan
	ch <- txBandwidthWan
	ch <- txBandwidthMaxWan
	ch <- rxBytesWan
	ch <- rxPacketsWan
	ch <- rxPacketsErrorsWan
	ch <- rxPacketsDiscardsWan
	ch <- rxLineOccupationWan
	ch <- rxBandwidthWan
	ch <- rxBandwidthMaxWan
	ch <- diagnosticsMinWan
	ch <- diagnosticsMaxWan
	ch <- diagnosticsAvgWan
	ch <- diagnosticsNumberOfSuccess
	ch <- diagnosticsNumberOfError
	ch <- diagnosticsNumberOfTries
	ch <- xDslLocalFEC
	ch <- xDslRemoteFEC
	ch <- xDslLocalCRC
	ch <- xDslRemoteCRC
	ch <- xDslLocalHEC
	ch <- xDslRemoteHEC
	ch <- xDslStatus
	ch <- xDslModulation
	ch <- xDslShowtime
	ch <- xDslATUR
	ch <- xDslATUC
	ch <- xDslSyncCount
	ch <- xDslUpBitrate
	ch <- xDslUpNoise
	ch <- xDslUpAttenuation
	ch <- xDslUpNoise
	ch <- xDslUpPower
	ch <- xDslUpNoise
	ch <- xDslUpBoost
	ch <- xDslUpInterleave
	ch <- xDslDownBitrate
	ch <- xDslDownNoise
	ch <- xDslDownAttenuation
	ch <- xDslDownNoise
	ch <- xDslDownPower
	ch <- xDslDownNoise
	ch <- xDslDownBoost
	ch <- xDslDownInterleave

}

func storeWanMetrics(ch chan<- prometheus.Metric, metrics bbox.WanMetrics) {
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Bytes), txBytesWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Packets), txPacketsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Packetserrors), txPacketsErrorsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Packetsdiscards), txPacketsDiscardsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Occupation), txLineOccupationWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.Bandwidth), txBandwidthWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Tx.MaxBandwidth), txBandwidthMaxWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Bytes), rxBytesWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Packets), rxPacketsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Packetserrors), rxPacketsErrorsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Packetsdiscards), rxPacketsDiscardsWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Occupation), rxLineOccupationWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.Bandwidth), rxBandwidthWan)
	storeMetric(ch, float64(metrics.IPStatistics[0].WAN.IP.Stats.Rx.MaxBandwidth), rxBandwidthMaxWan)

	for _, val := range metrics.DiagnosticsStatistics[0].Diags.DNS {
		if val.Tries > 0 {
			storeMetric(ch, float64(val.Min), diagnosticsMinWan, "dns")
			storeMetric(ch, float64(val.Max), diagnosticsMaxWan, "dns")
			storeMetric(ch, float64(val.Average), diagnosticsAvgWan, "dns")
			break
		}
	}
	for _, val := range metrics.DiagnosticsStatistics[0].Diags.Ping {
		if val.Tries > 0 {
			storeMetric(ch, float64(val.Min), diagnosticsMinWan, "ping")
			storeMetric(ch, float64(val.Max), diagnosticsMaxWan, "ping")
			storeMetric(ch, float64(val.Average), diagnosticsAvgWan, "ping")
			break
		}
	}
	for _, val := range metrics.DiagnosticsStatistics[0].Diags.HTTP {
		if val.Tries > 0 {
			storeMetric(ch, float64(val.Min), diagnosticsMinWan, "http")
			storeMetric(ch, float64(val.Max), diagnosticsMaxWan, "http")
			storeMetric(ch, float64(val.Average), diagnosticsAvgWan, "http")
			break
		}
	}
	storeMetric(ch, float64(metrics.XDslStatistics[0].Wan.XDsl.Stats.LocalFEC), xDslLocalFEC)
	storeMetric(ch, float64(metrics.XDslStatistics[0].Wan.XDsl.Stats.RemoteFEC), xDslRemoteFEC)
	storeMetric(ch, float64(metrics.XDslStatistics[0].Wan.XDsl.Stats.LocalCRC), xDslLocalCRC)
	storeMetric(ch, float64(metrics.XDslStatistics[0].Wan.XDsl.Stats.RemoteCRC), xDslRemoteCRC)
	storeMetric(ch, float64(metrics.XDslStatistics[0].Wan.XDsl.Stats.LocalHEC), xDslLocalHEC)
	storeMetric(ch, float64(metrics.XDslStatistics[0].Wan.XDsl.Stats.RemoteHEC), xDslRemoteHEC)

	if metrics.XDslInformations[0].Wan.XDsl.State == "Connected" {
		storeMetric(ch, 1.0, xDslStatus)
	} else {
		storeMetric(ch, 0.0, xDslStatus)
	}

	storeMetric(ch, 1.0, xDslModulation, metrics.XDslInformations[0].Wan.XDsl.Modulation)
	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Showtime), xDslShowtime)
	storeMetric(ch, 1.0, xDslATUR, metrics.XDslInformations[0].Wan.XDsl.ATURProvider)
	storeMetric(ch, 1.0, xDslATUC, metrics.XDslInformations[0].Wan.XDsl.ATUCProcider)
	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.SyncCount), xDslSyncCount)

	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Up.Biterates), xDslUpBitrate)
	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Up.Noise), xDslUpNoise)
	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Up.Attenuation), xDslUpAttenuation)
	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Up.Power), xDslUpPower)

	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Up.PhyR), xDslUpBoost, "phyr")
	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Up.GINP), xDslUpBoost, "ginp")
	storeMetric(ch, 0.0, xDslUpBoost, "nitro") // See bbox/wan.go

	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Up.InterleaveDelay), xDslUpInterleave)

	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Down.Biterates), xDslDownBitrate)
	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Down.Noise), xDslDownNoise)
	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Down.Attenuation), xDslDownAttenuation)
	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Down.Power), xDslDownPower)

	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Down.PhyR), xDslDownBoost, "phyr")
	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Down.GINP), xDslDownBoost, "ginp")
	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Down.Nitro), xDslDownBoost, "nitro")

	storeMetric(ch, float64(metrics.XDslInformations[0].Wan.XDsl.Down.InterleaveDelay), xDslDownInterleave)

}

func storeWanFtthMetric(ch chan<- prometheus.Metric, metric string) {
	fftStateValue := float64(0)
	if strings.ToUpper(metric) == "UP" {
		fftStateValue = float64(1)
	}
	ch <- prometheus.MustNewConstMetric(
		ftthState, prometheus.GaugeValue, fftStateValue,
	)
}
