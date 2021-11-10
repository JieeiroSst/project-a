package pagination

import (
	"github.com/JieeiroSst/itjob/model"
	"gorm.io/gorm"
	"math"
)

type paginationPage struct {
	pagination model.Pagination
}

type PaginationPage interface {
	GetOffset() int
	GetLimit() int
	GetPage() int
	GetSort() string
	Paginate(value interface{}, db *gorm.DB) (func(db *gorm.DB) *gorm.DB, model.Pagination)
}

func NewPaginationPage(pagination model.Pagination) PaginationPage {
	return &paginationPage{pagination: pagination}
}

func (p *paginationPage) GetOffset() int {
	return (p.GetPage() - 1) *p.GetLimit()
}

func (p *paginationPage) GetLimit() int {
	if p.pagination.Limit == 0 {
		p.pagination.Limit = 10
	}
	return p.pagination.Limit
}

func (p *paginationPage) GetPage() int {
	if p.pagination.Page == 0 {
		p.pagination.Page = 1
	}
	return p.pagination.Page
}

func (p *paginationPage) GetSort() string {
	if p.pagination.Sort == "" {
		p.pagination.Sort = "id desc"
	}
	return p.pagination.Sort
}

func (p *paginationPage) Paginate(value interface{}, db *gorm.DB) (func(db *gorm.DB) *gorm.DB, model.Pagination) {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	p.pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(p.pagination.Limit)))
	p.pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort())
	}, p.pagination
}