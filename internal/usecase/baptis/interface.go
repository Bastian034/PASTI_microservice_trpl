package baptis

import (
	"context"
	"gereja-services/internal/entities"
)

type BaptisService interface {
	FindAll(ctx context.Context) ([]entities.BaptisEntityModel, error)
	FindByID(ctx context.Context, payload *GetByIDRequest) (*GetByIDResponse, error)
	Create(ctx context.Context, payload *CreateRequest) (*CreateResponse, error)
	Update(ctx context.Context, payload *UpdateRequest) (*UpdateResponse, error)
	Delete(ctx context.Context, payload *DeleteRequest) (*DeleteResponse, error)
}
