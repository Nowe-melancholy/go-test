package createLUsecase

import (
	"context"
	iMiddlewareRepo "go-test/domain/interface"
)

type iCreateL interface {
	CreateL(ctx context.Context, lId int) (int, error)
}

type sCreateL struct {
	repo iMiddlewareRepo.IMiddlewareRepo
}

func NewCreateLUsecase(repo iMiddlewareRepo.IMiddlewareRepo) iCreateL {
	return &sCreateL{repo}
}

func (usecase *sCreateL) CreateL(ctx context.Context, lId int) (int, error) {
	return usecase.repo.CreateL(ctx, lId)
}
