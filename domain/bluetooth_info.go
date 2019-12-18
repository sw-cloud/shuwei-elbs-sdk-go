package domain

/**
 * 蓝牙信息
 */

type BluetoothInfo struct {
	Name           string `json:"name,omitempty"`           // 设备名称
	Mac            string `json:"mac,omitempty"`            // MAC 地址，去除冒号分隔符保持小写
	Rssi           int    `json:"rssi,omitempty"`           // 信号强度（dBm）
	IbeaconUUID    string `json:"ibeaconUUID,omitempty"`    // 设备 UUID
	IbeaconMajorId int    `json:"ibeaconMajorId,omitempty"` // 设备 MajorId
	IbeaconMinorId int    `json:"ibeaconMinorId,omitempty"` // 设备 MinorId
	Timestamp      int64  `json:"timestamp,omitempty"`      // 采集时间戳
}
