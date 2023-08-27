package models

type RequestForm struct {
	IdSupplier  int64            `json:"id_supplier" binding:"required"`
	Date        string           `json:"tanggal" binding:"required"`
	IdGudang    int64            `json:"id_gudang" binding:"required"`
	ListProduct []RequestProduct `json:"list_produk"`
	IdCustomer  int64            `json:"id_customer" binding:"required"`
}

type RequestProduct struct {
	Idproduct int32 `json:"id_produk" binding:"required"`
	QtyDus    int32 `json:"qty_dus" binding:"required"`
	QtyPcs    int32 `json:"qty_pcs" binding:"required"`
}

type RequestDetailProductTrx struct {
	IdTrx     string `json:"TrxInNo"`
	Idproduct int32  `json:"id_produk"`
	QtyDus    int32  `json:"qty_dus"`
	QtyPcs    int32  `json:"qty_pcs"`
}

type SearchStock struct {
	Gudang  string `json:"gudang"`
	Product string `json:"produk"`
}
