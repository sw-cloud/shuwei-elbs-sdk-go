package domain

import "github.com/sw-cloud/shuwei-elbs-sdk-go/constant"

/**
 * ELBS响应信息
 */

type ELBSResponse struct {
	RetCode int                    `json:"retCode"` //响应码
	Msg     string                 `json:"msg"`     //响应信息
	Data    map[string]interface{} `json:"data"`    //定位信息
}

func NewELBSResponse(retCode constant.RetCode) *ELBSResponse {
	resp := ELBSResponse{
		RetCode: retCode.RetCode,
		Msg:     retCode.Msg,
	}
	return &resp
}

func (r *ELBSResponse) SetData(data map[string]interface{}) {
	r.Data = data
}
