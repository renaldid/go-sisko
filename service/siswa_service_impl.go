package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/renaldid/go-sisko/exception"
	"github.com/renaldid/go-sisko/helper"
	"github.com/renaldid/go-sisko/model/domain"
	"github.com/renaldid/go-sisko/model/web"
	"github.com/renaldid/go-sisko/repository"
)

type SiswaServiceImpl struct {
	SiswaRepo repository.SiswaRepo
	DB        *sql.DB
	Validate  *validator.Validate
}

func NewSiswaService(siswaRepo repository.SiswaRepo, DB *sql.DB, Validate *validator.Validate) SiswaService {
	return &SiswaServiceImpl{
		SiswaRepo: siswaRepo,
		DB:        DB,
		Validate:  Validate,
	}
}

func (service *SiswaServiceImpl) Create(ctx context.Context, request web.SiswaCreateRequest) web.SiswaResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	siswa := domain.Siswa{
		Nama:          request.Nama,
		Alamat:        request.Alamat,
		TanggalLahir:  request.TanggalLahir,
		TempatLahir:   request.TempatLahir,
		JenisKelamin:  request.JenisKelamin,
		Agama:         request.Agama,
		GolonganDarah: request.GolonganDarah,
		NomorTelepon:  request.NomorTelepon,
	}
	siswa = service.SiswaRepo.Save(ctx, tx, siswa)
	return helper.ToSiswaResponse(siswa)

}

func (service *SiswaServiceImpl) Update(ctx context.Context, request web.SiswaUpdateRequest) web.SiswaResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	siswa, err := service.SiswaRepo.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	siswa.Nama = request.Nama
	siswa.Alamat = request.Alamat
	siswa.TanggalLahir = request.TanggalLahir
	siswa.TempatLahir = request.TempatLahir
	siswa.JenisKelamin = request.JenisKelamin
	siswa.Agama = request.Agama
	siswa.GolonganDarah = request.GolonganDarah
	siswa.NomorTelepon = request.NomorTelepon

	siswa = service.SiswaRepo.Update(ctx, tx, siswa)
	return helper.ToSiswaResponse(siswa)
}

func (service *SiswaServiceImpl) Delete(ctx context.Context, siswaId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	siswa, err := service.SiswaRepo.FindById(ctx, tx, siswaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.SiswaRepo.Delete(ctx, tx, siswa)
}

func (service *SiswaServiceImpl) FindById(ctx context.Context, siswaId int) web.SiswaResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	siswa, err := service.SiswaRepo.FindById(ctx, tx, siswaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToSiswaResponse(siswa)
}

func (service *SiswaServiceImpl) FindByAll(ctx context.Context) []web.SiswaResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	siswas := service.SiswaRepo.FindByAll(ctx, tx)
	return helper.ToSiswaResponses(siswas)
}
