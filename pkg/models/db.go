package models

import (
	"time"

	"github.com/go-courier/sqlx/v2"
	"github.com/saitofun/items/pkg/depends"
)

var DB = sqlx.NewDatabase("item")

type OperationTimes struct {
	CreatedAt depends.Timestamp `db:"f_created_at,default='0'" json:"createdAt"` // 创建时间
	UpdatedAt depends.Timestamp `db:"f_updated_at,default='0'" json:"updatedAt"` // 更新时间
}

func (times *OperationTimes) MarkUpdatedAt() {
	times.UpdatedAt = depends.Timestamp(time.Now())
}

func (times *OperationTimes) MarkCreatedAt() {
	times.MarkUpdatedAt()
	times.CreatedAt = times.UpdatedAt
}

type OperationTimesWithDeletedAt struct {
	OperationTimes
	DeletedAt depends.Timestamp `db:"f_deleted_at,default='0'" json:"-"` // 删除时间
}

func (times *OperationTimesWithDeletedAt) MarkDeletedAt() {
	times.MarkUpdatedAt()
	times.DeletedAt = times.UpdatedAt
}

type PrimaryID struct {
	ID uint64 `db:"f_id,autoincrement" json:"-"`
}

type Text string

func (Text) DataType(driver string) string {
	return "text"
}
