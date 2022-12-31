package iMiddlewareRepo

import "context"

type IMiddlewareRepo interface {
	CreateL(ctx context.Context, lId int) (int, error)
}
