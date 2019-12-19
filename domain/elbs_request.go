package domain

import (
	"github.com/sw-cloud/shuwei-elbs-sdk-go/utils"
	"strconv"
)

type ELBSRequest struct {
	Timestamp       int64            `json:"timestamp"`                 // 时间戳,毫秒
	Mac             string           `json:"mac,omitempty"`             // 终端网卡的MAC地址
	MacType         string           `json:"macType,omitempty"`         // 终端网卡的MAC地址类型
	Oid             string           `json:"oid"`                       // 手机标识
	OidType         string           `json:"oidType"`                   // 手机标识类型 1-IMEI号 4-iosIDFA
	NetworkMode     string           `json:"networkMode"`               // 联网方式 1-wifi 2-2G 3-3G 4-4G
	OsType          int              `json:"osType,omitempty"`          // 操作系统类型：0-不确定 1-Android 2-iOS
	AppChannel      int              `json:"appChannel,omitempty"`      // 应用流量来源的渠道 1-APP
	Ext             string           `json:"ext,omitempty"`             // 扩展参数
	Passthrough     string           `json:"passthrough,omitempty"`     // 透传信息,在响应时原样返回
	Signals         []*WifiItem      `json:"signals"`                   // wifi列表
	BaseStations    []*BaseStation   `json:"baseStations,omitempty"`    // 基站信息
	DeviceInfo      *DeviceInfo      `json:"deviceInfo,omitempty"`      // 设备信息
	AppInfo         []*AppInfo       `json:"appInfo,omitempty"`         // 应用信息
	BluetoothInfo   []*BluetoothInfo `json:"bluetoothInfo,omitempty"`   // 蓝牙信息
	MagneticInfo    *MagneticInfo    `json:"magneticInfo,omitempty"`    // 磁场信息
	OrientationInfo *OrientationInfo `json:"orientationInfo,omitempty"` // 姿态方向信息
	LightInfo       *LightInfo       `json:"lightInfo,omitempty"`       // 光强信息
	GpsInfo         *GpsInfo         `json:"gpsInfo,omitempty"`         // GPS信息
	CoordinateInfo  *CoordinateInfo  `json:"coordinateInfo,omitempty"`  // 经纬度信息
}

func (r *ELBSRequest) LackNeedParams() bool {
	return r.Timestamp == 0 ||
		utils.IsEmptyString(r.Oid) ||
		utils.IsEmptyString(r.OidType) ||
		utils.IsEmptyString(r.NetworkMode) ||
		len(r.Signals) == 0
}

func (r *ELBSRequest) FormatCheck() bool {
	return len(strconv.FormatInt(r.Timestamp, 10)) == 13
}
