package routers

import (
	"database/sql"
	"errors"

	"laporan_stock_barang/constants"
	"laporan_stock_barang/database"
	"laporan_stock_barang/exceptions"
	"laporan_stock_barang/helper"
	"laporan_stock_barang/middleware"
	"net/http"

	formatter "github.com/abdullahPrasetio/validation-formatter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type routes struct {
	router             *gin.Engine
	db                 *sql.DB
	validatorFormatter formatter.ValidateFormatter
	validate           *validator.Validate
}

func SetupRouter() *gin.Engine {
	db, err := database.GetConnection()
	helper.PanicIfError(err)

	r := routes{
		router:   gin.Default(),
		db:       db,
		validate: formatter.Validate,
	}
	r.router.Use(gin.CustomRecovery(exceptions.ErrorHandlerRecovery))
	r.router.Use(cors.Default())
	api := r.router.Group(constants.ServerDefaultRoute, middleware.AddDefaultHeader())
	r.addLaporanPenerimaanRoute(api)
	r.addLaporanPengeluaranRoute(api)
	r.router.NoRoute(func(c *gin.Context) {
		c.JSON(404, helper.APIResponseError("Page not found", constants.ErrorPageNotFound, errors.New("Page not found").Error()))
	})
	r.router.GET("/healthz", checkHealtz)
	return r.router
}

func checkHealtz(c *gin.Context) {
	c.JSON(http.StatusOK, helper.APIResponseSuccess("success", "ok"))
}
