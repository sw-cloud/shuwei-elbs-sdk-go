package domain

/**
 * 坐标信息API请求参数
 */
type CoordinateInfo struct {
	CoordinateSystem string  `json:"coordinateSystem,omitempty"` // 坐标系
	Longitude        float64 `json:"longitude,omitempty"`        // 经度
	Latitude         float64 `json:"latitude,omitempty"`         // 纬度
	Timestamp        int64   `json:"timestamp,omitempty"`        // 采集时间戳
}
