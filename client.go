package shuwei_elbs_sdk_go

import (
	"encoding/base64"
	"encoding/json"
	"github.com/sw-cloud/shuwei-elbs-sdk-go/constant"
	"github.com/sw-cloud/shuwei-elbs-sdk-go/domain"
	"github.com/sw-cloud/shuwei-elbs-sdk-go/log"
	"github.com/sw-cloud/shuwei-elbs-sdk-go/utils"
	"strings"
)

const (
	// 加密key的长度
	LENGTH_OF_ENCRYPT_KEY = 16
)

type IELBSClient interface {
	Location(elbsRequest *domain.ELBSRequest) *domain.ELBSResponse
}

type ELBSClient struct {
	elbsProfile *ELBSProfile
}

func NewELBSClient(profile *ELBSProfile) IELBSClient {
	return &ELBSClient{
		elbsProfile: profile,
	}
}

func (this *ELBSClient) Location(elbsRequest *domain.ELBSRequest) (locateResp *domain.ELBSResponse) {

	if elbsRequest == nil || this.elbsProfile == nil {
		return domain.NewELBSResponse(constant.REQUEST_DATA_IS_NULL)
	}
	//log.Infof("requestInfo|elbsRequest=%v|profile=%v\n", utils.ToJsonStr(elbsRequest), utils.ToJsonStr(this.elbsProfile))

	if this.elbsProfile.LackNeedParams() || elbsRequest.LackNeedParams() {
		log.Errorf("doResponse|%v\n", constant.LACK_NEED_PARAMS.RetCode)
		return domain.NewELBSResponse(constant.LACK_NEED_PARAMS)
	}

	if !this.elbsProfile.FormatCheck() || !elbsRequest.FormatCheck() {
		log.Errorf("doResponse| %v\n", constant.PARAMS_FORMAT_ERROR.RetCode)
		return domain.NewELBSResponse(constant.PARAMS_FORMAT_ERROR)
	}

	locationRequest := domain.NewLocationFromRequestData(elbsRequest, this.elbsProfile.AppId, this.elbsProfile.AppKey)
	requestBody, err := json.Marshal(locationRequest)
	if err != nil {
		log.Errorf("locationRequest to json: %v\n", err.Error())
		return domain.NewELBSResponse(constant.PARAMS_FORMAT_ERROR)
	}

	appKeyTmp := strings.ReplaceAll(this.elbsProfile.AppKey, "-", "")
	sKey := appKeyTmp[0:LENGTH_OF_ENCRYPT_KEY]
	ivP := appKeyTmp[LENGTH_OF_ENCRYPT_KEY : LENGTH_OF_ENCRYPT_KEY*2]

	encryptedBody, err := utils.AesCBCEncrypt(requestBody, []byte(sKey), []byte(ivP))
	if err != nil {
		log.Errorf("encrypt request data error: %v\n", err)
		return domain.NewELBSResponse(constant.PARAMS_FORMAT_ERROR)
	}

	respBody, err := utils.HttpPost(this.elbsProfile.Url, []byte(base64.StdEncoding.EncodeToString(encryptedBody)), locationRequest.Authorization)
	if err != nil {
		log.Errorf("httpPost|%v\n", constant.SERVER_ERROR.RetCode)
		return domain.NewELBSResponse(constant.SERVER_ERROR)
	}
	log.Infof("httpPost|%v\n", string(respBody))

	response := domain.ELBSResponse{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		data, err := base64.StdEncoding.DecodeString(string(respBody))
		if err != nil {
			log.Errorf("base64 decode|%v\n", err)
			return domain.NewELBSResponse(constant.SERVER_ERROR)
		}

		decryptedData, err := utils.AesCBCDecrypt(data, []byte(sKey), []byte(ivP))
		if err != nil {
			log.Errorf("decrypted |%v\n", err)
			return domain.NewELBSResponse(constant.SERVER_ERROR)
		}
		err = json.Unmarshal(decryptedData, &response)
		if err != nil {
			log.Errorf("decrypted data|%v\n", decryptedData)
			log.Errorf("decrypted data raw|%v\n", string(decryptedData))
			return domain.NewELBSResponse(constant.SERVER_ERROR)
		}
	}
	return &response
}
