package flink

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type DeleteArtifactRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
	Namespace string `position:"Path" name:"namespace"`
	Filename  string `position:"Query" name:"filename"`
}
type DeleteArtifactRequestResp struct {
	*CommonResponse
}

func (c Client) DeleteArtifact(req *DeleteArtifactRequest) (response *DeleteArtifactRequestResp, err error) {
	request := &DeleteArtifactRequest{
		RoaRequest: &requests.RoaRequest{},
		Workspace:  req.Workspace,
		Namespace:  req.Namespace,
		Filename:   req.Filename,
	}
	request.InitWithApiInfo("ververica",
		"2020-05-01",
		"DeleteArtifact",
		"/pop/workspaces/[workspace]/artifacts/v1/namespaces/[namespace]/artifacts:delete",
		"ververica",
		"openAPI")
	request.Method = requests.DELETE
	response = &DeleteArtifactRequestResp{
		CommonResponse: NewCommonResponse(),
	}
	err = c.Client.DoAction(request, response)
	return
}
