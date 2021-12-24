package pager

import "github.com/go-courier/sqlx/v2/builder"

type Pager struct {
	Size int64 `name:"pageSize,omitempty" in:"query" default:"10" validate:"@int64[-1,]"`
	NO   int64 `name:"pageNo,omitempty"   in:"query" default:"1"  validate:"@int64[1,]"`
}

func (p Pager) Offset() int64 {
	if p.Size <= 0 {
		return -1
	}
	return (p.NO - 1) * p.Size
}

func (p Pager) ToAddition(orders ...*builder.Order) []builder.Addition {
	var additions []builder.Addition

	if p.Size != -1 {
		additions = append(additions,
			builder.Limit(p.Size).Offset(p.Offset()))
	}
	additions = append(additions, builder.OrderBy(orders...))
	return additions
}

func (p Pager) ToAdditionWithComment(comment string, orders ...*builder.Order) []builder.Addition {
	return append(p.ToAddition(orders...), builder.Comment(comment))
}

type PageQueryResponse struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
