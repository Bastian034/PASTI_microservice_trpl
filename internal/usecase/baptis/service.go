package baptis

import (
	"context"
	"errors"
	"gereja-services/internal/entities"
	"gereja-services/internal/repository/interfaces"
	"gereja-services/pkg/response"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type service struct {
	repo interfaces.BaptisRepository
	log  *logrus.Logger
}

func NewBaptisServices(repo interfaces.BaptisRepository, log *logrus.Logger) *service {
	return &service{
		repo: repo,
		log:  log,
	}
}

func (s *service) FindAll(ctx context.Context) ([]entities.BaptisEntityModel, error) {
	datas := make([]entities.BaptisEntityModel, 0)

	datas, err := s.repo.FindAll(ctx)
	if err != nil {
		return datas, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	return datas, nil
}

func (s *service) FindByID(ctx context.Context, payload *GetByIDRequest) (*GetByIDResponse, error) {
	var result *GetByIDResponse

	data, err := s.repo.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return result, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result = &GetByIDResponse{
		Datas: *data,
	}

	return result, nil
}

func (s *service) Create(ctx context.Context, payload *CreateRequest) (*CreateResponse, error) {
	result := &entities.BaptisEntityModel{
		Entity:       entities.Entity{ID: uuid.NewString()},
		BaptisEntity: payload.BaptisEntity,
	}
	data, err := s.repo.Create(ctx, result)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}
	return &CreateResponse{data: *data}, nil
}

func (s *service) Update(ctx context.Context, payload *UpdateRequest) (*UpdateResponse, error) {
	data, err := s.repo.Update(ctx, &payload.ID, &payload.BaptisEntityModel)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}
	return &UpdateResponse{
		data: *data,
	}, nil

}

func (s *service) Delete(ctx context.Context, payload *DeleteRequest) (*DeleteResponse, error) {
	_, err := s.repo.Delete(ctx, &payload.ID)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}

	return &DeleteResponse{
		ID: &payload.ID,
	}, err
}
