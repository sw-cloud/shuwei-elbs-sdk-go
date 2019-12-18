package domain

import (
	"encoding/base64"
	"fmt"
	"shuwei-elbs-sdk-go/utils"
	"strconv"
	"strings"
)

type LocationRequest struct {
	ELBSRequest
	AppId         string `json:"appId"`      // appId
	SsidEncode    int    `json:"ssidEncode"` // ssid是否做base64加密：0-不加密,1-加密
	Sign          string `json:"sign"`       // 签名
	Authorization string `json:"-"`
}

func NewLocationFromRequestData(requestData *ELBSRequest, appId, appKey string) *LocationRequest {
	locationRequest := LocationRequest{
		AppId: appId,
	}
	locationRequest.ELBSRequest = *requestData
	locationRequest.Authorization = fmt.Sprintf("t=%d;a=%s", requestData.Timestamp, appId)

	locationRequest.handleSignals(requestData)
	locationRequest.handleSign(appKey)
	return &locationRequest
}

func (r *LocationRequest) handleSignals(requestData *ELBSRequest) {
	signals := requestData.Signals
	newSignals := make([]*WifiItem, 0)

	if len(signals) == 0 {
		return
	}

	for _, wifiItem := range signals {
		ssid := strings.TrimSpace(wifiItem.Ssid)
		if ssid == "" {
			continue
		}
		wifiItem.Ssid = base64.StdEncoding.EncodeToString([]byte(ssid))
		newSignals = append(newSignals, wifiItem)
	}
	r.Signals = newSignals
	r.SsidEncode = 1
}

func (r *LocationRequest) handleSign(appKey string) {
	r.Sign = utils.Sign(r.signParams(), appKey)
}

func (r *LocationRequest) signParams() map[string]string {
	params := make(map[string]string)
	params["timestamp"] =  strconv.FormatInt(r.Timestamp, 10)
	params["oid"] = r.Oid
	params["oidType"] = r.OidType
	params["networkMode"] = r.NetworkMode
	params["appId"] = r.AppId
	params["mac"] = r.Mac
	params["macType"] = r.MacType
	if r.SsidEncode > 0 {
		params["ssidEncode"] = strconv.Itoa(r.SsidEncode)
	}
	if r.OsType > 0 {
		params["osType"] = strconv.Itoa(r.OsType)
	}
	if r.AppChannel > 0 {
		params["appChannel"] = strconv.Itoa(r.AppChannel)
	}
	params["passthrough"] = r.Passthrough
	params["ext"] = r.Ext
	return params
}

