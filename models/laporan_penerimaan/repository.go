package laporanpenerimaan

import (
	"context"
	"database/sql"
	"fmt"
	"laporan_stock_barang/helper"
	"laporan_stock_barang/models"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type repository struct {
	db *sql.DB
}

const (
	tbl_customer          = "MasterCustomer"
	tbl_product           = "MasterProduct"
	tbl_supplier          = "MasterSupplier"
	tbl_barang            = "MasterWarehouse"
	tbl_header_penerimaan = "TransaksiPenerimaanBarangHeader"
	tbl_detail_penerimaan = "TransaksiPenerimaanBarangDetail"
)

type Repository interface {
	InputHeaderPenerimaan(param models.RequestForm) (models.ResultIdLaporan, error)
	InputDetailPenerimaan(param models.RequestDetailProductTrx) error
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) InputHeaderPenerimaan(param models.RequestForm) (models.ResultIdLaporan, error) {
	var result models.ResultIdLaporan

	location, _ := time.LoadLocation("Asia/Jakarta")

	now := time.Now().In(location)
	createdAt := now.Format("2006-01-02 15:04:05")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer tx.Rollback()

	lastID, err := getLastTransactionID(tx)
	helper.PanicIfError(err)

	trxNumber := extractAndIncrementTransactionNumber(lastID)

	idTrx := fmt.Sprintf("TRXIN%04d", trxNumber)

	query := sq.Insert(tbl_header_penerimaan).
		Columns("TrxInNo", "WhsIdf", "TrxInDate", "TrxInSuppIdf", "TrxInCustomer").
		Values(idTrx, param.IdGudang, createdAt, param.IdSupplier, param.IdCustomer)
	sql, args, err := query.ToSql()
	helper.PanicIfError(err)

	_, err = tx.ExecContext(ctx, sql, args...)
	helper.PanicIfError(err)

	err = tx.Commit()
	helper.PanicIfError(err)

	result.IdTrx = idTrx

	return result, nil
}

func getLastTransactionID(tx *sql.Tx) (int, error) {
	var lastID int
	err := tx.QueryRow("SELECT COALESCE(MAX(TrxInPK), 0) FROM TransaksiPenerimaanBarangHeader").Scan(&lastID)
	return lastID, err
}

func extractAndIncrementTransactionNumber(id int) int {
	return id + 1
}

func (r *repository) InputDetailPenerimaan(param models.RequestDetailProductTrx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer tx.Rollback()

	query := sq.Insert(tbl_detail_penerimaan).
		Columns("TrxInIDF", "TrxInDProductIdf", "TrxInDQtyDus", "TrxInDQtyPcs").
		Values(param.IdTrx, param.Idproduct, param.QtyDus, param.QtyPcs)
	sql, args, err := query.ToSql()
	helper.PanicIfError(err)

	_, err = tx.ExecContext(ctx, sql, args...)
	helper.PanicIfError(err)

	err = tx.Commit()
	helper.PanicIfError(err)
	return nil
}
