package models

import "time"

type Operator struct {
	Uuid       string     `json:"uuid"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	RolName    string     `json:"rol_name"`
	Locality   string     `json:"locality"`
	BaseSalary float64    `json:"base_salary"`
	Status     string     `json:"status"`
	CreatedAt  CustomTime `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy  *int64     `gorm:"not null" json:"created_by"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	DeletedBy  *int64     `json:"deleted_by,omitempty"`
}

type Operators []Operator
