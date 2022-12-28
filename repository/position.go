package repository

import (
	"context"
	"go-self-payroll-system/config"
	"go-self-payroll-system/model"
)

type positionRepo struct {
	Cfg config.Config
}

func NewPositionRepo(cfg config.Config) model.PositionRepo {
	return &positionRepo{Cfg: cfg}
}

func (p *positionRepo) PostPosition(ctx context.Context, position *model.Position) (*model.Position, error) {
	if err := p.Cfg.Database().WithContext(ctx).Create(&position).Error; err != nil {
		return nil, err
	}

	return position, nil
}
