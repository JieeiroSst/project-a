package utils

import (
	"github.com/JieeiroSst/itjob/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type paginationPage struct {

}

type PaginationPage interface {
	GeneratePaginationFromRequest(c *gin.Context) model.PaginationPage
}

func NewPaginationPage() PaginationPage {
	return &paginationPage{}
}

func (p *paginationPage) GeneratePaginationFromRequest(c *gin.Context) model.PaginationPage {
	limit := 2
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return model.PaginationPage{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}