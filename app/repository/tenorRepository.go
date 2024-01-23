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

// method to get data tenor by nik and date_transaction
func (t *TenorRepository) GetByNikAndDate(ctx context.Context, nik string, dateTransaction time.Time) (*entity.Tenor, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "TenorRepository GetByNikAndDate")
	defer span.Finish()

	// prepare query
	statement, err := t.db.PrepareContext(ctxTracing, "SELECT t.id, t.nik, t.bulan, t.start_date, t.end_date, t.tenor, t.created_at, t.updated_at FROM tenor t WHERE t.nik=? AND (? BETWEEN t.start_date AND t.end_date) LIMIT 1")
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// query
	row := statement.QueryRowContext(ctxTracing, nik, dateTransaction)
	if row.Err() != nil {
		return nil, customError.NewInternalSeverError(row.Err().Error())
	}

	var tenor entity.Tenor
	if err := row.Scan(&tenor.Id, &tenor.Nik, &tenor.Bulan, &tenor.StartDate, &tenor.EndDate, &tenor.Tenor, &tenor.CreatedAt, &tenor.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, customError.NewNotFoundError("record tenor not found")
		}

		return nil, customError.NewInternalSeverError(err.Error())
	}

	// success get data
	return &tenor, nil
}

// update limit by nik dan bulan
func (t *TenorRepository) UpdateLimit(ctx context.Context, limit float64, nik, bulan string) error {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "TenorRepository UpdateLimit")
	defer span.Finish()

	// prepare query
	statement, err := t.db.PrepareContext(ctxTracing, "UPDATE tenor SET tenor.tenor = ? WHERE nik=? AND bulan=?")
	defer statement.Close()
	if err != nil {
		return customError.NewInternalSeverError(err.Error())
	}

	// execute query
	result, err := statement.ExecContext(ctxTracing, limit, nik, bulan)
	if err != nil {
		return customError.NewInternalSeverError(err.Error())
	}

	if row, _ := result.RowsAffected(); row == 0 {
		return customError.NewInternalSeverError("cant update limit")
	}

	// success update limit
	return nil
}
