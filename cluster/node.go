package cluster

type Node struct {
	Id          int64  `json:"id"`
	Address     string `json:"addr"`
	Port        int    `json:"port"`
	StartUpTime int64  `json:"start-up-time"`
}
