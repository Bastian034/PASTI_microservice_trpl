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

type repoBaptisImpl struct {
	db  *gorm.DB
	log *logrus.Logger
	cfg *config.Configs
}

func NewRepoBaptis(log *logrus.Logger, cfg *config.Configs) *repoBaptisImpl {
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.SqlBaptis.User, cfg.SqlBaptis.Password, cfg.SqlBaptis.Host, cfg.SqlBaptis.Port, cfg.SqlBaptis.DbName)
	db, err := gorm.Open(mysql.Open(dbInfo))

	if err != nil {
		log.Error(err)
	}

	db.AutoMigrate(&entities.BaptisEntityModel{})

	return &repoBaptisImpl{
		db:  db,
		log: log,
		cfg: cfg,
	}
}

func (r *repoBaptisImpl) Create(ctx context.Context, e *entities.BaptisEntityModel) (*entities.BaptisEntityModel, error) {
	err := r.db.Create(&e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *repoBaptisImpl) Update(ctx context.Context, id *string, e *entities.BaptisEntityModel) (*entities.BaptisEntityModel, error) {
	err := r.db.Model(e).Where("id", id).Updates(e).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *repoBaptisImpl) Delete(ctx context.Context, id *string) (*entities.BaptisEntityModel, error) {
	var data *entities.BaptisEntityModel

	err := r.db.Where("collection_id = ?", id).Delete(&data).WithContext(ctx).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func (r *repoBaptisImpl) FindByID(ctx context.Context, id *string) (*entities.BaptisEntityModel, error) {
	var data entities.BaptisEntityModel

	err := r.db.Where("id = ?", id).First(&data).
		WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repoBaptisImpl) FindAll(ctx context.Context) ([]entities.BaptisEntityModel, error) {
	var datas []entities.BaptisEntityModel

	err := r.db.Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}
	return datas, nil
}
