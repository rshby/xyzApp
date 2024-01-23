package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"xyzApp/app/customError"
	"xyzApp/app/helper"
	"xyzApp/app/model/dto"
	"xyzApp/app/model/entity"
	"xyzApp/app/repository"
)

type TenorService struct {
	Validate  *validator.Validate
	TenorRepo repository.ITenorRepository
}

// function provider
func NewTenorService(validate *validator.Validate, tenorRepo repository.ITenorRepository) ITenorService {
	return &TenorService{
		Validate:  validate,
		TenorRepo: tenorRepo,
	}
}

// method insert limit
func (t *TenorService) InsertLimit(ctx context.Context, request *dto.InsertLimitRequest) (*dto.InsertLimitResponse, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "TenorService InsertLimit")
	defer span.Finish()

	// 1. validasi struct
	if err := t.Validate.StructCtx(ctxTracing, *request); err != nil {
		// send error validation
		return nil, err
	}

	if request.Bulan > 12 {
		return nil, customError.NewBadRequestError("bulan tidak boleh lebih dari 12")
	}

	// create entity tenor
	firstDate, lastDate := helper.GetFirstAndLastDate(request.Bulan, request.Tahun)
	input := entity.Tenor{
		Nik:       request.Nik,
		Bulan:     helper.MonthToText(request.Bulan),
		StartDate: firstDate,
		EndDate:   lastDate,
		Tenor:     request.Limit,
	}

	// call procedure in repository to insert
	tenor, err := t.TenorRepo.Insert(ctxTracing, &input)
	if err != nil {
		return nil, err
	}

	// mapping to response
	response := dto.InsertLimitResponse{
		Nik:       tenor.Nik,
		Bulan:     tenor.Bulan,
		StartDate: tenor.StartDate.Format("2006-01-02"),
		EndDate:   tenor.EndDate.Format("2006-01-02"),
		Limit:     tenor.Tenor,
	}

	return &response, nil
}
