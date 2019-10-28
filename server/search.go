package srv

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	pb "grpc_demo/proto"
)

type SearchService struct {
}

func (this *SearchService) Search(ctx context.Context, in *pb.SearchRequest) (out *pb.SearchResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "search")
	defer span.Finish()
	fmt.Println("请求req:", in.Request)
	out = &pb.SearchResponse{Response: "Fuck the world !", Lenth: 18}
	return
}

func (this *SearchService) Echo(ctx context.Context, in *pb.StringMessage) (out *pb.StringMessage, err error) {
	fmt.Println("echo:", in.Words)
	out = &pb.StringMessage{Words: in.Words}
	return
}
