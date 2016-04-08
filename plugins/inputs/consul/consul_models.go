package Consul

const meas = "consul"

type Event interface {
	NewMetric() telegraf.Metric
}

type Index struct {
	CreateIndex int `json:"CreateIndex"`
	ModifyIndex int `json:"ModifyIndex"`
}

type Datacenter string

type Node struct {
	Name    string `json:"Node"`
	Address string `json:"Address"`
}

type Service struct {
	ServiceID               string   `json:"ServiceID"`
	ServiceName             string   `json:"ServiceName"`
	ServiceTags             []string `json:"ServiceTags"`
	ServiceAddress          string   `json:"ServiceAddress"`
	ServicePort             int      `json:"ServicePort"`
	ServiceEnableTagOveride bool     `json:"ServiceEnableTagOveride"`
}

type Svc struct {
	ID                string   `json:"ID"`
	Name              string   `json:"Service"`
	Tags              []string `json:"Tags"`
	Address           string   `json:"Address"`
	Port              int      `json:"Port"`
	EnableTagOverride bool     `json:"EnableTagOverride"`
}

type Member struct {
	Name        string            `json:"Name"`
	Address     string            `json:"Addr"`
	Port        int               `json:"Port"`
	Tags        map[string]string `json:"Tags"`
	Status      int               `json:"Status"`
	ProtocolMin int               `json:"ProtocolMin"`
	ProtocolMax int               `json:"ProtocolMax"`
	ProtocolCur int               `json:"ProtocolCur"`
	DelegateMin int               `json:"DelegeateMin"`
	DelegateMax int               `json:"DelegeateMax"`
	DelegateCur int               `json:"DelegeateCur"`
}
