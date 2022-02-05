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

package bbox

import (
	"github.com/go-kit/kit/log/level"
)

type WanMetrics struct {
	IPInformations        []WanIPInformations  `json:"ip_informations"`
	IPStatistics          []WanIPStatistics    `json:"ip_statistics"`
	FtthStatistics        *FtthStatistics      `json:"ftth_statistics"`
	DiagnosticsStatistics []WanDiagsStatistics `json:"diagnostics"`
	XDslStatistics        []WanXDslStat        `json:"xdsl_statistics"`
	XDslInformations      []WanXDslInfo        `json:"xdsl_informations"`
}

type WanIPStatistics struct {
	WAN struct {
		IP struct {
			Stats struct {
				Rx struct {
					Packets         flexInt `json:"packets"`
					Bytes           flexInt `json:"bytes"` // See: https://github.com/nlamirault/bbox_exporter/issues/1
					Packetserrors   flexInt `json:"packetserrors"`
					Packetsdiscards flexInt `json:"packetsdiscards"`
					Occupation      float64 `json:"occupation"`
					Bandwidth       flexInt `json:"bandwidth"`
					MaxBandwidth    flexInt `json:"maxBandwidth"`
				} `json:"rx"`
				Tx struct {
					Packets         flexInt `json:"packets"`
					Bytes           flexInt `json:"bytes"` // See: https://github.com/nlamirault/bbox_exporter/issues/1
					Packetserrors   flexInt `json:"packetserrors"`
					Packetsdiscards flexInt `json:"packetsdiscards"`
					Occupation      float64 `json:"occupation"`
					Bandwidth       flexInt `json:"bandwidth"`
					MaxBandwidth    flexInt `json:"maxBandwidth"`
				} `json:"tx"`
			} `json:"stats"`
		} `json:"ip"`
	} `json:"wan"`
}

type FtthStatistics []struct {
	Ftth Ftth `json:"ftth"`
}

type Ftth struct {
	Wan struct {
		Ftth struct {
			Mode  string `json:"mode"`
			State string `json:"state"`
		} `json:"ftth"`
	} `json:"wan"`
}

type WanXDslStat struct {
	Wan struct {
		XDsl struct {
			Stats struct {
				LocalFEC  int `json:"local_fec"`
				RemoteFEC int `json:"remote_fec"`
				LocalCRC  int `json:"local_crc"`
				RemoteCRC int `json:"remote_crc"`
				LocalHEC  int `json:"local_hec"`
				RemoteHEC int `json:"remote_hec"`
			} `json:"stats"`
		} `json:"xdsl"`
	} `json:"wan"`
}

type WanXDslInfo struct {
	Wan struct {
		XDsl struct {
			State        string `json:"state"`
			Modulation   string `json:"modulation"`
			Showtime     int    `json:"showtime"`
			ATURProvider string `json:"atur_provider"`
			ATUCProcider string `json:"atuc_provider"`
			SyncCount    int    `json:"sync_count"`
			Up           struct {
				Biterates   int `json:"bitrates"`
				Noise       int `json:"noise"`
				Attenuation int `json:"attenuation"`
				Power       int `json:"power"`
				PhyR        int `json:"phyr"`
				GINP        int `json:"ginp"`
				// Nitro       int `json:"nitro"` // Not enabled because in my case, it's an empty string (but it's the int 0 in the down struct, go figure why)
				InterleaveDelay int `json:"interleave_delay"`
			} `json:"up"`
			Down struct {
				Biterates       int `json:"bitrates"`
				Noise           int `json:"noise"`
				Attenuation     int `json:"attenuation"`
				Power           int `json:"power"`
				PhyR            int `json:"phyr"`
				GINP            int `json:"ginp"`
				Nitro           int `json:"nitro"`
				InterleaveDelay int `json:"interleave_delay"`
			} `json:"down"`
		} `json:"xdsl"`
	} `json:"wan"`
}

type WanIPInformations struct {
	Wan struct {
		Internet struct {
			State int `json:"state"`
		} `json:"internet"`
		Interface struct {
			ID      int `json:"id"`
			Default int `json:"default"`
			State   int `json:"state"`
		} `json:"interface"`
		IP struct {
			Address    string        `json:"address"`
			State      string        `json:"state"`
			Gateway    string        `json:"gateway"`
			Dnsservers string        `json:"dnsservers"`
			Subnet     string        `json:"subnet"`
			IP6State   string        `json:"ip6state"`
			IP6Address []interface{} `json:"ip6address"`
			IP6Prefix  []interface{} `json:"ip6prefix"`
			Mac        string        `json:"mac"`
			Mtu        int           `json:"mtu"`
		} `json:"ip"`
		Link struct {
			State string `json:"state"`
			Type  string `json:"type"`
		} `json:"link"`
	} `json:"wan"`
}

type WanDiagsStatistics struct {
	Diags struct {
		DNS []struct {
			Min      float64 `json:"min"`
			Max      float64 `json:"max"`
			Average  float64 `json:"average"`
			Success  int     `json:"success"`
			Error    int     `json:"error"`
			Tries    int     `json:"tries"`
			Status   string  `json:"status"`
			Protocol string  `json:"protocol"`
		} `json:"dns"`
		Ping []struct {
			Min      float64 `json:"min"`
			Max      float64 `json:"max"`
			Average  float64 `json:"average"`
			Success  int     `json:"success"`
			Error    int     `json:"error"`
			Tries    int     `json:"tries"`
			Status   string  `json:"status"`
			Protocol string  `json:"protocol"`
		} `json:"ping"`
		HTTP []struct {
			Min      float64 `json:"min"`
			Max      float64 `json:"max"`
			Average  float64 `json:"average"`
			Success  int     `json:"success"`
			Error    int     `json:"error"`
			Tries    int     `json:"tries"`
			Status   string  `json:"status"`
			Protocol string  `json:"protocol"`
		} `json:"http"`
	} `json:"diags"`
}

func (client *Client) getWanMetrics() (*WanMetrics, error) {
	var metrics WanMetrics

	wanIPInformations, err := client.getWanInformations()
	if err != nil {
		return nil, err
	}
	metrics.IPInformations = wanIPInformations

	wanIPStats, err := client.getWanStatistics()
	if err != nil {
		return nil, err
	}
	metrics.IPStatistics = wanIPStats

	ftthStats, err := client.getWanFtthStatistics()
	if err != nil {
		return nil, err
	}
	metrics.FtthStatistics = ftthStats

	diagsStats, err := client.getWANDiagnostics()
	if err != nil {
		return nil, err
	}
	metrics.DiagnosticsStatistics = diagsStats

	xDslStats, err := client.getXDslStatistics()
	if err != nil {
		return nil, err
	}
	metrics.XDslStatistics = xDslStats

	xDslInfos, err := client.getXDslInformations()
	if err != nil {
		return nil, err
	}
	metrics.XDslInformations = xDslInfos

	return &metrics, nil
}

// getWanInformations returns WAN IP Information
// See: https://api.bbox.fr/doc/apirouter/#api-WAN-GetWANIP
func (client *Client) getWanInformations() ([]WanIPInformations, error) {
	level.Info(client.logger).Log("msg", "Retrieve WAN IP informations from Bbox")
	var informations []WanIPInformations
	if err := client.apiRequest("/wan/ip", &informations); err != nil {
		return nil, err
	}
	return informations, nil
}

// getWanStatistics returns WAN IP statistics
// See: https://api.bbox.fr/doc/apirouter/#api-WAN-GetWANIPStats
func (client *Client) getWanStatistics() ([]WanIPStatistics, error) {
	level.Info(client.logger).Log("msg", "Retrieve WAN metrics from Bbox")
	var metrics []WanIPStatistics
	if err := client.apiRequest("/wan/ip/stats", &metrics); err != nil {
		return nil, err
	}
	return metrics, nil
}

// getWanFtthStatistics returns information about FTTH
// See: https://api.bbox.fr/doc/apirouter/#api-WAN-GetFTTHStats
func (client *Client) getWanFtthStatistics() (*FtthStatistics, error) {
	level.Info(client.logger).Log("msg", "Retrieve WAN metrics from Bbox")
	var metrics FtthStatistics
	if err := client.apiRequest("/wan/ftth/stats", &metrics); err != nil {
		return nil, err
	}
	return &metrics, nil
}

// getWANDiagnostics return results of the tests to retrieve the real state of the Internet connectivity
// https://api.bbox.fr/doc/apirouter/index.html#api-WAN-GetWANDiags
func (client *Client) getWANDiagnostics() ([]WanDiagsStatistics, error) {
	level.Info(client.logger).Log("msg", "Retrieve WAN diagnostics from Bbox")
	var metrics []WanDiagsStatistics
	if err := client.apiRequest("/wan/diags", &metrics); err != nil {
		return nil, err
	}
	return metrics, nil
}

// getXDslInformations returns information about xDsl
// https://api.bbox.fr/doc/apirouter/index.html#api-WAN-GetWANXDSL
func (client *Client) getXDslInformations() ([]WanXDslInfo, error) {
	level.Info(client.logger).Log("msg", "Retrieve xDsl informations from Bbox")
	var metrics []WanXDslInfo
	if err := client.apiRequest("/wan/xdsl", &metrics); err != nil {
		return nil, err
	}
	return metrics, nil
}

// getXDslStatistics returns statistics about xDsl
// https://api.bbox.fr/doc/apirouter/index.html#api-WAN-GetWANXDSLStats
func (client *Client) getXDslStatistics() ([]WanXDslStat, error) {
	level.Info(client.logger).Log("msg", "Retrieve xDsl statistics from Bbox")
	var metrics []WanXDslStat
	if err := client.apiRequest("/wan/xdsl/stats", &metrics); err != nil {
		return nil, err
	}
	return metrics, nil
}
