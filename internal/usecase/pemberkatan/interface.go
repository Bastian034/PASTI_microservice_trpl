package pemberkatan

import (
	"context"
	"gereja-services/internal/entities"
)

type PemberkatanService interface {
	FindAll(ctx context.Context) ([]entities.PemberkatanEntityModel, error)
	FindByID(ctx context.Context, payload *GetByIDRequest) (*GetByIDResponse, error)
	Create(ctx context.Context, payload *CreateRequest) (*CreateResponse, error)
	Update(ctx context.Context, payload *UpdateRequest) (*UpdateResponse, error)
	Delete(ctx context.Context, payload *DeleteRequest) (*DeleteResponse, error)
}
