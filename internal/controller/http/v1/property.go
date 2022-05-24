package v1

import (
	"net/http"

	"github.com/adnanmhd/go-porto-homies/internal/entity"
	"github.com/adnanmhd/go-porto-homies/internal/usecase"
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

	h := handler.Group("/property")
	{
		h.GET("properties", r.propertyList)
	}
}

func (p *propertyRoutes) propertyList(c *gin.Context) {
	list, err := p.List(c.Request.Context())
	if err != nil {
		p.Error(err, "http - v1 - listProperty")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}
	c.JSON(http.StatusOK, Response{
		Message:          "success",
		Code:             "0",
		propertyResponse: propertyResponse{list},
	})
}
