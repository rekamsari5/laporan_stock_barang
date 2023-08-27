package routers

import (
	laporanpengeluaran "laporan_stock_barang/controllers/api/v1/laporan_pengeluaran"
	models "laporan_stock_barang/models/laporan_pengeluaran_barang"

	"github.com/gin-gonic/gin"
)

func (r *routes) addLaporanPengeluaranRoute(rg *gin.RouterGroup) {
	repository := models.NewRepository(r.db)
	service := models.NewService(repository)
	controller := laporanpengeluaran.NewController(service, r.validatorFormatter, r.validate)

	pengeluaran := rg.Group("pengeluaran")
	pengeluaran.POST("/from_pengeluaran_barang", controller.CreatePengeluaran)
}
