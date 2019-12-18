package domain

/**
 * 基站信息
 */
type BaseStation struct {
	Type            int    `json:"type,omitempty"`            // 所 属 网 络 0: 未 知 ,1:GSM,2:CDMA,3:WCDMA ,4:LTE
	Mcc             int    `json:"mcc,omitempty"`             // 移动国家代码 (GSM,CDMA,WCDMA,LTE)
	Mnc             int    `json:"mnc,omitempty"`             // 移动网络号码 (GSM,CDMA,WCDMA,LTE)
	Lac             int    `json:"lac,omitempty"`             // 位置区域码 (GSM,WCDMA)
	Cid             int    `json:"cid,omitempty"`             // UMTS 小区身份(GSM,WCDMA)
	Tac             int    `json:"tac,omitempty"`             // 跟踪区域码 (LTE)
	Ci              int    `json:"ci,omitempty"`              // 小区标识(LTE)
	Pci             int    `json:"pci,omitempty"`             // 物理小区 id (LTE)
	Psc             int    `json:"psc,omitempty"`             // WCDMA 主扰码(WCDMA)
	Sid             int    `json:"sid,omitempty"`             // CDMA System ID (CDMA)
	Nid             int    `json:"nid,omitempty"`             // Network ID (CDMA)
	Bid             int    `json:"bid,omitempty"`             // Base Station ID (CDMA)
	DBm             int    `json:"dBm,omitempty"`             // 1 毫瓦分贝数 (GSM,CDMA,WCDMA,LTE)
	AsuLevel        int    `json:"asuLevel,omitempty"`        // Asu 信号单元(GSM,CDMA,WCDMA,LTE)
	Level           int    `json:"level,omitempty"`           // 信号格(GSM,CDMA,WCDMA,LTE)
	Timestamp       int64  `json:"timestamp,omitempty"`       // 采集时间戳
	Register        int    `json:"register,omitempty"`        // 是否已注册: 1-注册 null-未注册
	BaseStationType string `json:"baseStationType,omitempty"` // 基站实例类型，区分到底是什么基站实例，也可以此来区分版本
	Earfcn          int    `json:"earfcn,omitempty"`          // Android N 新增，绝对射频频道号(LTE)
	BandWidth       int    `json:"bandWidth,omitempty"`       // 小区带宽(LTE)
	Arfcn           int    `json:"arfcn,omitempty"`           // Android N 新增，绝对射频频道号(GSM)
	Bsic            int    `json:"bsic,omitempty"`            // Android N 新增，基站识别码，物理小区id(GSM)
	Uarfcn          int    `json:"uarfcn,omitempty"`          // Android N 新增，绝对射频频道号(WCDMA)
	Rsrp            int    `json:"rsrp,omitempty"`            // 信号接收质量(LTE)
	Rssnr           int    `json:"rssnr,omitempty"`           // 信噪比(LTE)
	Cqi             int    `json:"cqi,omitempty"`             //信道质量指标(LTE)
	TimingAdvance   int    `json:"timingAdvance,omitempty"`   // 时间提前量(GSM,LTE)
	CsiRsrp         int    `json:"csiRsrp,omitempty"`         //csi信号接收强度(NR)
	CsiRsrq         int    `json:"csiRsrq,omitempty"`         //csi信号接收质量(NR)
	CsiSinr         int    `json:"csiSinr,omitempty"`         //csi信噪比(NR)
	SsRsrp          int    `json:"ssRsrp,omitempty"`          //ss信号接收强度(NR)
	SsRsrq          int    `json:"ssRsrq,omitempty"`          //ss信号接收质量(NR)
	SsSinr          int    `json:"ssSinr,omitempty"`          //ss信噪比(NR)
	Nci             int64  `json:"nci,omitempty"`             //小区标识(NR)
	Mrarfcn         int    `json:"nrarfcn,omitempty"`         //绝对射频频道号(NR)
	BitErrorRate    int    `json:"bitErrorRate,omitempty"`    //误码率(GSM，WCDMA)
	MCdmaDbm        int    `json:"mCdmaDbm,omitempty"`        //电信2G Dbm
	MCdmaEcio       int    `json:"mCdmaEcio,omitempty"`       //电信2G Ec/Io
	MEvdoDbm        int    `json:"mEvdoDbm,omitempty"`        //电信3G Dbm
	MEvdoEcio       int    `json:"mEvdoEcio,omitempty"`       //电信3G Ec/Io
	MEvdoSnr        int    `json:"mEvdoSnr,omitempty"`        //电信3G 信噪比
}
