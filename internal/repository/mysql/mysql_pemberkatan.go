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

type repoPemberkatanImpl struct {
	db  *gorm.DB
	log *logrus.Logger
	cfg *config.Configs
}

func NewRepoPemberkatan(log *logrus.Logger, cfg *config.Configs) *repoPemberkatanImpl {
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.SqlPemberkatan.User, cfg.SqlPemberkatan.Password, cfg.SqlPemberkatan.Host, cfg.SqlPemberkatan.Port, cfg.SqlPemberkatan.DbName)
	db, err := gorm.Open(mysql.Open(dbInfo))

	if err != nil {
		log.Error(err)
	}

	db.AutoMigrate(&entities.PemberkatanEntityModel{})

	return &repoPemberkatanImpl{
		db:  db,
		log: log,
		cfg: cfg,
	}
}

func (r *repoPemberkatanImpl) Create(ctx context.Context, e *entities.PemberkatanEntityModel) (*entities.PemberkatanEntityModel, error) {
	err := r.db.Create(&e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *repoPemberkatanImpl) Update(ctx context.Context, id *string, e *entities.PemberkatanEntityModel) (*entities.PemberkatanEntityModel, error) {
	err := r.db.Model(e).Where("id", id).Updates(e).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *repoPemberkatanImpl) Delete(ctx context.Context, id *string, e *entities.PemberkatanEntityModel) (*entities.PemberkatanEntityModel, error) {
	var data *entities.PemberkatanEntityModel

	err := r.db.Where("collection_id = ?", id).Delete(&data).WithContext(ctx).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func (r *repoPemberkatanImpl) FindByID(ctx context.Context, id *string) (*entities.PemberkatanEntityModel, error) {
	var data entities.PemberkatanEntityModel

	err := r.db.Where("id = ?", id).First(&data).
		WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repoPemberkatanImpl) FindAll(ctx context.Context) ([]entities.PemberkatanEntityModel, error) {
	var datas []entities.PemberkatanEntityModel

	err := r.db.Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}
	return datas, nil
}
