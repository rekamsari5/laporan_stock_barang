package laporanpenerimaan

import (
	"laporan_stock_barang/helper"
	"laporan_stock_barang/models"
)

type service struct {
	repository Repository
}

type Service interface {
	InputHeaderPenerimaan(input models.RequestForm) (models.ResultIdLaporan, error)
	InputDetailPenerimaan(input models.RequestDetailProductTrx) error
}

func NewService(repo Repository) *service {
	return &service{repository: repo}
}

func (s *service) InputHeaderPenerimaan(input models.RequestForm) (models.ResultIdLaporan, error) {
	var err error
	var result models.ResultIdLaporan
	result, err = s.repository.InputHeaderPenerimaan(input)
	helper.PanicIfError(err)
	return result, nil
}

func (s *service) InputDetailPenerimaan(input models.RequestDetailProductTrx) error {
	err := s.repository.InputDetailPenerimaan(input)
	helper.PanicIfError(err)
	return nil
}
