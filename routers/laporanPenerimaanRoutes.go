package routers

import (
	laporanpenerimaan "laporan_stock_barang/controllers/api/v1/laporan_penerimaan"
	models "laporan_stock_barang/models/laporan_penerimaan"

	"github.com/gin-gonic/gin"
)

func (r *routes) addLaporanPenerimaanRoute(rg *gin.RouterGroup) {
	repository := models.NewRepository(r.db)
	service := models.NewService(repository)
	controller := laporanpenerimaan.NewController(service, r.validatorFormatter, r.validate)

	penerimaan := rg.Group("laporan")
	penerimaan.POST("/from_penerimaan_barang", controller.CreatePenerimaan)
}
