package model

import (
	"context"
	"go-self-payroll-system/request"
	"time"
)

type Position struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Salary    int       `json:"salary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PositionRepo interface {
	PostPosition(ctx context.Context, position *Position) (*Position, error)
}

type PositionUsecase interface {
	PostPosition(ctx context.Context, req *request.PositionReq) (*Position, error)
}
