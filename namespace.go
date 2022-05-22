package flink

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type ListNamespacesRequest struct {
	*requests.RoaRequest
}

type ListNamespacesResp struct {
	*CommonResponse
	Data struct {
		Namespaces []struct {
			Name                      string        `json:"name"`
			Workspace                 string        `json:"workspace"`
			CreateTime                string        `json:"createTime"`
			RoleBindings              []interface{} `json:"roleBindings"`
			LifecyclePhase            string        `json:"lifecyclePhase"`
			PreviewSessionClusterName string        `json:"previewSessionClusterName"`
		} `json:"namespaces"`
	} `json:"data"`
}

func (c *Client) ListNamespaces() (response *ListNamespacesResp, err error) {
	request := &ListNamespacesRequest{
		&requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "ListNamespaces", ""+
		"/pop/namespaces/v1/namespaces", "ververica", "openAPI")
	request.Method = requests.GET
	response = &ListNamespacesResp{
		CommonResponse: NewCommonResponse(),
	}
	err = c.Client.DoAction(request, response)
	return
}
