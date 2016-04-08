package consul

import "github.com/influxdata/telegraf"

// Consul base struct
type Consul struct {
}

// SampleConfig returns the sample Telegraf config for Consul
func (c *Consul) SampleConfig() string {
	return ""
}

// Description returns a description of the plugin
func (c *Consul) Description() string {
	return ""
}

// Gather retrieves metrics and passes them on to the Accumulator
func (c *Consul) Gather(acc telegraf.Accumulator) error {
	return nil
}
