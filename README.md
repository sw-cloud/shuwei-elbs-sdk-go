# 使用示例
以下是调用数位ELBS定位接口的4个主要步骤：
1. 创建ELBSProfile实例并进行初始化  
2. 创建ELBSClient实例化  
3. 创建一个ELBSRequest请求结构体并设置值  
4. 通过client.Location()方法调用定位接口并接受定位结果  

简单代码示例如下：

```
package main

import (
	"encoding/json"
	"fmt"
	sdk "github.com/sw-cloud/shuwei-elbs-sdk-go"
	"github.com/sw-cloud/shuwei-elbs-sdk-go/domain"
	"time"
)

const (
	//这里填上实际的值
	APPID  = ""
	APPKEY = ""
	URL    = ""
)

var signals = `[{"bssid":"8f:85:f3:44:55:a0","rssi":-51,"ssid":"ChinaNet-pAZg"},{"bssid":"70:51:f3:78:55:8a","rssi":-75,"ssid":"qianweishuan2G"},{"bssid":"15:91:a3:78:55:8a","rssi":-80,"ssid":"ITTS"},{"bssid":"22:91:f1:78:55:8a","rssi":-82,"ssid":"QWS017"},{"bssid":"90:51:f3:78:55:8a","rssi":-84,"ssid":"qinxiaoguan"}]`

func main() {
	client := sdk.NewELBSClient(sdk.NewELBSProfile(APPID, APPKEY, URL))
	oid := "" //这里填实际的值
	request := &domain.ELBSRequest{
		Timestamp:   int64(time.Now().Unix()) * 1000,
		Oid:         oid,
		OidType:     "1",
		NetworkMode: "1",
		Signals:     make([]*domain.WifiItem, 0),
	}
	json.Unmarshal([]byte(signals), &request.Signals)

	resp := client.Location(request)
	fmt.Printf("response| %#v\n", resp)
}

```


