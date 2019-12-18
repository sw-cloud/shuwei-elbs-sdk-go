package domain

/**
 * wifi信号
 */
type WifiItem struct {
	Bssid        string `json:"bssid,omitempty"`        // mac
	Rssi         int    `json:"rssi,omitempty"`         // 信号强度
	Ssid         string `json:"ssid,omitempty"`         // 名称
	ConnectState int    `json:"connectState,omitempty"` // 是否为连接信号（1表示连接，默认为null）
	Timestamp    int64  `json:"timestamp,omitempty"`    // wifi采集的时间戳(ms)
	Band         int    `json:"band,omitempty"`         //频段（1:2.4G，2:5G）
	Channel      int    `json:"channel,omitempty"`      //信道
}
