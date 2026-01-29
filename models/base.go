package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time       `gorm:"not null" json:"created_at"`
	CreatedBy uint            `gorm:"not null" json:"created_by"`
	UpdatedAt *time.Time      `gorm:"default:null" json:"updated_at"`
	UpdatedBy *uint           `gorm:"default:null" json:"updated_by"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
	DeletedBy *uint           `gorm:"default:null"`
}

type UUIDModel struct {
	UUID string `gorm:"type:varchar(500);uniqueIndex;not null" json:"uuid"`
}
