package repository

import (
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go"
	"time"
	"xyzApp/app/customError"
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
	statement, err := k.db.PrepareContext(ctxTracing, "INSERT INTO konsumer(nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie, created_at, updated_at) VALUE (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// execute query
	result, err := statement.ExecContext(ctxTracing, entity.Nik, entity.FullName, entity.LegalName, entity.TempatLahir, entity.TanggalLahir, entity.Gaji, entity.FotoKtp, entity.FotoSelfie, time.Now(), time.Now())
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// check row affected
	if row, _ := result.RowsAffected(); row == 0 {
		return nil, customError.NewInternalSeverError("cant insert data")
	}

	// success insert
	return entity, nil
}

// method get all data konsumer
func (k *KonsumerRepository) GetAll(ctx context.Context) ([]entity.Konsumer, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "KonsumerRepository GetAll")
	defer span.Finish()

	// create prepare query
	statement, err := k.db.PrepareContext(ctxTracing, "SELECT nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie, created_at, updated_at FROM konsumer")
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// execute query
	rows, err := statement.QueryContext(ctxTracing)
	defer rows.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	var konsumer []entity.Konsumer
	for rows.Next() {
		var data entity.Konsumer
		if err := rows.Scan(&data.Nik, &data.FullName, &data.LegalName, &data.TempatLahir, &data.TanggalLahir, &data.Gaji, &data.FotoKtp, &data.FotoSelfie, &data.CreatedAt, &data.UpdatedAt); err != nil {
			return nil, customError.NewInternalSeverError(err.Error())
		}

		// append
		konsumer = append(konsumer, data)
	}

	// check if not found
	if len(konsumer) == 0 {
		return nil, customError.NewNotFoundError("record konsumer not found")
	}

	// success get all data
	return konsumer, nil
}

// method get data konsumer by nik
func (k *KonsumerRepository) GetByNik(ctx context.Context, nik string) (*entity.Konsumer, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "KonsumerRepository GetByNik")
	defer span.Finish()

	// statement
	statement, err := k.db.PrepareContext(ctxTracing, "SELECT nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie, created_at, updated_at FROM konsumer WHERE nik=?")
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// execute query
	row := statement.QueryRowContext(ctxTracing, nik)
	if row.Err() != nil {
		return nil, customError.NewNotFoundError(row.Err().Error())
	}

	// scan data
	var konsumer entity.Konsumer
	if err = row.Scan(&konsumer.Nik, &konsumer.FullName, &konsumer.LegalName, &konsumer.TempatLahir, &konsumer.TanggalLahir, &konsumer.Gaji, &konsumer.FotoKtp, &konsumer.FotoSelfie, &konsumer.CreatedAt, &konsumer.UpdatedAt); err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// success get data by nik
	return &konsumer, nil
}
