package objects

import (
	"math"

	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"
)

type Pagination struct {
	Page         int
	Limit        uint32
	Search       string
	Prev         int
	Next         int
	TotalPages   int
	TotalRecords int
}

func NewPagination() *Pagination {
	return &Pagination{}
}

func (p *Pagination) MapFromRequest(req common_input_handler.PaginationRequest) {
	p.Page = req.Page
	p.Limit = uint32(req.Limit)
	p.Search = req.Search

	if p.Page == 0 {
		p.Page = constants.DefaultPage
	}
	if p.Limit == 0 {
		p.Limit = constants.DefaultLimit
	}
}

func (p *Pagination) MapToResponse() *common_input_handler.Pagination {
	return &common_input_handler.Pagination{
		Page:         p.Page,
		Limit:        int(p.Limit),
		Prev:         p.Prev,
		Next:         p.Next,
		TotalPages:   p.TotalPages,
		TotalRecords: p.TotalRecords,
	}
}

func (p *Pagination) AllData() *Pagination {
	return &Pagination{
		Page:  constants.DefaultPage,
		Limit: constants.DefaultUnlimited,
	}
}

func (p *Pagination) GetPagination(totalRow int) {
	p.TotalPages = int(math.Ceil(float64(totalRow) / float64(p.Limit)))
	prev := 1
	if p.Page > 1 {
		prev = p.Page - 1
	}
	next := p.TotalPages
	if p.TotalPages != p.Page {
		next = p.Page + 1
	}
	p.Prev = prev
	p.Next = next
	p.TotalRecords = totalRow
}
