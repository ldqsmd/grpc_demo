package srv

import (
	"context"
	"fmt"
	pb "grpc_demo/proto"
)

type SearchService struct { }

func (this *SearchService)Search(ctx context.Context, in *pb.SearchRequest)(out *pb.SearchResponse,err error) {

	fmt.Println("请求req:",in.Request)
	out  = &pb.SearchResponse{Response:"Fuck the world !",Lenth:18}
	return
}