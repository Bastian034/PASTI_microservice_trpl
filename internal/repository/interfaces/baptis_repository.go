package interfaces

import (
	"context"
	"gereja-services/internal/entities"
)

type BaptisRepository interface {
	Create(ctx context.Context, e *entities.BaptisEntityModel) (*entities.BaptisEntityModel, error)
	Update(ctx context.Context, id *string, e *entities.BaptisEntityModel) (*entities.BaptisEntityModel, error)
	Delete(ctx context.Context, id *string) (*entities.BaptisEntityModel, error)
	FindByID(ctx context.Context, id *string) (*entities.BaptisEntityModel, error)
	FindAll(ctx context.Context) ([]entities.BaptisEntityModel, error)
}
