package domain

/**
 * 卫星信息
 */

type SatelliteInfo struct {
	HasEphemeris       bool    `json:"mHasEphemeris"`       //GPS引擎是否有卫星星历(有则返回true)
	HasAlmanac         bool    `json:"mHasAlmanac"`         //GPS引擎是否有近似轨道信息(有则返回true)
	UsedInFix          bool    `json:"mUsedInFix"`          //卫星是否被GPS引擎用于计算最近位置(是则返回true)
	Prn                int     `json:"mPrn"`                //伪随机噪声码
	Snr                float64 `json:"mSnr"`                //卫星的信噪比
	Elevation          float64 `json:"mElevation"`          //卫星的高度角
	Azimuth            float64 `json:"mAzimuth"`            //卫星的方位角
	CarrierFrequencyHz float64 `json:"mCarrierFrequencyHz"` //所跟踪信号的载波频率
}
