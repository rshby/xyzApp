package repository

import (
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go"
	"xyzApp/app/model/entity"
)

type KonsumerRepository struct {
	db *sql.DB
}

// function provider
func NewKonsumerRepository(db *sql.DB) IKonsumerRepository {
	return &KonsumerRepository{db: db}
}

// method insert
func (k *KonsumerRepository) Insert(ctx context.Context, entity *entity.Konsumer) (*entity.Konsumer, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "KonsumerRepo Insert")
	defer span.Finish()

	// create statement query
	statement, err := k.db.PrepareContext(ctxTracing, "INSERT INTO konsumer (nik, )") // TODO : kerjakan query insert
}

// method get all data konsumer
func (k *KonsumerRepository) GetAll(ctx context.Context) ([]entity.Konsumer, error) {
	//TODO implement me
	panic("implement me")
}

// method get data konsumer by nik
func (k *KonsumerRepository) GetByNik(ctx context.Context, nik string) (*entity.Konsumer, error) {
	//TODO implement me
	panic("implement me")
}
