package baptis

import (
	"gereja-services/internal/repository/interfaces"
	"gereja-services/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service *service
}

func NewHandler(r interfaces.BaptisRepository, l *logrus.Logger) *handler {
	service := NewBaptisServices(r, l)
	return &handler{
		service: service,
	}
}

// Create godoc
// @Summary Create Baptis
// @Description Create Baptis
// @Tags Baptis
// @Accept  json
// @Produce  json
// @Param request body CreateRequest true "request body"
// @Success 200 {object} CreateResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /baptis [post]
func (h *handler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var payload CreateRequest

	if err := c.Bind(payload); err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Create(ctx, &payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}
	return response.SuccessResponse(result).Send(c)
}

// Update godoc
// @Summary Update Baptis
// @Description Update Baptis
// @Tags Baptis
// @Accept  json
// @Produce  json
// @Param request body UpdateRequest true "request body"
// @Success 200 {object} UpdateResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /baptis [patch]
func (h *handler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	var payload UpdateRequest

	if err := c.Bind(payload); err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Update(ctx, &payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}
	return response.SuccessResponse(result).Send(c)
}

// Delete godoc
// @Summary Delete Baptis
// @Description Delete Baptis
// @Tags Baptis
// @Accept  json
// @Produce  json
// @Param id path string true "id path"
// @Success 200 {object} DeleteResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /baptis/{id} [delete]
func (h *handler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	var payload DeleteRequest

	if err := c.Bind(payload); err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Delete(ctx, &payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}
	return response.SuccessResponse(result).Send(c)
}

// GetByID
// @Summary Get Baptis By ID
// @Description Get Baptis By ID
// @Tags Baptis
// @Accept json
// @Produce json
// @Param id path string true "id path"
// @Success 200 {object} GetByIDResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /baptis/{id} [get]
func (h *handler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	var payload GetByIDRequest

	if err := c.Bind(payload); err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.FindByID(ctx, &payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}
	return response.SuccessResponse(result).Send(c)
}

// GetAll
// @Summary Get All Baptis
// @Description Get ALl Baptis
// @Tags Baptis
// @Accept json
// @Produce json
// @Success 200 {object} GetAllResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /baptis [get]
func (h *handler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := h.service.FindAll(ctx)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}
	return response.SuccessResponse(result).Send(c)
}
