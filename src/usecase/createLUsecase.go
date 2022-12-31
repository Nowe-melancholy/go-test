package createLUsecase

import (
	"context"
	iMiddlewareRepo "go-test/domain/interface"
)

type CreateLUsecase interface {
	CreateL(ctx context.Context, lId int) (int, error)
}

type createLUsecase struct {
	repo iMiddlewareRepo.IMiddlewareRepo
}

func NewCreateLUsecase(repo iMiddlewareRepo.IMiddlewareRepo) CreateLUsecase {
	return &createLUsecase{repo}
}

func (usecase *createLUsecase) CreateL(ctx context.Context, lId int) (int, error) {
	return usecase.repo.CreateL(ctx, lId)
}
