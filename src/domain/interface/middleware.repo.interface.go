package iMiddlewareRepo

import "context"

type IMiddlewareRepo interface {
	CreateL(ctx context.Context, lId int) (int, error)

	CreateD(ctx context.Context, dId int) (int, error)
}
