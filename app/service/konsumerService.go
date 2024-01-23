package service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"time"
	"xyzApp/app/customError"
	"xyzApp/app/helper"
	"xyzApp/app/model/dto"
	"xyzApp/app/model/entity"
	"xyzApp/app/repository"
)

type KonsumerService struct {
	Validate     *validator.Validate
	KonsumerRepo repository.IKonsumerRepository
}

// function provider
func NewKonsumerService(validate *validator.Validate, konsumerRepo repository.IKonsumerRepository) IKonsumerService {
	return &KonsumerService{
		Validate:     validate,
		KonsumerRepo: konsumerRepo,
	}
}

// method register
func (k *KonsumerService) Register(ctx context.Context, request *dto.RegisterKonsumerRequest) (*dto.RegisterKonsumerResponse, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "KonsumerService Register")
	defer span.Finish()

	// validate
	err := k.Validate.StructCtx(ctxTracing, *request)
	if err != nil {
		return nil, err
	}

	tanggalLahir, err := helper.StringToDateTime(fmt.Sprintf("%v 00:00:00", request.TanggalLahir))
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// create entity
	data := entity.Konsumer{
		Nik:          request.Nik,
		FullName:     request.FullName,
		LegalName:    request.LegalName,
		TempatLahir:  request.TempatLahir,
		TanggalLahir: tanggalLahir,
		Gaji:         request.Gaji,
		FotoKtp:      request.FotoKtp,
		FotoSelfie:   request.FotoSelfie,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// call procedure to insert
	konsumer, err := k.KonsumerRepo.Insert(ctxTracing, &data)
	if err != nil {
		return nil, err
	}

	// mapping to response
	response := dto.RegisterKonsumerResponse{
		Nik:          konsumer.Nik,
		FullName:     konsumer.FullName,
		LegalName:    konsumer.LegalName,
		TempatLahir:  konsumer.TempatLahir,
		TanggalLahir: konsumer.TanggalLahir.Format("2006-04-02"),
		Gaji:         konsumer.Gaji,
		FotoKtp:      konsumer.FotoKtp,
		FotoSelfie:   konsumer.FotoSelfie,
		CreatedAt:    helper.DateTimeToString(konsumer.CreatedAt),
		UpdatedAt:    helper.DateTimeToString(konsumer.UpdatedAt),
	}

	// success -> return response
	return &response, nil
}
