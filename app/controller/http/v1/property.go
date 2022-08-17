package v1

import (
	"github.com/adnanmhd/go-porto-homies/app/entity"
	"github.com/adnanmhd/go-porto-homies/app/usecase"
	"net/http"

	"github.com/adnanmhd/go-porto-homies/pkg/logger"
	"github.com/gin-gonic/gin"
)

type propertyRoutes struct {
	usecase.Property
	logger.Interface
}

type propertyResponse struct {
	PropertyList []entity.Property `json:"property_list"`
}

type Response struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	propertyResponse
}

func newPropertyRoutes(handler *gin.RouterGroup, p usecase.Property, l logger.Interface) {
	r := &propertyRoutes{p, l}

	handler.GET("properties", r.propertyList)
}

func (p *propertyRoutes) propertyList(c *gin.Context) {
	list, err := p.List(c.Request.Context())
	var resp entity.Response
	if err != nil {
		p.Error(err, "http - v1 - listProperty")
		resp.Message = "database problems"
		resp.Status = "error"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp = entity.Response{
		Status:     "0",
		Message:    "Success",
		Properties: list,
	}

	c.JSON(http.StatusOK, resp)
}
