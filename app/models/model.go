// Package models 模型通用属性和方法
package models

import "time"

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsFiels 时间戳
type CommonTimestampsFiels struct {
	CreatedAt time.Time `gorm:"created_at;index" json:"created_at,omitempty" `
	UpdatedAt time.Time `gorm:"updated_at;index" json:"updated_at,omitempty"`
}
