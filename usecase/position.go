package usecase

import (
	"context"
	"go-self-payroll-system/model"
	"go-self-payroll-system/request"
)

type positionUsecase struct {
	positionRepo model.PositionRepo
}

func NewPositionUsecase(repo model.PositionRepo) model.PositionUsecase {
	return &positionUsecase{positionRepo: repo}
}

func (p *positionUsecase) PostPosition(ctx context.Context, req *request.PositionReq) (*model.Position, error) {
	position := &model.Position{
		Name: req.Name,
		Salary: req.Salary,
	}

	data, err := p.positionRepo.PostPosition(ctx, position)
	if err != nil {
		return nil, err
	}

	return data, nil

}