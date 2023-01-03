package middlewareRepo

import (
	"context"
	"fmt"
	iMiddlewareRepo "go-test/domain/interface"
	"go-test/ent"

	_ "github.com/lib/pq"
)

type middlewareRepo struct{}

func NewMiddlewareRepo() iMiddlewareRepo.IMiddlewareRepo {
	return &middlewareRepo{}
}

func (repo *middlewareRepo) CreateL(ctx context.Context, lId int) (int, error) {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		"db", "5432", "db", "root", "root"))

	if err != nil {
		return 0, err
	}

	defer client.Close()
	middleware, err := client.Debug().Middleware.
		Create().
		SetLID(lId).
		SetNillableDID(nil).
		SetSysID(1).
		Save(ctx)
	if err != nil {
		return 0, err
	}

	return middleware.ID, nil
}

func (repo *middlewareRepo) CreateD(ctx context.Context, dId int) (int, error) {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		"db", "5432", "db", "root", "root"))

	if err != nil {
		return 0, err
	}

	defer client.Close()
	middleware, err := client.Debug().Middleware.
		Create().
		SetDID(dId).
		SetNillableLID(nil).
		SetSysID(2).
		Save(ctx)
	if err != nil {
		return 0, err
	}

	return middleware.ID, nil
}
