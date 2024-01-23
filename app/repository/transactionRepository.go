package repository

import (
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go"
	"xyzApp/app/customError"
	"xyzApp/app/model/entity"
)

type TransactionRepository struct {
	db *sql.DB
}

// function provider
func NewTransactionRepository(db *sql.DB) ITransactionRepository {
	return &TransactionRepository{db: db}
}

func (t *TransactionRepository) Insert(ctx context.Context, input *entity.Transaction) (*entity.Transaction, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "TransactionRepositry Insert")
	defer span.Finish()

	// create query statement
	statement, err := t.db.PrepareContext(ctxTracing, "INSERT INTO transactions (reff_number, nik, date_transaction, otr, admin_fee, jumlah_cicilan, jumlah_bunga, aset, total_debet) VALUES (?, ?, ?, ?, ?, ?, ?, ? ,?)")
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// execute query
	result, err := statement.ExecContext(ctxTracing, input.ReffNumber, input.Nik, input.DateTransaction, input.Otr, input.AdminFee, input.JumlahCicilan, input.JumlahBunga, input.Aset, input.TotalDebet)
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}
	// cek row affected
	if row, _ := result.RowsAffected(); row == 0 {
		return nil, customError.NewInternalSeverError("cant insert data transaction")
	}

	// success insert
	return input, nil
}
