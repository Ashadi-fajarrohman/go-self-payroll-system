package delivery

import (
	"go-self-payroll-system/model"
	"go-self-payroll-system/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type positionDelivery struct {
	positionUsecase model.PositionUsecase
}

type PositionDelivery interface {
	Mounting(group *echo.Group)
}

func NewPositionDelivery(usecase model.PositionUsecase) PositionDelivery {
	return &positionDelivery{positionUsecase: usecase}
}

func (p positionDelivery) Mounting(group *echo.Group) {
	group.POST("", p.PostPosition)
}

func (p *positionDelivery) PostPosition(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.PositionReq

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	data, err := p.positionUsecase.PostPosition(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
