package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"sync"
	"xyzApp/app/customError"
	"xyzApp/app/helper"
	"xyzApp/app/model/dto"
	"xyzApp/app/model/entity"
	"xyzApp/app/repository"
)

type TransactionService struct {
	Validate        *validator.Validate
	KonsumerRepo    repository.IKonsumerRepository
	TenorRepo       repository.ITenorRepository
	TransactionRepo repository.ITransactionRepository
}

// function provider
func NewTransactionService(validate *validator.Validate, konsumerRepo repository.IKonsumerRepository,
	tenorRepo repository.ITenorRepository, transactionRepo repository.ITransactionRepository) ITransactionService {
	return &TransactionService{
		Validate:        validate,
		KonsumerRepo:    konsumerRepo,
		TenorRepo:       tenorRepo,
		TransactionRepo: transactionRepo,
	}
}

// method Buy
func (t *TransactionService) Buy(ctx context.Context, request *dto.BuyRequest) (*dto.BuyResponse, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "TransactionService Buy")
	defer span.Finish()

	// validate request
	if err := t.Validate.StructCtx(ctxTracing, *request); err != nil {
		return nil, err
	}

	// define waitgroup
	wg := &sync.WaitGroup{}

	errKonsumerChan := make(chan error, 1)
	tenorChan := make(chan entity.Tenor, 1)
	errTenorChan := make(chan error, 1)
	defer func() {
		close(errTenorChan)
		close(tenorChan)
		close(errKonsumerChan)
	}()

	// get data konsumer by nik
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup, nik string, errChan chan error) {
		defer wg.Done()
		_, err := t.KonsumerRepo.GetByNik(ctx, nik)
		if err != nil {
			errChan <- err
			return
		}

		// tidak ada error
		errChan <- nil
	}(ctxTracing, wg, request.Nik, errKonsumerChan)

	// get data tenor apakah ada
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup, nik string, date string, tenorChan chan entity.Tenor, errChan chan error) {
		defer wg.Done()

		// ubah date (string) menjadi time.Time
		dateTrx, _ := helper.StringToDateTime(date)
		data, err := t.TenorRepo.GetByNikAndDate(ctx, nik, dateTrx)
		if err != nil {
			tenorChan <- entity.Tenor{}
			errChan <- customError.NewNotFoundError("no record limit was found by the specified date_transaction")
			return
		}

		// success get data
		tenorChan <- *data
		errChan <- nil
	}(ctxTracing, wg, request.Nik, request.DateTransaction, tenorChan, errTenorChan)

	// wait
	wg.Wait()

	// cek apakah nik tersebut ada
	if err := <-errKonsumerChan; err != nil {
		return nil, customError.NewNotFoundError("record konsumer not found")
	}

	tenor := <-tenorChan
	err := <-errTenorChan
	if err != nil {
		return nil, err
	}

	// hitung otr/jumlah_cicilan
	var basePrice float64 = request.Otr / float64(request.JumlahCicilan)
	var bunga float64 = (request.Bunga / 100) * basePrice
	totalDebet := basePrice + bunga + request.AdminFee

	// cek apakah limit cukup
	if totalDebet > tenor.Tenor {
		return nil, customError.NewInternalSeverError("limit not enough")
	}

	// hitung sisa limit
	sisaLimit := tenor.Tenor - totalDebet

	// update sisa limit di tenor
	if err := t.TenorRepo.UpdateLimit(ctxTracing, sisaLimit, request.Nik, tenor.Bulan); err != nil {
		return nil, err
	}

	// insert ke tabel transaction
	dateTrx, _ := helper.StringToDateTime(request.DateTransaction)
	transaction := entity.Transaction{
		ReffNumber:      helper.GenerateReffNumber(request.DateTransaction, request.Nik),
		Nik:             request.Nik,
		DateTransaction: dateTrx,
		Otr:             request.Otr,
		AdminFee:        request.AdminFee,
		JumlahCicilan:   request.JumlahCicilan,
		JumlahBunga:     request.Bunga,
		Aset:            request.Aset,
		TotalDebet:      totalDebet,
	}
	_, err = t.TransactionRepo.Insert(ctxTracing, &transaction)
	if err != nil {
		return nil, err
	}

	// create response
	response := dto.BuyResponse{
		ReffNumber:    helper.GenerateReffNumber(request.DateTransaction, request.Nik),
		Nik:           request.Nik,
		Otr:           request.Otr,
		AdminFee:      request.AdminFee,
		JumlahCicilan: request.JumlahCicilan,
		Bunga:         request.Bunga,
		Aset:          request.Aset,
		TotalDebet:    totalDebet,
		SisaLimit:     sisaLimit,
	}

	return &response, nil
}
