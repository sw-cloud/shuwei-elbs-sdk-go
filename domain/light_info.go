package domain

/**
 * 光强信息
 */

type LightInfo struct {
	Lux           float64 `json:"lux,omitempty"`           //光强信息
	SensorName    string  `json:"sensorName,omitempty"`    //传感器名称
	SensorVendor  string  `json:"sensorVendor,omitempty"`  //传感器供应商
	SensorVersion string  `json:"sensorVersion,omitempty"` //传感器版本
	Timestamp     int64   `json:"timestamp,omitempty"`     //采集时间戳
}
