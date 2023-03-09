package endpoint

import (
	"com/lhh/micro/kit/library-book-grpc-service/service"
	"context"

	pbbook "com/lhh/micro/kit/protos/book"
	"github.com/go-kit/kit/endpoint"
)

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