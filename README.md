aliyun 实时计算版 golang sdk

# 封装过程

1. client 初始化的过程参考 [foas sdk](https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/services/foas/client.go)
2. api参考 [java 版sdk](./Demo.bak)封装,[官网链接](https://help.aliyun.com/document_detail/194744.html).
   [java 版本 1.0.6](https://search.maven.org/artifact/com.aliyun/aliyun-java-sdk-ververica/1.0.6/jar)
3. api的 host 是从java版sdk中抓取处理的,设置 HttpUtil.setIsHttpDebug(true); 可以查看请求的详细信息.
