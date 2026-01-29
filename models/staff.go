package models

import "time"

// Staff representa la tabla staffs con información adicional de usuarios
type Staff struct {
	ID                int64      `gorm:"primaryKey" json:"id"`
	Email             *string    `gorm:"size:500;uniqueIndex:idx_staffs_email" json:"email"`
	PhoneNumber       string     `gorm:"size:500;not null" json:"phone_number"`
	FirstName         *string    `gorm:"size:500" json:"first_name"`
	LastName          *string    `gorm:"size:500" json:"last_name"`
	ProfilePictureURL *string    `gorm:"size:500" json:"profile_picture_url"`
	Address           *string    `gorm:"size:100" json:"address"`
	Profession        *string    `gorm:"size:100" json:"profession"`
	Gender            *string    `gorm:"size:10" json:"gender"`
	BirthDate         *string    `gorm:"size:100" json:"birth_date"`
	LocalityID        *int64     `json:"locality_id"`
	Locality          *Locality  `gorm:"foreignKey:LocalityID" json:"locality,omitempty"`
	User              *User      `gorm:"-" json:"user,omitempty"`
	CreatedAt         time.Time  `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	CreatedBy         int64      `gorm:"not null" json:"created_by"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
	DeletedBy         *int64     `json:"deleted_by,omitempty"`
}

// TableName especifica el nombre de la tabla
func (Staff) TableName() string {
	return "staffs"
}
