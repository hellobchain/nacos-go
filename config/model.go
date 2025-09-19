package config

type ConfigItem struct {
	DataId  string `json:"dataId"`
	Group   string `json:"group"`
	Tenant  string `json:"tenant,omitempty"`
	Content string `json:"content"`
	Md5     string `json:"md5"`
	Type    string `json:"type"`
	BetaIps string `json:"betaIps"`
	SrcUser string `json:"srcUser"`
	SrcIp   string `json:"srcIp"`
	AppName string `json:"appName"`
}
