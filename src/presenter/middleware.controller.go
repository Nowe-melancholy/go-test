package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	middlewareRepo "go-test/repos"
	createDUsecase "go-test/usecase/createD"
)

type Output struct {
	Id int
}

type Input struct {
	DId int `json:"did"`
}

func CreateD(context echo.Context) error {
	input := new(Input)
	if err := context.Bind(input); err != nil {
		panic(err)
	}
	log.Println(input.DId)

	repo := middlewareRepo.NewMiddlewareRepo()
	usecase := createDUsecase.NewCreateDUsecase(repo)

	result, err := usecase.CreateD(context.Request().Context(), input.DId)
	if err != nil {
		return err
	}

	output := Output{Id: result}

	return context.JSON(http.StatusOK, output)
}
