package models

type Laporan struct {
	ID     int32 `json:"id"`
	Gudang int32 `json:"gudang"`
	Produk int32 `json:"produk"`
	QtyDus int32 `json:"qty_dus"`
	QtyPcs int32 `json:"qty_pcs"`
}

type ResultLaporan struct {
	ID          int32   `json:"id"`
	IdTrx       string  `json:"id_trx"`
	IdSupplier  int64   `json:"id_supplier"`
	Date        string  `json:"tanggal"`
	IdGudang    int32   `json:"id_gudang"`
	ListProduct Product `json:"list_produk"`
	IdCustomer  int32   `json:"id_customer"`
}

type Product struct {
	ID        int32 `json:"id"`
	Idproduct int32 `json:"id_produk"`
	QtyDus    int32 `json:"qty_dus"`
	QtyPcs    int32 `json:"qty_pcs"`
}

type ResultIdLaporan struct {
	IdTrx string `json:"id_trx"`
}
