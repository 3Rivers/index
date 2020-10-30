package main

import (
        log "github.com/micro/go-micro/v2/logger"
	      "net/http"
        "github.com/micro/go-micro/v2/web"
        "github.com/3Rivers/index/handler"
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
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/index/call", handler.IndexCall)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
