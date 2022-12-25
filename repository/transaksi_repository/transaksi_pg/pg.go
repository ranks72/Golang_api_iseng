package transaksi_pg

import (
	"MALIKI-KARIM/repository/transaksi_repository"

	"gorm.io/gorm"
)

type transaksiPG struct {
	db *gorm.DB
}

func NewTransaksiPG(db *gorm.DB) transaksi_repository.TransaksiRepository {
	return &transaksiPG{
		db: db,
	}
}
