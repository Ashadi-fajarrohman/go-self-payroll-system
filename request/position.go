package request

type PositionReq struct {
	Name   string `json:"name" validate:"required"`
	Salary int    `json:"salary" validate:"required"`
}
