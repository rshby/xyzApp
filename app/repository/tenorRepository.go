package repository

import (
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go"
	"time"
	"xyzApp/app/customError"
	"xyzApp/app/model/entity"
)

type TenorRepository struct {
	db *sql.DB
}

// func provider
func NewTenorRepository(db *sql.DB) ITenorRepository {
	return &TenorRepository{db}
}

// method insert data tenor
func (t *TenorRepository) Insert(ctx context.Context, input *entity.Tenor) (*entity.Tenor, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "TenorRepository Insert")
	defer span.Finish()

	// prepare query
	statement, err := t.db.PrepareContext(ctxTracing, "INSERT INTO tenor (nik, bulan, start_date, end_date, tenor, created_at, updated_at) VALUE (?, ?, ?, ?, ?, ?, ?)")
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// execute
	result, err := statement.ExecContext(ctxTracing, input.Nik, input.Bulan, input.StartDate, input.EndDate, input.Tenor, time.Now(), time.Now())
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// check row affeected
	if row, _ := result.RowsAffected(); row == 0 {
		return nil, customError.NewInternalSeverError("cant insert data tenor")
	}

	// success insert
	return input, nil
}
