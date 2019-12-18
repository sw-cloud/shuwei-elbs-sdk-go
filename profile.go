package shuwei_elbs_sdk_go

import (
	"regexp"
	"shuwei-elbs-sdk-go/constant"
	"shuwei-elbs-sdk-go/utils"
)

type ELBSProfile struct {
	AppId  string
	AppKey string
	Url    string
}

func NewELBSProfile(appId, appKey, url string) *ELBSProfile {
	return &ELBSProfile{
		AppId:  appId,
		AppKey: appKey,
		Url:    url}
}

func (p *ELBSProfile) LackNeedParams() bool {
	return utils.IsEmptyString(p.AppId) || utils.IsEmptyString(p.AppKey) || utils.IsEmptyString(p.Url)
}

func (p *ELBSProfile) FormatCheck() bool {
	match, err := regexp.MatchString(constant.PATTERN_APPID, p.AppId)
	if !match || err != nil {
		return false
	}
	match, err = regexp.MatchString(constant.PATTERN_APPKEY, p.AppKey)
	if !match || err != nil {
		return false
	}
	return true
}
