package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/go-sisko/helper"
	"github.com/renaldid/go-sisko/model/domain"
)

type SiswaRepoImpl struct {
}

func NewSiswaRepo() SiswaRepo {
	return &SiswaRepoImpl{}
}

func (s SiswaRepoImpl) Save(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa {
	SQL := "insert into siswa(nama, alamat, tanggal_lahir, tempat_lahir, jenis_kelamin, agama, golongan_darah, nomor_telepon) value (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, siswa.Nama, siswa.Alamat, siswa.TanggalLahir, siswa.TempatLahir, siswa.JenisKelamin, siswa.Agama, siswa.GolonganDarah, siswa.NomorTelepon)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	siswa.Id = int(id)
	return siswa
}

func (s SiswaRepoImpl) Update(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa {
	SQL := "update siswa set nama =?, alamat=?, tanggal_lahir=?, tempat_lahir=?, jenis_kelamin=?, agama=?, golongan_darah=?, nomor_telepon=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, siswa.Nama, siswa.Alamat, siswa.TanggalLahir, siswa.TempatLahir, siswa.JenisKelamin, siswa.Agama, siswa.GolonganDarah, siswa.NomorTelepon, siswa.Id)
	helper.PanicIfError(err)
	return siswa
}

func (s SiswaRepoImpl) Delete(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) {
	SQL := "delete from siswa where id = ?"
	_, err := tx.ExecContext(ctx, SQL, siswa.Id)
	helper.PanicIfError(err)
}

func (s SiswaRepoImpl) FindById(ctx context.Context, tx *sql.Tx, siswaId int) (domain.Siswa, error) {
	SQL := "select id, nama, alamat, tanggal_lahir, tempat_lahir, jenis_kelamin, agama, golongan_darah, nomor_telepon from siswa where id =?"
	rows, err := tx.QueryContext(ctx, SQL, siswaId)
	helper.PanicIfError(err)
	defer rows.Close()

	siswa := domain.Siswa{}
	if rows.Next() {
		err := rows.Scan(&siswa.Id, &siswa.Nama, &siswa.Alamat, &siswa.TanggalLahir, &siswa.TempatLahir, &siswa.JenisKelamin, &siswa.Agama, &siswa.GolonganDarah, &siswa.NomorTelepon)
		helper.PanicIfError(err)
		return siswa, nil
	} else {
		return siswa, errors.New("siswa is not found")
	}
}

func (s SiswaRepoImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Siswa {
	SQL := "select id, nama, alamat, tanggal_lahir, tempat_lahir, jenis_kelamin, agama, golongan_darah, nomor_telepon from siswa"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var siswas []domain.Siswa
	for rows.Next() {
		siswa := domain.Siswa{}
		err := rows.Scan(&siswa.Id, &siswa.Nama, &siswa.Alamat, &siswa.TanggalLahir, &siswa.TempatLahir, &siswa.JenisKelamin, &siswa.Agama, &siswa.GolonganDarah, &siswa.NomorTelepon)
		helper.PanicIfError(err)
		siswas = append(siswas, siswa)
	}
	return siswas
}
