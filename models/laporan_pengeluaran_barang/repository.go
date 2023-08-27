package laporanpengeluaranbarang

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
	tbl_customer           = "MasterCustomer"
	tbl_product            = "MasterProduct"
	tbl_supplier           = "MasterSupplier"
	tbl_barang             = "MasterWarehouse"
	tbl_header_pengeluaran = "TransaksiPengeluaranBarangHeader"
	tbl_detail_pengeluaran = "TransaksiPengeluaranBarangDetail"
)

type Repository interface {
	InputHeaderPengeluaran(param models.RequestForm) (models.ResultIdLaporan, error)
	InputDetailPengeluaran(param models.RequestDetailProductTrx) error
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) InputHeaderPengeluaran(param models.RequestForm) (models.ResultIdLaporan, error) {
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

	idTrx := fmt.Sprintf("TRXOUT%04d", trxNumber)

	query := sq.Insert(tbl_header_pengeluaran).
		Columns("TrxOutNo", "WhsIdf", "TrxOutDate", "TrxOutSuppIdf", "TrxOutCustomer").
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
	err := tx.QueryRow("SELECT COALESCE(MAX(TrxOutPK), 0) FROM TransaksiPengeluaranBarangHeader").Scan(&lastID)
	return lastID, err
}

func extractAndIncrementTransactionNumber(id int) int {
	return id + 1
}

func (r *repository) InputDetailPengeluaran(param models.RequestDetailProductTrx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer tx.Rollback()

	query := sq.Insert(tbl_detail_pengeluaran).
		Columns("TrxOutIDF", "TrxOutDProductIdf", "TrxOutDQtyDus", "TrxOutDQtyPcs").
		Values(param.IdTrx, param.Idproduct, param.QtyDus, param.QtyPcs)
	sql, args, err := query.ToSql()
	helper.PanicIfError(err)

	_, err = tx.ExecContext(ctx, sql, args...)
	helper.PanicIfError(err)

	err = tx.Commit()
	helper.PanicIfError(err)
	return nil
}
