package models

import "time"

// Employee representa la tabla employees
type Employee struct {
	ID           int64      `gorm:"primaryKey" json:"id"`
	HiringDate   time.Time  `gorm:"type:date;not null" json:"hiring_date"`
	SocialNumber string     `gorm:"size:100;not null" json:"social_number"`
	BaseSalary   float64    `gorm:"type:decimal(20,2);not null" json:"base_salary"`
	Status       string     `gorm:"type:varchar(20);not null" json:"status"`
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	CreatedBy    int64      `gorm:"not null" json:"created_by"`
	UpdatedAt    *time.Time `json:"updated_at"`
	UpdatedBy    *int64     `json:"updated_by"`
	DeletedAt    *time.Time `json:"deleted_at"`
	DeletedBy    *int64     `json:"deleted_by"`
	User         *User      `gorm:"-" json:"user,omitempty"`
}

// TableName especifica el nombre de la tabla
func (Employee) TableName() string {
	return "employees"
}

// Status constants
const (
	EmployeeStatusActive   = "active"
	EmployeeStatusInactive = "inactive"
)
