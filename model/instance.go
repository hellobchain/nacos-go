package model

type Instance struct {
	ID          uint64  `json:"id"`
	IP          string  `json:"ip"`
	Port        uint64  `json:"port"`
	ServiceName string  `json:"serviceName"`
	GroupName   string  `json:"groupName"`
	ClusterName string  `json:"clusterName"`
	Weight      float64 `json:"weight"`
	Healthy     bool    `json:"healthy"`
	Ephemeral   bool    `json:"ephemeral"`
	Metadata    string  `json:"metadata"` // json string
	ExpireTime  int64   `json:"expireTime"`
}
