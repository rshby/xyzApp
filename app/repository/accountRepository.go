package repository

import (
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go"
	"xyzApp/app/customError"
	"xyzApp/app/model/entity"
)

type AccountRepository struct {
	db *sql.DB
}

// function provider
func NewAccountRepository(db *sql.DB) IAccountRepository {
	return &AccountRepository{db: db}
}

// method insert
func (a *AccountRepository) Insert(ctx context.Context, entity *entity.Account) (*entity.Account, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "AccountRepository Insert")
	defer span.Finish()

	// preapre query
	statement, err := a.db.PrepareContext(ctxTracing, "INSERT INTO accounts(nik, email, password) VALUES (?, ?, ?)")
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// execute
	result, err := statement.ExecContext(ctxTracing, entity.Nik, entity.Email, entity.Password)
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// cek row affected
	if row, _ := result.RowsAffected(); row == 0 {
		return nil, customError.NewInternalSeverError("cant insert data account to database")
	}

	// success
	return entity, nil
}

// method get by id
func (a *AccountRepository) GetByEmail(ctx context.Context, email string) (*entity.Account, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "AccountRepository GetByEmail")
	defer span.Finish()

	// statement
	statement, err := a.db.PrepareContext(ctxTracing, "SELECT nik, email, password FROM accounts WHERE email=?")
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// query
	row := statement.QueryRowContext(ctxTracing, email)
	if row.Err() != nil {
		return nil, customError.NewInternalSeverError(row.Err().Error())
	}

	var account entity.Account
	if err := row.Scan(&account.Nik, &account.Email, &account.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, customError.NewNotFoundError("record account not found")
		}

		return nil, customError.NewInternalSeverError(err.Error())
	}

	// success get data account by email
	return &account, nil
}
