package category

import (
	"context"
	"log"
	pb "shop/proto/proto"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestCreateCategory(t *testing.T) {
	conn, err := grpc.Dial(":8002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	grpcClient := pb.NewCategoryClient(conn)
	ctx := context.Background()
	category := pb.CreateCategory{Name: "Men", IsParent: true, Sort: 0, ParentId: 0}
	res, err := grpcClient.Create(ctx, &category)
	if err != nil {
		panic(err)
	}
	log.Println(res)

}

func TestFindAll(t *testing.T) {
	conn, err := grpc.Dial(":8002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	grpcClient := pb.NewCategoryClient(conn)
	ctx := context.Background()
	res, err := grpcClient.GetAllCategory(ctx, &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	log.Println(res)

}
