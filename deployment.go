package flink

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

// 获取Deployment

// 更新Deployment状态
type UpdateDeploymentDesiredStateRequest struct {
	*requests.RoaRequest
	Workspace    string                             `position:"Path" name:"workspace"`
	Namespace    string                             `position:"Path" name:"namespace"`
	DeploymentId string                             `position:"Path" name:"deploymentId"`
	Params       UpdateDeploymentDesiredStateParams `json:"params"`
}

type UpdateDeploymentDesiredStateParams struct {
	State State `json:"state"`
}

// 更新Deployment状态
type UpdateDeploymentDesiredStateResp struct {
	*CommonResponse
}

// 获取Deployments列表
type ListDeploymentsRequest struct {
	*requests.RoaRequest
	Workspace string           `position:"Path" name:"workspace"`
	Namespace string           `position:"Path" name:"namespace"`
	Creator   string           `position:"Query" name:"creator"`
	Modifier  string           `position:"Query" name:"modifier"`
	Priority  string           `position:"Query" name:"priority"`
	BatchMode string           `position:"Query" name:"batchMode"`
	SortName  string           `position:"Query" name:"sortName"`
	PageSize  requests.Integer `position:"Query" name:"pageSize"`
	PageIndex requests.Integer `position:"Query" name:"pageIndex"`
	Name      string           `position:"Query" name:"name"`
	SortOrder string           `position:"Query" name:"sortOrder"`
	State     string           `position:"Query" name:"state"`
	Status    string           `position:"Query" name:"status"`
}

type ListDeploymentsResp struct {
	*CommonResponse
	Data struct {
		Metadata struct {
			TotalSize int `json:"totalSize"`
			TotalPage int `json:"totalPage"`
		} `json:"metadata"`
		ApiVersion string       `json:"apiVersion"`
		Kind       string       `json:"kind"`
		Items      []Deployment `json:"items"`
	} `json:"data"`
}

type GetDeploymentRequest struct {
	*requests.RoaRequest
	Workspace    string `position:"Path" name:"workspace"`
	Namespace    string `position:"Path" name:"namespace"`
	DeploymentId string `position:"Path" name:"deploymentId"`
	Name         string
}

type GetDeploymentResp struct {
	*CommonResponse
	Data Deployment `json:"data"`
}

type UpdateDeploymentRequest struct {
	*requests.RoaRequest
	Workspace    string `position:"Path" name:"workspace"`
	Namespace    string `position:"Path" name:"namespace"`
	DeploymentId string `position:"Path" name:"deploymentId"`
	Params       Deployment
}

type UpdateDeploymentResp struct {
	*CommonResponse
	Data Deployment `json:"data"`
}

func (c *Client) ListDeployments(request *ListDeploymentsRequest) (response *ListDeploymentsResp, err error) {
	request.RoaRequest = &requests.RoaRequest{}
	request.InitWithApiInfo("ververica", "2020-05-01", "ListDeployments",
		"/pop/workspaces/[workspace]/api/v1/namespaces/[namespace]/deployments", "ververica", "openAPI")
	request.Method = requests.GET
	response = &ListDeploymentsResp{
		CommonResponse: NewCommonResponse(),
	}
	err = c.Client.DoAction(request, response)
	return
}

func (c *Client) UpdateDeploymentDesiredState(request *UpdateDeploymentDesiredStateRequest) (response *UpdateDeploymentDesiredStateResp, err error) {
	request.RoaRequest = &requests.RoaRequest{}
	request.InitWithApiInfo("ververica",
		"2020-05-01",
		"UpdateDeploymentDesiredState",
		"/pop/workspaces/[workspace]/api/v1/namespaces/[namespace]/deployments/[deploymentId]/state",
		"ververica", "openAPI")
	request.Method = requests.PUT
	setJsonRequest(request.RoaRequest, request.Params)

	response = &UpdateDeploymentDesiredStateResp{
		CommonResponse: NewCommonResponse(),
	}
	err = c.Client.DoAction(request, response)
	return
}

func (c *Client) GetDeployment(request *GetDeploymentRequest) (response *GetDeploymentResp, err error) {
	request.RoaRequest = &requests.RoaRequest{}
	request.InitWithApiInfo("ververica",
		"2020-05-01",
		"GetDeployment",
		"/pop/workspaces/[workspace]/api/v1/namespaces/[namespace]/deployments/[deploymentId]",
		"ververica", "openAPI")
	request.Method = requests.GET
	response = &GetDeploymentResp{
		CommonResponse: NewCommonResponse(),
	}
	err = c.Client.DoAction(request, response)
	return
}

func (c *Client) UpdateDeployment(request *UpdateDeploymentRequest) (response *UpdateDeploymentResp, err error) {
	request.RoaRequest = &requests.RoaRequest{}
	request.InitWithApiInfo("ververica",
		"2020-05-01",
		"GetDeployment",
		"/pop/workspaces/[workspace]/api/v1/namespaces/[namespace]/deployments/[deploymentId]/patch",
		"ververica", "openAPI")
	request.Method = requests.PUT
	setJsonRequest(request.RoaRequest, request.Params)

	response = &UpdateDeploymentResp{
		CommonResponse: NewCommonResponse(),
	}
	err = c.Client.DoAction(request, response)
	return
}
