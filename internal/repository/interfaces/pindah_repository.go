package interfaces

import (
	"context"
	"gereja-services/internal/entities"
)

type PindahRepository interface {
	Create(ctx context.Context, e *entities.PindahEntityModel) (*entities.PindahEntityModel, error)
	Update(ctx context.Context, id *string, e *entities.PindahEntityModel) (*entities.PindahEntityModel, error)
	Delete(ctx context.Context, id *string, e *entities.PindahEntityModel) (*entities.PindahEntityModel, error)
	FindByID(ctx context.Context, id *string) (*entities.PindahEntityModel, error)
	FindAll(ctx context.Context) ([]entities.PindahEntityModel, error)
}
