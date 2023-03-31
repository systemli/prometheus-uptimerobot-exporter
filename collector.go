package main

import (
	"log"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	Client  *UptimerobotClient
	Monitor *prometheus.Desc
}

func NewCollector(apiKey string) *Collector {
	return &Collector{
		Client: NewUptimerobotClient(apiKey),
		Monitor: prometheus.NewDesc(
			"uptimerobot_monitor_up",
			"Status of the UptimeRobot monitor",
			[]string{"id", "friendly_name", "url", "type", "status"},
			nil,
		),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Monitor
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	monitors, err := c.Client.GetMonitors()
	if err != nil {
		log.Print(err)
		return
	}

	for _, monitor := range monitors {
		ch <- prometheus.MustNewConstMetric(
			c.Monitor,
			prometheus.GaugeValue,
			float64(monitor.Status),
			strconv.Itoa(monitor.ID),
			monitor.FriendlyName,
			monitor.URL,
			strconv.Itoa(monitor.Type),
			strconv.Itoa(monitor.Status),
		)
	}
}
