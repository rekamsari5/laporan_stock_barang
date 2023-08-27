package laporanpenerimaan

import (
	"fmt"
	"laporan_stock_barang/constants"
	"laporan_stock_barang/helper"
	"net/http"

	modelss "laporan_stock_barang/models"
	model "laporan_stock_barang/models/laporan_penerimaan"

	formattervalidator "github.com/abdullahPrasetio/validation-formatter"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type penerimaanController struct {
	service            model.Service
	validatorFormatter formattervalidator.ValidateFormatter
	validate           *validator.Validate
}

func NewController(service model.Service, validatorFormatter formattervalidator.ValidateFormatter, validate *validator.Validate) *penerimaanController {
	return &penerimaanController{
		service:            service,
		validatorFormatter: validatorFormatter,
		validate:           validate,
	}
}

func (h *penerimaanController) CreatePenerimaan(c *gin.Context) {
	var input modelss.RequestForm
	err := c.ShouldBindBodyWith(&input, binding.JSON)

	if err != nil {
		errorStr := fmt.Sprintf("JSON error: %v", err)
		response := helper.APIResponseError("Validation Error", constants.ErrorValidate, errorStr)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	resultHeader, err := h.service.InputHeaderPenerimaan(input)
	if err != nil {
		response := helper.APIResponseError("Error to create Header", constants.ErrorGeneral, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	for _, product := range input.ListProduct {
		detailParam := modelss.RequestDetailProductTrx{
			IdTrx:     resultHeader.IdTrx,
			Idproduct: product.Idproduct,
			QtyDus:    product.QtyDus,
			QtyPcs:    product.QtyPcs,
		}
		err := h.service.InputDetailPenerimaan(detailParam)
		if err != nil {
			response := helper.APIResponseError("Error to create detail", constants.ErrorGeneral, err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	response := helper.APIResponseSuccess("Successfully created laporan penerimaan", resultHeader)
	c.JSON(http.StatusOK, response)
}
