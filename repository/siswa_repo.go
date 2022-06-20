package repository

import (
	"context"
	"database/sql"
	"github.com/renaldid/go-sisko/model/domain"
)

type SiswaRepo interface {
	Save(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa
	Update(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa
	Delete(ctx context.Context, tx *sql.Tx, siswa domain.Siswa)
	FindById(ctx context.Context, tx *sql.Tx, siswaId int) (domain.Siswa, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.Siswa
}
