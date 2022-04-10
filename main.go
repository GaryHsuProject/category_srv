package main

import (
	"fmt"
	"net"
	"shop/driver"
	"shop/internal/category"
	pb "shop/proto/proto"

	"google.golang.org/grpc"
)

func main() {

	driver.Init()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", driver.GlobalConfig.Port))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	repo := category.NewCategoryRepository(driver.DB)
	pb.RegisterCategoryServer(grpcServer, category.NewCategoryService(repo))
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	driver.Logger.Info("grpc server is running on :", driver.GlobalConfig.Port)
}
