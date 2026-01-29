package models

import "time"

// PasswordReset representa la tabla password_resets
type PasswordReset struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	UserID    int64     `gorm:"not null" json:"user_id"`
	User      *User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	Token     string    `gorm:"size:500;not null" json:"token"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
}

// TableName especifica el nombre de la tabla
func (PasswordReset) TableName() string {
	return "password_resets"
}

// PhoneVerification representa la tabla phone_verifications
type PhoneVerification struct {
	ID               int64      `gorm:"primaryKey" json:"id"`
	UserID           int64      `gorm:"not null" json:"user_id"`
	User             *User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	VerificationCode int        `gorm:"not null" json:"verification_code"`
	ExpiresAt        time.Time  `gorm:"default:CURRENT_TIMESTAMP;not null" json:"expires_at"`
	Attempts         int        `gorm:"default:0" json:"attempts"`
	IsUsed           bool       `gorm:"default:false" json:"is_used"`
	CreatedAt        time.Time  `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UsedAt           *time.Time `json:"used_at"`
}

// TableName especifica el nombre de la tabla
func (PhoneVerification) TableName() string {
	return "phone_verifications"
}
