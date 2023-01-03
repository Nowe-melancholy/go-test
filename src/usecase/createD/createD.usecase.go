package createDUsecase

import (
	"context"
	iMiddlewareRepo "go-test/domain/interface"
)

type iCreateD interface {
	CreateD(ctx context.Context, dId int) (int, error)
}

type sCreateD struct {
	repo iMiddlewareRepo.IMiddlewareRepo
}

func NewCreateDUsecase(repo iMiddlewareRepo.IMiddlewareRepo) iCreateD {
	return &sCreateD{repo}
}

func (usecase *sCreateD) CreateD(ctx context.Context, dId int) (int, error) {
	return usecase.repo.CreateD(ctx, dId)
}
