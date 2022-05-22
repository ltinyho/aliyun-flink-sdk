package flink

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-hangzhou": "ververica.cn-hangzhou.aliyuncs.com",
			"cn-shanghai": "ververica.cn-shanghai.aliyuncs.com",
			"cn-shenzhen": "ververica.cn-shenzhen.aliyuncs.com",
			"cn-beijing":  "ververica.cn-beijing.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
