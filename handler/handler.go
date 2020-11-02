package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	"net/http"
	"time"

	hystrixGo "github.com/afex/hystrix-go/hystrix"

	index "github.com/3Rivers/helloworld/proto/helloworld"
)

var etcdReg registry.Registry

func  init()  {
	//新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
	etcdReg = etcd.NewRegistry(
		registry.Addrs("192.168.2.254:12379", "192.168.2.254:22379", "192.168.2.254:32379"),
	)
}

func IndexCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println("new")

	app := micro.NewService(
		micro.Name("go.micro.client.helloworld"),
		micro.Registry(etcdReg),
		micro.WrapClient(
			// 引入hystrix包装器
			hystrix.NewClientWrapper(),
		),
	)
	// call the backend service
	indexClient := index.NewHelloworldService("go.micro.service.helloworld", app.Client())
	hystrixGo.ConfigureCommand("go.micro.service.helloworld",
		hystrixGo.CommandConfig{
			MaxConcurrentRequests: 50, //最大并发数
			Timeout:               1000,//超时时间
		})
	rsp, err := indexClient.Call(context.TODO(), &index.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
