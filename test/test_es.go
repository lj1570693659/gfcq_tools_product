package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	// 自定义配置
	cfg := elasticsearch.Config{
		// 有多个节点时需要配置
		Addresses: []string{
			"http://10.80.28.218:9201",
		},
		// 配置HTTP传输对象
		Transport: &http.Transport{
			//MaxIdleConnsPerHost 如果非零，控制每个主机保持的最大空闲(keep-alive)连接。如果为零，则使用默认配置2。
			MaxIdleConnsPerHost: 10,
			//ResponseHeaderTimeout 如果非零，则指定在写完请求(包括请求体，如果有)后等待服务器响应头的时间。
			ResponseHeaderTimeout: time.Second,
			//DialContext 指定拨号功能，用于创建不加密的TCP连接。如果DialContext为nil(下面已弃用的Dial也为nil)，那么传输拨号使用包网络。
			DialContext: (&net.Dialer{Timeout: time.Second}).DialContext,
			// TLSClientConfig指定TLS.client使用的TLS配置。
			//如果为空，则使用默认配置。
			//如果非nil，默认情况下可能不启用HTTP/2支持。
			TLSClientConfig: &tls.Config{
				MaxVersion: tls.VersionTLS11,
				//InsecureSkipVerify 控制客户端是否验证服务器的证书链和主机名。
				InsecureSkipVerify: true,
			},
		},
	}

	// 创建一个 Elasticsearch 客户端
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	fmt.Println("es-------------", es)

	// 在索引中创建或更新文档。
	//res, err := es.Index(
	//	"test",                                   // Index name
	//	strings.NewReader(`{"title" : "Test2"}`), // Document body
	//	//es.Index.WithDocumentID("1"),            // Document ID
	//	//es.Index.WithRefresh("true"),               // Refresh
	//)
	//fmt.Println("res-------------", res)

	req := esapi.SearchRequest{
		Index:        []string{"test"},
		DocumentType: []string{"_doc"},
		Body: strings.NewReader(`{
		  "query": {
			"match": {
			  "title": "Test2"
			}
		  }
		}`),
	}
	getRes, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer getRes.Body.Close()

	fmt.Println("getRes-------------", getRes)

	//D:\code\product\data\elasticsearch\kibana\config\kibana.yml
	//docker run -d -p 5601:5601 -v D:\code\product\data\elasticsearch\kibana\config\kibana.yml:/usr/share/kibana/config/kibana.yml --restart=always --name kibana docker.io/library/kibana:7.6.1

	//./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.6.1/elasticsearch-analysis-ik-7.6.1.zip
	///usr/share/elasticsearch/plugins/ik

	//SELECT * from MBP_META_OBJCLS_PROPVALUE WHERE COBJCLASSID = 'cd6f924e-d1ae-4e39-a799-df28346dd183' AND CPROPERTYID = 'e59a9030-a40a-480e-b6a0-79db5310b244'
}
