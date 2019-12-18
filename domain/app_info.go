package domain

/**
 * 应用信息
 */
type AppInfo struct {
	ApplicationId string `json:"applicationId,omitempty"` // 应用的 PackageName
	VersionCode   string `json:"versionCode,omitempty"`   // 版本号
	FirstTime     int64  `json:"firstTime,omitempty"`     // 首次安装时间 ms
	LastTime      int64  `json:"lastTime,omitempty"`      // 最新更新时间 ms
	TotalTime     int64  `json:"totalTime,omitempty"`     // App前台使用时长
}
