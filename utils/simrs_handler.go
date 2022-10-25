package utils

import (
	"context"
	"github.com/dedinirtadinata/bpjs-helper/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SIMRSHandler struct {
}

func (h SIMRSHandler) GetInfoRuang(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

}

func (h SIMRSHandler) UpdateRuang(c *gin.Context) {

	var request models.BPJSRuanganModel2

	c.ShouldBind(&request)

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var ruangbaru models.BPJSRuanganModel2 = models.BPJSRuanganModel2{
		Kodekelas: request.Kodekelas,
		Koderuang: request.Koderuang,
		Namaruang: request.Namaruang,
		Kapasitas: request.Kapasitas,
		Tersedia:  request.Tersedia,
	}

	response := PostDataAplicare("aplicaresws/rest/bed/update/0059R005", ruangbaru)
	if response.Status != "OK" {
		if response.Message == "Data tidak ada di database." {
			response := PostDataAplicare("aplicaresws/rest/bed/create/0059R005", ruangbaru)
			if response.Status != "OK" {
				c.JSON(http.StatusOK, models.ResponseModel{
					Status:  "NOT_OK",
					Message: response.Message,
				})
				return
			} else {
				c.JSON(http.StatusOK, models.ResponseModel{
					Status: "OK",
				})
				return
			}
		}
		c.JSON(http.StatusOK, models.ResponseModel{
			Status:  "NOT_OK",
			Message: response.Message,
		})
		return
	} else {
		c.JSON(http.StatusOK, models.ResponseModel{
			Status: "OK",
		})
		return
	}
}

func NewSIMRSHandler(g *gin.Engine) {
	handler := &SIMRSHandler{}

	g.POST("v1/updateRuang", handler.UpdateRuang)
}
