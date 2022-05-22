package flink

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CommonResponse struct {
	*responses.BaseResponse
	RequestId string      `json:"requestId"`
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
}

func NewCommonResponse() (response *CommonResponse) {
	return &CommonResponse{
		BaseResponse: &responses.BaseResponse{},
	}
}

type commonParamsJson struct {
	ParamsJson string `json:"paramsJson"`
}

func setJsonRequest(request *requests.RoaRequest, params interface{}) {
	data, _ := json.Marshal(params)
	paramsJson := commonParamsJson{
		ParamsJson: string(data),
	}
	paramsJsonData, _ := json.Marshal(paramsJson)
	request.SetContentType("application/json")
	request.SetContent(paramsJsonData)
}
