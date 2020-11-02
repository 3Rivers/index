package main

import (
	"github.com/3Rivers/index/handler/order"
	log "github.com/micro/go-micro/v2/logger"
	"net/http"
	"time"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/3Rivers/index/handler"
	"github.com/micro/go-micro/v2/web"
)

var etcdReg registry.Registry

func  init()  {
        //新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
        etcdReg = etcd.NewRegistry(
                registry.Addrs("192.168.2.254:12379", "192.168.2.254:22379", "192.168.2.254:32379"),
        )
}

func main() {
	// create new web service
        service := web.NewService(
                web.Name("go.micro.web.index"),
                web.Version("latest"),
                web.Registry(etcdReg),
                web.RegisterTTL(time.Second * 30),
                web.RegisterInterval(time.Second * 3),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/hello", handler.IndexCall)
    service.HandleFunc("/order", order.OrderCall)


	// run service
	if err := service.Run(); err != nil {
	        log.Fatal(err)
	}
}
