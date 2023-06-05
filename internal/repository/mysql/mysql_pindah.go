package mysql

import (
	"context"
	"fmt"
	"gereja-services/config"
	"gereja-services/internal/entities"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type repoPindahImpl struct {
	db  *gorm.DB
	log *logrus.Logger
	cfg *config.Configs
}

func NewRepoPindah(log *logrus.Logger, cfg *config.Configs) *repoPindahImpl {
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.SqlPindah.User, cfg.SqlPindah.Password, cfg.SqlPindah.Host, cfg.SqlPindah.Port, cfg.SqlPindah.DbName)
	db, err := gorm.Open(mysql.Open(dbInfo))

	if err != nil {
		log.Error(err)
	}

	db.AutoMigrate(&entities.PindahEntityModel{})

	return &repoPindahImpl{
		db:  db,
		log: log,
		cfg: cfg,
	}
}

func (r *repoPindahImpl) Create(ctx context.Context, e *entities.PindahEntityModel) (*entities.PindahEntityModel, error) {
	err := r.db.Create(&e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *repoPindahImpl) Update(ctx context.Context, id *string, e *entities.PindahEntityModel) (*entities.PindahEntityModel, error) {
	err := r.db.Model(e).Where("id", id).Updates(e).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *repoPindahImpl) Delete(ctx context.Context, id *string, e *entities.PindahEntityModel) (*entities.PindahEntityModel, error) {
	var data *entities.PindahEntityModel

	err := r.db.Where("collection_id = ?", id).Delete(&data).WithContext(ctx).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func (r *repoPindahImpl) FindByID(ctx context.Context, id *string) (*entities.PindahEntityModel, error) {
	var data entities.PindahEntityModel

	err := r.db.Where("id = ?", id).First(&data).
		WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repoPindahImpl) FindAll(ctx context.Context) ([]entities.PindahEntityModel, error) {
	var datas []entities.PindahEntityModel

	err := r.db.Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}
	return datas, nil
}
