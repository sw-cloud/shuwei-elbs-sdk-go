package constant

type RetCode struct {
	RetCode int
	Msg     string
}

var (
	/** 成功 */
	SUCCESS = RetCode{0, "OK"}
	/** 服务器错误 */
	SERVER_ERROR = RetCode{100001, "System error"}
	/** 环境错误 */
	ENVIRONMENT_ERROR = RetCode{100002, "ELBSProfile error"}
	/** 缺少必传参数 */
	LACK_NEED_PARAMS = RetCode{100003, "Miss required parameter"}
	/** 参数格式错误 */
	PARAMS_FORMAT_ERROR = RetCode{100004, "Parameter value invalid"}
	/** 签名错误 */
	SIGN_ERROR = RetCode{100005, "Sign error"}
	/** 包名错误 */
	PACKAGE_NAME_ERROR = RetCode{100006, "Package name error"}
	/** 请求数据为空 */
	REQUEST_DATA_IS_NULL = RetCode{100007, "Request data is null"}
	/** 头部缺少授权数据 */
	LACK_AUTHORIZATION = RetCode{100008, "Restricted access.Lack authorization"}
	/** 达到qps上限 */
	APP_TPS_REACH_LIMIT = RetCode{100009, "Restricted access.App TPS reach limit"}
	/** 每日配额达到上限 */
	REACH_DAILY_TPD_LIMIT = RetCode{100010, "Restricted access.Transaction reach daily limit"}
	/** Server tps 达到上限 */
	SERVER_TPS_REACH_LIMIT = RetCode{100011, "Restricted access.TPS reach limit"}
	/** 接口版本错误 */
	API_VERSION_ERROR = RetCode{100012, "Api version is wrong"}
	/** 没有权限执行此操作 */
	NO_AUTHORITY = RetCode{100016, "No permission to perform this operation"}
	/** 信号无效**/
	SIGNAL_INVALID = RetCode{100018, "Signal invalid"}
	/** ssidEncode 错误**/
	SSID_ENCODE_VALUE_ERROR = RetCode{100019, "SsidEncode value error"}
	/** 无定位结果 */
	SCENE_NO_RESULT = RetCode{201001, "No result"}
)
