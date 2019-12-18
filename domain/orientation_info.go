package domain

/**
 * 移动设备的姿态方向信息
 */
type OrientationInfo struct {
	Pitch         float64 `json:"pitch,omitempty"`         //俯仰角，绕X轴产生的角
	Roll          float64 `json:"roll,omitempty"`          //翻转角，绕y轴产生的角
	Azimuth       float64 `json:"azimuth,omitempty"`       //方向角，绕Z轴产生的角
	SensorName    string  `json:"sensorName,omitempty"`    //传感器名称
	SensorVendor  string  `json:"sensorVendor,omitempty"`  //传感器供应商
	SensorVersion string  `json:"sensorVersion,omitempty"` //传感器版本
	Timestamp     int64   `json:"timestamp,omitempty"`     //采集时间戳
}
