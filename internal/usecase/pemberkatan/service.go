package pemberkatan

import (
	"context"
	"gereja-services/internal/repository/interfaces"
	"github.com/sirupsen/logrus"
)

type service struct {
	repo *interfaces.PemberkatanRepository
	log  *logrus.Logger
}

func NewPemberkatanServices(repo *interfaces.PemberkatanRepository, log *logrus.Logger) *service {
	return &service{
		repo: repo,
		log:  log,
	}
}

func (s *service) FindAll(ctx context.Context) (GetAllResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindByID(ctx context.Context, payload GetByIDRequest) (GetByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Create(ctx context.Context, payload CreateRequest) (CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Update(ctx context.Context, payload UpdateRequest) (UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Delete(ctx context.Context, payload DeleteRequest) (DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}
