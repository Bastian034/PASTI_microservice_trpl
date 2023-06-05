package interfaces

import (
	"context"
	"gereja-services/internal/entities"
)

type PemberkatanRepository interface {
	Create(ctx context.Context, e *entities.PemberkatanEntityModel) (*entities.PemberkatanEntityModel, error)
	Update(ctx context.Context, id *string, e *entities.PemberkatanEntityModel) (*entities.PemberkatanEntityModel, error)
	Delete(ctx context.Context, id *string, e *entities.PemberkatanEntityModel) (*entities.PemberkatanEntityModel, error)
	FindByID(ctx context.Context, id *string) (*entities.PemberkatanEntityModel, error)
	FindAll(ctx context.Context) ([]entities.PemberkatanEntityModel, error)
}
