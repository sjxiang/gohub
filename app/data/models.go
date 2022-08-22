// 数据模型 model 通用属性和方法

package data

import (
	"fmt"
	"time"
)

// 基类 Model
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"` 
}


// 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index" json:"updated_at,omitempty"`
}


// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GeyStringID() string {
	return fmt.Sprintf("%v", a.ID)
} 
