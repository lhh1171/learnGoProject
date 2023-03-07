package endpoint

import (
	"context"
	"micro/kit/library-book-grpc-service/service"

	"github.com/go-kit/kit/endpoint"
	pbbook "micro/kit/protos/book"
)

// BookEndpoints 每一个请求对应一个endpoint。
type BookEndpoints struct {
	FindBooksByUserIDEndpoint endpoint.Endpoint
}

func NewFindBooksByUserIDEndpoint(bookService service.BookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pbbook.BooksByUserIDRequest)
		res, err := bookService.FindBooksByUserID(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
