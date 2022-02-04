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

import "github.com/go-kit/kit/log/level"

type DeviceMetrics struct {
	Informations []DeviceInformations `json:"informations"`
	Memory       []DeviceMemory       `json:"device"`
	CPU          []DeviceCPU          `json:"cpu"`
}

type DeviceInformations struct {
	Device struct {
		Now           string  `json:"now"`
		Status        float64 `json:"status"`
		NumberOfBoots float64 `json:"numberofboots"`
		Uptime        int     `json:"uptime"`
		ModelName     string  `json:"modelname"`
		Temperature   struct {
			Current float64 `json:"current"`
			Status  string  `json:"status"`
		} `json:"temperature"`
		Using struct {
			IPv4 int `json:"ipv4"`
			IPv6 int `json:"ipv6"`
			FTTH int `json:"ftth"`
			ADSL int `json:"adsl"`
			VDSL int `json:"vdsl"`
		} `json:"using"`
	} `json:"device"`
}

type DeviceMemory struct {
	Device struct {
		Memory struct {
			Total  float64 `json:"total"`
			Free   float64 `json:"free"`
			Cached float64 `json:"cached"`
		} `json:"mem"`
	} `json:"device"`
}

type DeviceCPU struct {
	Device struct {
		CPU struct {
			Time struct {
				Total  int `json:"total"`
				User   int `json:"user"`
				Nice   int `json:"nice"`
				System int `json:"system"`
				IO     int `json:"io"`
				Idle   int `json:"idle"`
				Irq    int `json:"irq"`
			} `json:"time"`
			Process struct {
				Created int `json:"created"`
				Running int `json:"running"`
				Blocked int `json:"blocked"`
			} `json:"process"`
		} `json:"cpu"`
	} `json:"device"`
}

func (client *Client) getDeviceMetrics() (*DeviceMetrics, error) {
	var deviceStats DeviceMetrics

	informations, err := client.getDeviceInformations()
	if err != nil {
		return nil, err
	}
	deviceStats.Informations = informations

	cpu, err := client.getDeviceCPU()
	if err != nil {
		return nil, err
	}
	deviceStats.CPU = cpu

	memory, err := client.getDeviceMemory()
	if err != nil {
		return nil, err
	}
	deviceStats.Memory = memory

	return &deviceStats, nil
}

// getDeviceInformations returns Bbox information
// See: https://api.bbox.fr/doc/apirouter/#api-Device-GetDevice
func (client *Client) getDeviceInformations() ([]DeviceInformations, error) {
	level.Info(client.logger).Log("msg", "Retrieve device informations")
	var informations []DeviceInformations
	if err := client.apiRequest("/device", &informations); err != nil {
		return nil, err
	}
	return informations, nil
}

// getDeviceCPU returns Bbox CPU information
// See: https://api.bbox.fr/doc/apirouter/#api-Device-GetDeviceCPU
func (client *Client) getDeviceCPU() ([]DeviceCPU, error) {
	level.Info(client.logger).Log("msg", "Retrieve device CPU")
	var cpu []DeviceCPU
	if err := client.apiRequest("/device/cpu", &cpu); err != nil {
		return nil, err
	}
	return cpu, nil
}

// getDeviceMemory returns Bbox Memory information
// See: https://api.bbox.fr/doc/apirouter/#api-Device-GetDeviceMem
func (client *Client) getDeviceMemory() ([]DeviceMemory, error) {
	level.Info(client.logger).Log("msg", "Retrieve device memory")
	var memory []DeviceMemory
	if err := client.apiRequest("/device/mem", &memory); err != nil {
		return nil, err
	}
	return memory, nil
}
