package endpoint

import (
	"context"
	"fmt"

	pbbook "com/lhh/micro/kit/protos/book"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type BookRPCEndpoints struct {
	FindBooksEndpoint endpoint.Endpoint
}

// MakeFindBooksEndpoint 使查找书籍端点
func MakeFindBooksEndpoint(instance string) endpoint.Endpoint {
	//WithInsecure 返回一个DialOption禁用此ClientConn的传输安全。
	//请注意，除非设置了 WithInsecure，否则需要传输安全性。
	conn, err := grpc.Dial(instance, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//NewClient 为单个远程端点构造一个可用的Client 。
	//将 RPC 响应类型的零值 protobuf 消息作为 grpcReply 参数传递。
	findBooksEndpoint := grpctransport.NewClient(
		conn, "book.Book", "FindBooksByUserID",
		encodeGRPCFindBooksRequest,
		decodeGRPCFindBooksResponse,
		pbbook.BooksResponse{},
	).Endpoint()
	return findBooksEndpoint
}

func encodeGRPCFindBooksRequest(_ context.Context, r interface{}) (interface{}, error) {
	userID := r.(uint64)
	return &pbbook.BooksByUserIDRequest{
		UserID: userID,
	}, nil
}

func decodeGRPCFindBooksResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pbbook.BooksResponse)
	return resp.Books, nil
}
