package domain

/*
**
* gps信息
 */
type GpsInfo struct {
	SatelliteInfo            []SatelliteInfo `json:"satelliteInfo,omitempty"`             //卫星信息
	Provider                 string          `json:"mProvider,omitempty"`                 //获取是网络定位、GPS定位还是被动定位
	HorizontalAccuracyMeters float64         `json:"mHorizontalAccuracyMeters,omitempty"` //获取此位置的估计精确度
	Altitude                 float64         `json:"mAltitude,omitempty"`                 //海拔高度
	Bearing                  float64         `json:"mBearing,omitempty"`                  //方向度数（0,360）表示正北偏东多少度（0->北，90->东，180->南，270->西）
	ElapsedRealtimeNanos     int64           `json:"mElapsedRealtimeNanos,omitempty"`     //获取时间（单位：纳秒）
	Speed                    float64         `json:"mSpeed,omitempty"`                    //自传速度（单位：m/s）
	SatelliteCount           int             `json:"mSatelliteCount,omitempty"`           //扫描到的卫星总数
	Timestamp                int64           `json:"timestamp,omitempty"`                 // 采集时间戳
}
