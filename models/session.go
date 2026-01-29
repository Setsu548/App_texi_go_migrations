package models

import "time"

// Session representa la tabla sessions para registrar logins
type Session struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	UserID    int64     `gorm:"not null" json:"user_id"`
	User      *User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	Token     string    `gorm:"size:500;not null" json:"token"`
	IPAddress string    `gorm:"size:500;not null" json:"ip_address"`
	LastUsed  time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"last_used"`
	IsRevoked bool      `gorm:"default:true;not null" json:"is_revoked"`
	Brand     *string   `gorm:"size:100" json:"brand"`
	Model     *string   `gorm:"size:100" json:"model"`
	OS        *string   `gorm:"size:100" json:"os"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
}

// TableName especifica el nombre de la tabla
func (Session) TableName() string {
	return "sessions"
}
