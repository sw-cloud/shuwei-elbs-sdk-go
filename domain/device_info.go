package domain

/**
 * 设备信息
 */
type DeviceInfo struct {
	Imei    string  `json:"imei,omitempty"`    // Android 系统的设备号
	Idfa    string  `json:"idfa,omitempty"`    // 仅iOS6.0以上系统的IDFA,如4CFD11F0-09D0-4BF3-91CED50600BD0E64
	Andid   string  `json:"andid,omitempty"`   // 用户终端的AndroidID,如9774d56d682e549c
	Device  string  `json:"device,omitempty"`  // 设 备 品 牌 和 型 号 ， 如 honor v8 、Sumsang S6
	Osv     string  `json:"osv,omitempty"`     // 操作系统版本，如 iPhone 8.1.2 的参数填写 8.1.2
	Carrier string  `json:"carrier,omitempty"` // 运营商： 0：unknown 1：CMCC 2：CUCC 3：CTCC
	Imsi    string  `json:"imsi,omitempty"`    // 国际移动用户识别码
	RomInfo string  `json:"romInfo,omitempty"` // 系统romInfo
	Width   float32 `json:"width,omitempty"`   // 屏幕宽度
	Height  float32 `json:"height,omitempty"`  // 屏幕高度
}
