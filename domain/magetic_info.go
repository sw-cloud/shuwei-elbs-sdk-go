package domain

/**
 * 移动设备所在位置的磁场信息
 */
type MagneticInfo struct {
	MagX          float64 `json:"magX,omitempty"`          // X轴方向的磁感应强度
	MagY          float64 `json:"magY,omitempty"`          // Y轴方向的磁感应强度
	MagZ          float64 `json:"magZ,omitempty"`          // Z轴方向的磁感应强度
	SensorName    string  `json:"sensorName,omitempty"`    // 传感器名称
	SensorVendor  string  `json:"sensorVendor,omitempty"`  // 传感器供应商
	SensorVersion string  `json:"sensorVersion,omitempty"` // 传感器版本
	Timestamp     int64   `json:"timestamp,omitempty"`     // 采集时间戳
}
