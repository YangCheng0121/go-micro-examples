package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"go-micro-examples/registerConfiguration/config"
	"go-micro-examples/registerConfiguration/handler"
	pb "go-micro-examples/registerConfiguration/proto"
)

func main() {
	// Register consul
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})

	// 配置中心
	consulConfig, err := config.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		logger.Fatal(err)
	}

	// Mysql配置信息
	mysqlInfo, err := config.GetMysqlFromConsul(consulConfig, "mysql")
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Mysql配置信息:", mysqlInfo)

	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.srv.registerconfiguration"),
		micro.Version("latest"),
		// 注册consul中心
		micro.Registry(reg),
	)

	// Register handler
	if err := pb.RegisterRegisterConfigurationHandler(srv.Server(), new(handler.RegisterConfiguration)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
