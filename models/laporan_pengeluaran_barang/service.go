package laporanpengeluaranbarang

import (
	"laporan_stock_barang/helper"
	"laporan_stock_barang/models"
)

type service struct {
	repository Repository
}

type Service interface {
	InputHeaderPengeluaran(input models.RequestForm) (models.ResultIdLaporan, error)
	InputDetailPengeluaran(input models.RequestDetailProductTrx) error
}

func NewService(repo Repository) *service {
	return &service{repository: repo}
}

func (s *service) InputHeaderPengeluaran(input models.RequestForm) (models.ResultIdLaporan, error) {
	var err error
	var result models.ResultIdLaporan
	result, err = s.repository.InputHeaderPengeluaran(input)
	helper.PanicIfError(err)
	return result, nil
}

func (s *service) InputDetailPengeluaran(input models.RequestDetailProductTrx) error {
	err := s.repository.InputDetailPengeluaran(input)
	helper.PanicIfError(err)
	return nil
}
